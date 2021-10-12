package main

func main() {
	a1 := 1
	a2 := 1
	println(a1)
	for i := 0; i < 10; i++ {
		println(a2)
		tmp := a2
		a2 = a1 + a2
		a1 = tmp
	}
}
