package main

import "fmt"

func f1() {
	fmt.Println("F1 function")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "retorno1", "retorno2"
}

func main() {
	f1()
	f2("param1", "param2")

	r3, r4 := f3(), f4("param1", "param2")
	fmt.Println(r3)
	fmt.Println(r4)
	fmt.Println(f5())
}
