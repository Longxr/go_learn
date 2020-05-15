package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
	. "go_learn/rpc/jsonrpc_http/protocol"
)

func main() {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	server.RegisterService(new(Arith), "")
	r := mux.NewRouter()
	r.Handle("/rpc", server)
	log.Println("JSON RPC service listen and serving on port 2334")
	if err := http.ListenAndServe(":2334", r); err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}
