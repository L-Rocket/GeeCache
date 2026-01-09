package main

import "fmt"

type bird interface {
	fly()
}

type sparrow struct {
	name string
	age  int
}

func (b *sparrow) fly() {
	b.age = 3
	fmt.Printf("%s is flying\n", b.name)
}

func main() {
	var sparrow1 = sparrow{name: "sparrow1", age: 2}
	var b = sparrow1
	fmt.Println(sparrow1.age)
	fmt.Println(b.age)
	b.fly()
	fmt.Println("---------")
	fmt.Println(sparrow1.age)
	fmt.Println(b.age)
}

/*
func main() {
	// var b bird = &sparrow{name: "sparrow", age: 2}
	// b := &sparrow{name: "sparrow", age: 2}
	var s sparrow = sparrow{name: "sparrow", age: 2}
	var b bird = s
	b.fly()
	// b.fly() -> (&b).fly() no!
	// b.fly()
}


func (b sparrow) fly() {
	fmt.Printf("%s is flying\n", b.name)
}

func main() {
	// var b bird = &sparrow{name: "sparrow", age: 2}
	// b := &sparrow{name: "sparrow", age: 2}
	var s *sparrow = &sparrow{name: "sparrow", age: 2}
	var b bird = s
	b.fly()
	// b.fly() -> (*b).fly() yes!

	// b.fly()
}

*/
