package _go

import "fmt"

func subStrIndex(s string) int  {
	m := make(map[rune]int)
	start := 0
	maxLength := 0
	for i,ch := range []rune(s){
		if lastI,ok := m[ch];ok&&lastI>=start{
			start = lastI + 1
		}
		if i-start+1 > maxLength{
			maxLength = i-start+1
		}
		m[ch] = i
	}
	return maxLength
}
func main()  {
	fmt.Println(subStrIndex("bbbb"))
	fmt.Println(subStrIndex("abcabcbb"))
	fmt.Println(subStrIndex("我爱中国"))
	fmt.Println(subStrIndex("一二三二一"))
}
