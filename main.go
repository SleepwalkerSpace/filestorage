package main

import "filestorage/src/server"

func main() {
	srv := server.NewFileStorageServer()
	srv.Run()
}
