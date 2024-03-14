package main

import (
	"fmt"
	"hash/fnv"
)

type HashMap struct{
	size int
	Bucket []LinkedList
}

func NewHashMap(size int) *HashMap{
	return &HashMap{
		size:size,
		Bucket:make([]LinkedList, size),
	}
}

func (hashMap HashMap) hash(item string) int{
	h := fnv.New32a()
    h.Write([]byte(string(item)))
    // Get the 32-bit hash value as an unsigned integer
    hash := h.Sum32()
	fmt.Println("hash :",int(hash) % hashMap.size)
    // Return the hash value modulo the array size to ensure it fits within the array bounds
    return int(hash) % hashMap.size
}

func (hashMap *HashMap) put(key string,value string){
	index:=hashMap.hash(key);
	hashMap.Bucket[index].Append(key,value); 
}

// func (hashMap HashMap) get(key string) any{
// 	index := hashMap.hash(key)
// 	if(hashMap.Bucket[index]!=nil){
// 		return hashMap.Bucket[hashMap.hash(key)];
// 	}
// 	return -1
// }

func (hashMap HashMap) getAll(){
	for i := 0; i < hashMap.size; i++ {
		hashMap.Bucket[i].PrintLL()	
	}
}

//linked list
type Node struct{
	key string
	value string
	next *Node
}

type LinkedList struct{
	head *Node
}

func (l *LinkedList) Append(key string,value string) {
    newNode := &Node{key:key,value: value}

    // If the list is empty, set the new node as the head
    if l.head == nil {
        l.head = newNode
        return
    }

    // Traverse the list to find the last node
    current := l.head
    for current.next != nil {
        current = current.next
    }

    // Append the new node to the end of the list
    current.next = newNode
}

func (l *LinkedList) PrintLL() {
    current := l.head
    for current != nil {
        fmt.Println("ll=>","{",current.key,",",current.value,"}")
        current = current.next
    }
    fmt.Println("nil")
}