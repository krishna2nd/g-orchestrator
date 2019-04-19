package main

import (
	"gecgithub01.walmart.com/k0k00bt/g-orchestrator/cmd"
	"gecgithub01.walmart.com/k0k00bt/g-orchestrator/server"
)

func main() {
	var flags *cmd.CommandFlags = cmd.CmdParse()
	server.Start(flags)
}
