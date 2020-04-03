package myLib

func Add(x, y int) int {
	return x + y
}

func add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	a := add(x, y)
	println(a)
	return x - y
}
