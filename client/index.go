package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type SetArgs struct {
    A, B string
}

type GetArgs struct {
    A string
}

func SET(args SetArgs,result* string){
    client, err := rpc.DialHTTP("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Error connecting to server:", err)
    }
    err = client.Call("Arith.SET", args, &result)
    if err != nil {
        log.Fatal("Error calling remote function:", err)
    }
    fmt.Println("Result:", result)
}

func GET_ALL(args GetArgs, result* string){
    client, err := rpc.DialHTTP("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Error connecting to server:", err)
    }
    err = client.Call("Arith.GET_ALL",args, &result)
    if err != nil {
        log.Fatal("Error calling remote function:", err)
    }
    fmt.Println("Result:", &result)
}

func GET(args GetArgs, result* string){
    client, err := rpc.DialHTTP("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Error connecting to server:", err)
    }
    err = client.Call("Arith.GET",args, &result)
    if err != nil {
        log.Fatal("Error calling remote function:", err)
    }
    fmt.Println("Result:", *result)
}

func main() {
	fn := os.Args[1]
	var key string;
	var value string;
	switch(fn){
	case "SET":
		key = os.Args[2]
		value = os.Args[3]
		args := SetArgs{key,value}
		var result string
		SET(args,&result);
		break;
	case "GET_ALL":
		key:=os.Args[2]
		args := GetArgs{key}
		var result string
		GET_ALL(args,&result);
		break;
	case "GET":
		key:=os.Args[2]
		args := GetArgs{key}
		var result string
		GET(args,&result);
		break;
	default:
		fmt.Print("Invalid command")
		break;
	}
}