// Package cmd have functions for parsing cmd flags "CommandFlags", "CmdParse"
package cmd

import (
	"flag"
	"time"
)

// CommandFlags to store configuration provided via command line
type CommandFlags struct {
	Wait time.Duration
	Host,
	Port,
	Config string
}

// cmd.CmdParse for command line flags parsing return *CommandFlags
func CmdParse() *CommandFlags {
	var flags = CommandFlags{}
	flag.DurationVar(&flags.Wait, "wait", time.Second*15, "The duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&flags.Host, "host", "0.0.0.0", "Host name to run the server")
	flag.StringVar(&flags.Port, "port", "8080", "Server port to run the server")
	flag.StringVar(&flags.Config, "config", "config/config.json", "Server configuration")

	flag.Parse()
	return &flags
}
