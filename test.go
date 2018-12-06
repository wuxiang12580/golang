package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

//全局变量不能用:=
var aa = 3;
var ss = "kk"

var (
	dd = 44
	ee = "rr"
)

func param()  {
	var a int
	var b string
	fmt.Printf("%d %q\n",a,b)
}
func paramInit()  {
	var a int = 3
	var b string = "hh"
	fmt.Println(a,b)
}

func paramInteType()  {
	a,b := 3,4
	c := "qq"
	b = 7
	fmt.Println(a,b,c)
}

func t()  {
	a,b := 3,4
	var c int
	c = int(math.Sqrt(float64(a*a + b* b)))
	fmt.Println(c)
}

func ela()  {
	a := 3 + 4i
	fmt.Println(cmplx.Abs(a))
}
func main()  {
	fmt.Println("hello world")
	param()
	paramInit()
	paramInteType()
	fmt.Println(aa,ss,dd,ee)
	ela()
	t()
}
