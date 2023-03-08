package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var s []string
	var k int
	fmt.Scan(&k)
	for i := 0; i < k*2; i++ {
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		s = append(s, line)
	}

	for _, bytes := range s {
		fmt.Println(bytes)
	}
	/*
		var T []byte
		in := bufio.NewReader(os.Stdin)
		fmt.Fscan(in, &T)
		s := string(T)
		fmt.Println(s)
		// Scanner может просканировать построчный ввод
		/*sc := bufio.NewScanner(os.Stdin)
			for sc.Scan() {
				txt := sc.Text()
				fmt.Printf("Эхо: %s\n", txt)
			}


			for {
				data := make([]byte, 8)
				n, err := os.Stdin.Read(data)
				if err == nil && n > 0 {
					process(data)
				} else {
					break
				}
			}

		}

		func process(data []byte) {
			fmt.Printf("Получено: %X 	%s\n", data, string(data))
		}*/
}
