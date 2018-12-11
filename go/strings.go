package _go

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	s:="yes我爱慕课网！"
	fmt.Println(s)
	fmt.Println(len(s))
	for _,b := range []byte(s){
		fmt.Printf("%Xn ",b)
	}
	fmt.Println("\n")


	for i,b := range  s{
		fmt.Printf("(%d,%X)",i,b)
	}
	fmt.Println()

	for i,ch := range []rune(s){
		fmt.Printf("(%d,%c)",i,ch)
	}
	fmt.Println()

	fmt.Println(utf8.RuneCountInString(s))
}
