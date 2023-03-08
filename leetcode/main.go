package main

import (
	"fmt"
)

func main() {
	b := isValid("()[]{}")
	fmt.Println(b)
}
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0, len(s))
	hashMap := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	for _, char := range s {
		fmt.Println("взяли " + string(char))
		if _, ok := hashMap[char]; !ok {
			fmt.Println("найдено совпадение")
			fmt.Println(string(hashMap[char]))
			// this bracket is open
			stack = append(stack, char)
			fmt.Println(stack)
		} else {
			fmt.Println("сработало иначе")
			// this bracket is close
			if len(stack) == 0 || (stack[len(stack)-1] != hashMap[char]) {
				return false
			}
			stack = stack[:len(stack)-1]
			fmt.Println(stack)
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
