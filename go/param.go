package _go

import (
	"io/ioutil"
	"fmt"
)

func grade(score int) string  {
	var g = ""
	switch {
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	default:
		panic(fmt.Sprintf("wrong score %d",score))
	}
	return  g
}

func main()  {
	const filename  = "abc.txt"
	if content,err := ioutil.ReadFile(filename);err!=nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n",content)
	}
	/*if err != nil {
		fmt.Println(err)
	}else{
		fmt.Printf("%s\n",content)
	}*/

	score := grade(80)
	fmt.Println(score)
	fmt.Println(grade(0))
	fmt.Println(grade(95))

}
