package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

func welcomeClient(){
	
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!") // Write the response to the client
}

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

var myHash *HashMap;

func main() {

 	mapSize := 5;
	myHash = NewHashMap(mapSize);
	
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

// func main(){
// 	fmt.Println("poopoo connected");

// 	mapSize := 5;
// 	myHash := NewHashMap(mapSize);

// 	scanner := bufio.NewScanner(os.Stdin)

// 	x:=0;
// 	for{
// 		if(x >= mapSize){
// 			break;
// 		}

// 		fmt.Print("Enter key:")
// 		scanner.Scan()
// 		key := scanner.Text()
// 		fmt.Print("Enter value:")
// 		scanner.Scan()
// 		value := scanner.Text()

// 		fmt.Println(key,value);
// 		myHash.put(key,value);
// 		x+=1;
// 		fmt.Println(x)
// 	}

// 	myHash.getAll()
// 	// fmt.Println("map:",myHash.getAll());

// 	// myHash.put(1,10)
// 	// myHash.put(2,20)
// 	// fmt.Println(myHash.getAll())
// 	// myHash.put(2,25)
// 	// myHash.put(3,30)
// 	// fmt.Println(myHash.getAll())

// }

// func main(){
// 	arr := Data{
// 		Size: 5,
// 		A: "sad",
// 	}

// 	scanner := bufio.NewScanner(os.Stdin)

// 	// Prompt the user to enter their name
// 	fmt.Print("Enter your name: ")

// 	// Read the input from the command line
// 	scanner.Scan()
// 	name := scanner.Text()

// 	arr.put(name);

// 	fmt.Println("",arr.get());
// }

// func main() {
//     // Define a handler function for the root route "/"
//     handler := func(w http.ResponseWriter, r *http.Request) {
//         fmt.Fprintf(w, "Hello, World!") // Write "Hello, World!" to the response writer
//     }

//     // Register the handler function for the root route "/"
//     http.HandleFunc("/", handler)

//     // Start the HTTP server on port 8080
//     fmt.Println("Server listening on port 8080...")
//     http.ListenAndServe(":8080", nil)
// }

// DB
// |-Client
// | |--project
// |	 |---table
// |	 |---table
// | |--project
// |-Client
// |-Client