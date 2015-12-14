package main

import "fmt"

func main(){
	list :=[]int{32,43,56,79,61,42,46,52,86,57}
	fmt.Println("unsorted: ", list)

	insertionSort(list)
	fmt.Println("sorted: ", list)
}

func insertionSort(intArray []int){
	size := len(intArray)
	var inner int
	var temp int
	for outer:=1;outer<size;outer++{
		temp=intArray[outer]
		inner=outer
		for inner>0 && intArray[inner-1]>=temp{
			intArray[inner]=intArray[inner-1]
			inner--
		}
		intArray[inner]=temp
		fmt.Println("Iteration number: ",outer," result: ",intArray)
	}
}
