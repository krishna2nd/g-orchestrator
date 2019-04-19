SHELL :=/bin/bash
FILES :=$(shell find . | grep -v 'github\|golang' | grep '\.go')
PACKAGES := $(shell ls | grep -v 'github\|golang\|test')
PLUGIN_DIR := plugins
SYMBOLS_DIR := sym
PLUGINS := $(shell find plugins -type d -maxdepth 1 -mindepth 1 -exec basename {} \;)

.PHONY = all clean
CC = go

SRC := main.go
BIN := orchestrator
COV := coverage
TEST:= ${CC} test -v 
LINT:=../../../../bin/golint
BLACK:=0
RED:=1
GREEN:=2
YELLOW:=3
BLUE:=4
VIOLET:=5
LBLUE:=6
WHITE:=7

define clrprint
    @tput setab $1
	@tput setaf $2
	@tput bold
    @echo $3
    @tput sgr0
endef

all: clean setup build

clean:
	@$(call clrprint, ${BLACK}, ${VIOLET}, "Cleaning..")
	@rm -f orchestrator & true
	@rm -f sym/* & true

setup:
	@export GOPATH=`pwd`/;
	@[ -d sym ] ||  mkdir sym
	#@[ -f ./bin/golint ] || ( echo "Downloading golint"  && ${CC} get -u golang.org/x/lint/golint )
	#@[ -f ./bin/dep ] || ( echo "Downloading dep"  && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | bash)
	#./bin/dep init

plugin: ${SYMBOLS_DIR}
	@$(call clrprint, ${BLACK}, ${VIOLET}, "Building plugins..")
	@for plugin in ${PLUGINS}; do \
		${CC} build -buildmode=plugin -o ${SYMBOLS_DIR}/$${plugin}.so ${PLUGIN_DIR}/$${plugin}/$${plugin}.go; \
	done

build: main.go clean setup lint plugin
	@$(call clrprint, ${BLACK}, ${GREEN}, "Building orchestrator..")
	@${CC} build -o ${BIN} ${SRC};
	@chmod 755 ${BIN}

run: main.go setup
	@${CC} run ${SRC};

test: lint unit_test func_test bench_test

func_test: lint test/ setup
	@$(call clrprint, ${BLACK}, ${GREEN}, "Functional testing starts")
	@${TEST} ./test  -cover -outputdir ${COV} -coverprofile func_cover.out

unit_test: lint setup
	@$(call clrprint, ${BLACK}, ${LBLUE}, "Unit testing starts")
	@${TEST} ${PACKAGES} -cover -outputdir ${COV} -coverprofile=unit_coverage.out;
	@tput sgr0

bench_test: lint setup
	@$(call clrprint, ${BLACK}, ${VIOLET}, "Benchmark testing starts")
	@${TEST} -benchmem -outputdir ${COV} \
		-blockprofile bench_block.out \
		-coverprofile bench_cover.out \
		-cpuprofile bench_cpu.out \
		-memprofile bench_mem.out \
		-trace bench_trace.out

lint: ${FILES}
	@$(call clrprint, ${BLACK}, ${YELLOW}, "Formating files..")
	@for file in ${FILES}; do ${CC} fmt $${file}; done
	@$(call clrprint, ${BLACK}, ${BLUE}, "Linting files..")
	@for file in ${FILES}; do ${LINT} $${file}; done
