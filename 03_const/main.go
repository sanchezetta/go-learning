package main

import "fmt"

const pi float32 = 3.1415

func main() {
	//pi = 2 - ошибка

	/*
		иницилизация константы возможна только значениями или другими константами, переменными нельзя
		возможна иницилизация вида:

			const (
				name1 type1 = val1
				name2 type2 = val2
			)

			const name = val

			const name1, name2 = val1, val2

		обязательна иницилизация начальными значениями

		автоматическая иницилизация предыдущими значениями:
			const (
				a = 1
				b
				c
				d =2
				e
				f
			)
			fmt.Println(a, b, c, d, e, f) выведет 1, 1, 1, 2, 2, 2

		iota - авто счетчик для констант

		const (
			a = iota равно 0
			b = iota равно 1
			c = iota равно 2
		)

		const (
			a = iota равно 0
			b        равно 1
			c        равно 2
		)
	*/
	fmt.Println(pi)
}
