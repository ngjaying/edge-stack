package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"time"
)

type Handler struct{}

func (e *Handler) Echo123(args []string, response *string) error {
	log.Printf("args=%#v\n", args)
	*response = fmt.Sprintf("%#v", args)
	return nil
}

// dup'd in cli.go
type EchoArgument struct {
	A string
	B string
	C string
}

func (e *Handler) EchoStruct(arg EchoArgument, response *string) error {
	log.Printf("args=%#v\n", arg)
	*response = fmt.Sprintf("%#v", arg)
	return nil
}

func (e *Handler) SayHello(name string, response *map[string]interface{}) error {
	log.Printf("args=%#v\n", name)
	(*response)["message"] = name
	return nil
}

func main() {
	rpcSrv := rpc.NewServer()
	_ = rpcSrv.Register(&Handler{})
	srvRpc := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", "127.0.0.1", 52200),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      rpcSrv,
	}
	fmt.Println("Start rpc server")
	_ = srvRpc.ListenAndServe()
	fmt.Println("Exit rpc server")
}
