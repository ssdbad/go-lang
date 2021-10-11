package assignment001

import (
	"container/list"
	"fmt"
)

func main() {

	var searchString = "WXYZ"

	listOfString := list.New()
	abcd := listOfString.PushFront("ABCD")
	wxyz := listOfString.PushBack("WXYZ")

	listOfString.InsertBefore("sample", abcd)
	listOfString.InsertAfter("example", wxyz)

	listOfString.InsertAfter("EFGH", abcd)
	listOfString.InsertBefore("STUV", wxyz)

	for i := listOfString.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, ", ")
	}

	fmt.Println("\n  Size of List : ", listOfString.Len())
	var count = 1
	for i := listOfString.Front(); i != nil; i = i.Next() {
		if i.Value == searchString {
			fmt.Println("\n\n", searchString, "is found in current list at index : ", count)
			break
		}
		count++
	}
}
