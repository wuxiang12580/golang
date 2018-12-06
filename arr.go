package main

import "fmt"

func printArr(arr *[5]int){
	arr[0] = 100
	for i,v := range arr {
		fmt.Println(i,v)
	}
}
func main()  {
	var arr [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{2,3,4,5}
	var grid [4][5] int
	fmt.Println(arr,arr2,arr3,grid)

	for i:=0;i<len(arr2) ;i++  {
		fmt.Println(arr2[i])
	}
	for i,v := range arr2{
		fmt.Println(i,v)
	}

	z := 0
	for _,i := range arr3 {
		z += i
	}
	fmt.Println(z)

	printArr(&arr)
}
