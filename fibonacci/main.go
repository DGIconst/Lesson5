package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibo(n-2) + fibo(n-1)
}

func main() {
	var a int
	fmt.Println("Введите число")
	fmt.Scanln(&a)
	fmt.Println("Числом Фибоначи", a, "является", fibo(a))

}
