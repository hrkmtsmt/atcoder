package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var size, start, end int

	fmt.Scanf("%d %d %d", &size, &start, &end)

	array := make([]string, size)

	reverseBiginNumber := end

	for i := range array {
		currentNumber := i + 1
		diff := end - start + 1

		if diff == 0 || start > currentNumber || end < currentNumber {
			array[i] = strconv.Itoa(currentNumber)

			continue
		}

		array[i] = strconv.Itoa(reverseBiginNumber)
		reverseBiginNumber = reverseBiginNumber - 1
	}

	fmt.Println(strings.Join(array, " "))
}
