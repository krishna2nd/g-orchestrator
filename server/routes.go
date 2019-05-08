package server

import (
	"fmt"
	"net/http"
	"plugin"

	"github.com/gorilla/mux"
	"github.com/krishna2nd/g-orchestrator/service"
)

func InitRouter(config *Config) *mux.Router {
	var r *mux.Router = mux.NewRouter()

	// r.HandleFunc("/stores/{id}", Handler).Methods("GET")
	for _, plugin := range config.Plugins {
		linkHandlers(r, plugin, config.PluginDir)
	}
	r.Use(LoggingMiddleware)
	http.Handle("/", r)

	return r
}

func linkHandlers(r *mux.Router, pluginName, PluginDir string) {
	plugType, err := plugin.Open(PluginDir + "/" + pluginName + ".so")
	if err != nil {
		fmt.Println("Error in opening checkout symbol", err)
	}
	serviceType, errL := plugType.Lookup("Service")
	if errL != nil {
		fmt.Println("Error in Lookup Interface", errL)
	}

	serviceInterface, _ := serviceType.(service.Service)
	r.HandleFunc(serviceInterface.Handle(), serviceInterface.HandlerFunc).Methods("GET")

}
