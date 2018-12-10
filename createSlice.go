package main

import "fmt"

func sliceMake(s []int)  {
	fmt.Printf("s=%d,len=%d,cap=%d\n",s,len(s),cap(s))
}
func main()  {
	var s []int

	for i := 0;i < 100;i++{
		s = append(s,2 * i + 1)
		sliceMake(s)
	}

	s2 :=[]int{2,3,4,5,6}
	sliceMake(s2)

	s3 := make([]int,16)
	sliceMake(s3)

	s4 := make([]int,10,32)
	sliceMake(s4)

	//-----copy
	copy(s3,s2)
	sliceMake(s3)

	//------delete
	s5 := append(s3[:3],s3[4:] ...)
	fmt.Println(s5)
}
