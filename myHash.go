package main

type HashMap struct{
	size int
	Node []any
}

func NewHashMap(size int) *HashMap{
	return &HashMap{
		size:size,
		Node:make([]any, size),
	}
}

func (hashMap HashMap) hash(item int) int{
	return item % hashMap.size;
}

func (hashMap *HashMap) put(key int,value any){
	index:=hashMap.hash(key);
	hashMap.Node[index] = value;
}

func (hashMap HashMap) get(key int) any{
	index := hashMap.hash(key)
	if(hashMap.Node[index]!=nil){
		return hashMap.Node[hashMap.hash(key)];
	}
	return -1
}

func (hashMap HashMap) getAll() []any{
	return hashMap.Node;
}