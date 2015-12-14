package main

import "fmt"

func main(){
	list :=[]int{32,43,56,79,61,42,46,52,86,57}
	fmt.Println("unsorted: ", list)

	selectionSort(list)
	fmt.Println("sorted: ", list)
}

func selectionSort(intArray []int){
	size := len(intArray)
	var minIndex int
	for outer:=0;outer<size;outer++{
		minIndex=outer;
		for inner:=outer+1;inner<size;inner++{
			if intArray[inner]<intArray[minIndex]{
				minIndex=inner
			}
			intArray[minIndex],intArray[outer]=intArray[outer], intArray[minIndex]
		}
		fmt.Println("Iteration number: ",outer," result: ",intArray)
	}
}
