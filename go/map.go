package _go

import "fmt"

func main()  {
	m := map[string]string{
		"name" : "java",
		"course" : "golang",
		"num" : "22",
		"site" : "imooc",
	}
	fmt.Println(m)

	m2 := make(map[string]int)
	fmt.Println(m2)

	for k,v := range m {
		fmt.Println(k,v)
	}

	if num,ok := m["num"];ok{
		fmt.Println(num)
	}else{
		fmt.Println("key not exist")
	}

	if site,ok := m["site"];ok{
		delete(m,site)
	}else{
		fmt.Println("key not exist")
	}
	fmt.Println(m)
}
