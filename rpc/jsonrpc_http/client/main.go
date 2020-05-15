package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
	. "go_learn/rpc/jsonrpc_http/protocol"
)

func main() {
	url := "http://localhost:2334/rpc"
	req := ArithRequest{7, 2}
	var result ArithResponse

	message, err := json.EncodeClientRequest("Arith.Multiply", req)
	if (err != nil) {
		log.Fatalf("json EncodeClientRequest %s", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(message))
	defer resp.Body.Close()
	if (err != nil) {
		log.Fatalf("http post %s", err)
	}

	err = json.DecodeClientResponse(resp.Body, &result)
	if (err != nil) {
		log.Fatalf("json DecodeClientResponse %s", err)
	}
	fmt.Printf("%d * %d = %d\n", req.A, req.B, result.Pro)

	message, err = json.EncodeClientRequest("Arith.Divide", req)
	if (err != nil) {
		log.Fatalf("json EncodeClientRequest %s", err)
	}

	resp, err = http.Post(url, "application/json", bytes.NewReader(message))
	if (err != nil) {
		log.Fatalf("http post %s", err)
	}

	err = json.DecodeClientResponse(resp.Body, &result)
	if (err != nil) {
		log.Fatalf("json DecodeClientResponse %s", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, result.Quo, result.Rem)
}