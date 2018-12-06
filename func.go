package main

import (
	"fmt"
	"math"
	"runtime"
	"reflect"
)

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

func div(a,b int)(int,int){
	return a/b,a%b
}

func div2(a,b int)(int ,error){
	return a/b,nil
}

func apply(op func(int,int) int ,a,b int) int  {
	fmt.Printf("calling %s with %d,%d\n",runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(),a,b)
	return op(a,b)
}

func sum(numbers... int) int  {
	var s = 0;
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
func main()  {
	fmt.Println(sum(1,2,3,4,5))

	res := eval(3,4,"+")
	fmt.Println(res)

	a,b := div(13,4)
	fmt.Printf("13除以4等于，商%d,余数%d \n",a,b)

	c,_ := div2(13,3)
	fmt.Printf("13除以3等于，商%d,余数%d \n",c)

	//if e,err := div2(13,0); err!=nil{
	//	panic(err)
	//}else {
	//	fmt.Println(e)
	//}

	fmt.Println(apply(
		func(a int, b int) int {
			return int(math.Pow(float64(a),float64(b)))
		},3,4))

}
