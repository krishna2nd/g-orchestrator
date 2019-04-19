package server

import (
	"log"
	"net/http"
	"time"

	"gecgithub01.walmart.com/k0k00bt/g-orchestrator/cmd"
	"github.com/gorilla/mux"
)

func Start(flags *cmd.CommandFlags) {
	var config *Config = ReadConfig(flags.Config)
	var r *mux.Router = InitRouter(config)
	RouteWalk(r)
	server := &http.Server{
		Addr:         flags.Host + ":" + flags.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	GracefulExit(server, flags)
}
