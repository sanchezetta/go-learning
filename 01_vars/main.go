package main

import "fmt"

func main() {
	/*
		var name type
		var name type = value
		var name = value
		name := value
		name1, name2 := val1, val2
		var name1, name2 type
		var (
			name1 type1
			name2 type2
		)
		var (
			name1 = var1
			name2 = var2
		)
	*/
	var hello string
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)
	a, b, c = 4, 5, 6
	fmt.Println(a, b, c)
	hello = "hi!!!"
	fmt.Println(hello)
}
