package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

const filename  = "abc.txt"

func readFile()  {
	file,err := os.Open(filename)
	if err !=nil{
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan()  {
		fmt.Println(scanner.Text())
	}
}


func forever()  {
	for  {
		fmt.Println("abc")
	}
}

func intToBin(n int) string  {
	result := ""
	for ; n > 0 ; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func testInt(n int)  {
	//n /= 2 // n = n / 2
	n %= 2
	fmt.Println(n)
}
func main()  {
	forever()
	readFile()
	fmt.Println(intToBin(5))
	fmt.Println(intToBin(13))

	testInt(13)

	sum := 0
	for i:= 0;i <= 100;i++{
		sum += i
	}
	fmt.Println(sum)
}
