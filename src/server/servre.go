package server

import (
	"filestorage/src/library"
	"fmt"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Server file storage server
type Server struct {
	*logrus.Logger
	config  *ConfigModel
	gateway *http.Server
}

// NewFileStorageServer new file storage server
func NewFileStorageServer() Server {
	server := Server{
		Logger: library.NewLogger("./log", "logger", true),
		config: &ConfigModel{},
	}
	if err := library.NewConfig("./cfg", "config", "json", server.config); err != nil {
		server.Fatalf("setup config error: %v", err)
	}
	fmt.Printf("conf: %+v", *server.config)
	server.gateway = &http.Server{
		Addr:    server.config.ListenAddress,
		Handler: server.router(server.config.StoragePath, server.config.DirList),
	}
	return server
}

// Run ..
func (srv *Server) Run() {
	if err := srv.gateway.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
