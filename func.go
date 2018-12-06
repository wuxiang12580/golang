package main

import "fmt"

func eval(a,b int ,op string) int  {
	switch op {
	case "+":
		return a+b
	case "-":
		return a-b
	case "*":
		return a*b
	case "/":
		return a/b
	default:
		panic("not support "+op)
	}
}
func main()  {
	res := eval(3,4,"+")
	fmt.Println(res)

	res1 := eval(3,4,"*")
	fmt.Println(res1)

	res2 := eval(1,3,"/")
	fmt.Println(res2)

	res3 := eval(1,3,"-")
	fmt.Println(res3)
}
