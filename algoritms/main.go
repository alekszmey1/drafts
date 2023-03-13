package main

import "fmt"

func main() {
	text := "giperboloid ingenera garina"

	fmt.Println(first(text))
	fmt.Println(second(text))
}

// поиск маскимального количества повторений символа в тексте
// O(n*n)
func first(s string) string {
	ans := ""
	ansCnt := 0
	for _, i2 := range s {
		nowCnt := 0
		for _, i4 := range s {
			if i2 == i4 {
				nowCnt++
			}
			if nowCnt > ansCnt {
				ans = string(i2)
				ansCnt = nowCnt
			}
		}
	}
	return ans
}

// O(n+k)
func second(s string) string {
	ans := ""
	ansCnt := 0
	m := make(map[string]int)
	for i := 0; i < len(s); i++ {
		m[string(s[i])]++
	}
	for sim, i := range m {
		if i > ansCnt {
			ansCnt = i
			ans = sim
		}
	}
	fmt.Println(m)
	return ans
}
