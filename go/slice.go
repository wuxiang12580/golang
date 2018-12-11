//切片
package _go

import "fmt"

func updateSice(arr []int)  {
	arr[0] = 100
}
func main()  {
	//slice 可以超过len 但是不能超过cap
	arr := [...]int{0,1,2,3,4,5,6}
	s := arr[2:6]
	s2 := s[3:5]
	fmt.Println(s,"-----",s2,"---")
	a := arr[:4]
	b := arr[5:]
	d := arr[:]
	e := arr[6:7]
	fmt.Println(s,a,b,d,e)
	updateSice(a)
	fmt.Println(a)
	updateSice(b)
	fmt.Println(b)

	//append-----
	arr2 := [...]int{0,1,2,3,4,5,6}
	s4 := arr2[2:6]
	s5 := s4[3:5]
	s6 := append(s4,10)
	s7 := append(s5,11)
	fmt.Println(s6,s7)
	fmt.Println(arr2)


}