package main

import (
	"github.com/krishna2nd/g-orchestrator/cmd"
	"github.com/krishna2nd/g-orchestrator/server"
)

func main() {
	var flags *cmd.CommandFlags = cmd.CmdParse()
	server.Start(flags)
}
