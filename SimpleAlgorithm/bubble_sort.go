package main

import "fmt"

func main(){
	list :=[]int{32,43,56,79,61,42,46,52,86,57}
	fmt.Println("unsorted: ", list)

	bubblesort(list)
	fmt.Println("sorted: ", list)
}

func bubblesort(intArray []int){
	size := len(intArray)
	for i:=0;i<size;i++{
		for j:=size-1;j>=i+1;j--{
			if intArray[j]<intArray[j-1]{
				intArray[j],intArray[j-1]=intArray[j-1], intArray[j]
			}
		}
	}
}
