// package main

// import (
// 	"fmt"
// )

// // func main() {

// // 	f()
// // 	fmt.Println("Returned normally from f.")
// // }

// func f() {
// 	defer func() {
// 		// 可以取得 panic 的回傳值
// 		r := recover()
// 		if r != nil {
// 			fmt.Println("Recovered in f", r)
// 		}
// 	}()

// 	fmt.Println("Calling g.")
// 	g(0)
// 	fmt.Println("Returned normally from g.")
// }

// func g(i int) {
// 	if i > 3 {
// 		panicNumber := fmt.Sprintf("%v", i)
// 		fmt.Println("Panicking!", panicNumber)
// 		// log.Fatalln(panicNumber)  // 如果使用 fatal 程式會直接終止，defer 也不會執行到
// 		panic(panicNumber)
// 	}
// 	defer fmt.Println("Defer in g", i)
// 	fmt.Println("Printing in g", i)
// 	g(i + 1)
// }

// ////////////////////////////////////////
// type BaseBird struct {
// 	age int
// }

// func (this *BaseBird) Cal() {
// 	this.Add()
// }
// func (this *BaseBird) Add() {
// 	fmt.Printf("before add: age=%d\n", this.age)
// 	this.age = this.age + 1
// 	fmt.Printf("after add: age=%d\n", this.age)
// }

// type DerivedBird struct {
// 	BaseBird
// }

// func (this *DerivedBird) Add() {
// 	fmt.Printf("before add: age=%d\n", this.age)
// 	this.age = this.age + 2
// 	fmt.Printf("after add: age=%d\n", this.age)
// }

// func main() {
// 	// var b1, b2 Bird
// 	// b1 = &BaseBird{age: 1}
// 	// b2 = &DerivedBird{BaseBird{age: 1}}
// 	// Cal(b1)
// 	// Cal(b2)

// 	var b1 BaseBird
// 	var b2 DerivedBird
// 	b1 = BaseBird{age: 1}
// 	b1.Cal()
// 	b2 = DerivedBird{BaseBird{1}}
// 	b2.Cal()
// }
