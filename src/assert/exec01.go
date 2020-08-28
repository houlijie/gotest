package main

import "fmt"

// 断言实例2

func typeCheck(items ...interface{}) {
	for index, v := range items {
		// fmt.Println(v.(type)) // error  use of .(type) outside type switch
		switch v.(type) {
		case bool:
			fmt.Printf("index=%d, v=%v, type=bool", index, v)
		}
	}
}
func main() {
	n1 := 1
	n2 := 2.3
	n3 := false
	typeCheck(n1, n2, n3)
}
