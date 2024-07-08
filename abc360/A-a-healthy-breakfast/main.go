package main

import (
	"fmt"
	"strings"
)

func main() {
	var menu string
	fmt.Scan(&menu)

	const RICE string = "R"
	const MISO_SOUP string = "M"

	array := strings.Split(menu, "")
	left := array[0]
	center := array[1]

	if MISO_SOUP == left {
		fmt.Println("No")

		return
	}

	if MISO_SOUP == center && RICE != left {
		fmt.Println("No")

		return
	}

	fmt.Println("Yes")
}
