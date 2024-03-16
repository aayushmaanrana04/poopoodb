package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)


type Args struct {
    A, B string
}
type GetArgs struct {
    A string
}

type Arith string

func (t *Arith) SET(args *Args, reply *string) error {
	myHash.put(args.A,args.B);
    fmt.Println(myHash.getAll())

    *reply =  myHash.getAll()
    return nil
}

func (t *Arith) GET_ALL(args *GetArgs,reply *string) error {
    *reply = myHash.getAll()
    return nil
}

func (t *Arith) GET(args *GetArgs,reply *string) error {
    *reply = myHash.Get(args.A);
    return nil
}

func _init(){
	arith := new(Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()
    l, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    // defer l.Close() 
    fmt.Println("Server listening on port 8080")
    http.Serve(l, nil)
}