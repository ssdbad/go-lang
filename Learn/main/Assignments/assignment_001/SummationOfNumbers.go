package assignment001

import (
	"container/list"
	"fmt"
	"strconv"
)

func main() {

	listOfNumbers := list.New()
	listOfNumbers.PushFront(5)
	listOfNumbers.PushBack(4)
	listOfNumbers.PushFront(10)
	listOfNumbers.PushBack(15)
	var sum = 0
	for i := listOfNumbers.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, ",")
		value := fmt.Sprintf("%v", i.Value)
		intType, _ := strconv.Atoi(value)
		sum = sum + intType

	}

	fmt.Println("\n\n\nSummation of all numbers : ", sum)
}
