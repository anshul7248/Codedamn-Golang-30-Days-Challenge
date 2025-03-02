package main

import "fmt"

func main() {
	var str string

	fmt.Println("Enter the string")
	fmt.Scan(&str)

	reverseString(str)

}

func reverseString(str string) {
	strCount := len([]rune(str))

	for i := strCount - 1; i >= 0; i-- {
		fmt.Printf("%c", str[i])

	}
	fmt.Println()

	fmt.Println("String count-->>", strCount)
}
