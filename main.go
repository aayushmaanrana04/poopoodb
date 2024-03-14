package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("poopoo connected");

	myHash := NewHashMap(5);

	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user to enter their name
	fmt.Print("Enter key,value:")

	// Read the input from the command line
	scanner.Scan()
	name := scanner.Text()

	myHash.put(1,name);

	fmt.Println("",myHash.getAll());

	// myHash.put(1,10)
	// myHash.put(2,20)
	// fmt.Println(myHash.getAll())
	// myHash.put(2,25)
	// myHash.put(3,30)
	// fmt.Println(myHash.getAll())

}

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