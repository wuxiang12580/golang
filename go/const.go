package _go

import "fmt"

const x  = 3.14

func main()  {
	const a,b,c  int = 3,4,5
	const e string = "1111"
	fmt.Println(a,b,c,e,x)

	const  (
		n int = 9
		m string = "hello"
	)
	fmt.Println(n,m)

	const  (
		j = iota + 1
		k
		l
		p
	)
	fmt.Println(j,k,l,p)
}
