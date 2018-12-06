package main

import "fmt"

func swap(a,b *int)  {
	*a,*b = *b,*a
}
func main()  {
	a,b := 3,4
	swap(&a,&b)
	fmt.Println(a,b)

	//var a int = 2
	//var p *int = &a
	//*p = 4
	//fmt.Println(a)
}
