package main

import "fmt"

func fibonacci() func() int {
	numbers := make(map[int]int)
	n := 0
	return func() int {
		if n == 0 {
			numbers[n] = 0
			n++
			return 0
		}
		if n == 1 {
			numbers[n] = 1
			n++
			return 1
		}
		number := numbers[n-1] + numbers[n-2]
		numbers[n] = number
		n++
		return number
	}
}

func main() {
	var a int
	fmt.Println("Введите число")
	fmt.Scanln(&a)
	f := fibonacci()
	for i := -1; i < a; i++ {
		fmt.Print(" ", f())
	}
}

// не уверен что правильно понял задание, 3  дня читал изучал  справочную литературу и вот это вышло))
// прошу пояснить если сделал не правильно