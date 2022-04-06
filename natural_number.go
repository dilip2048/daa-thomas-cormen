package main

func Sum(n int) int {
	if n == 1 {
		return 1
	} else {
		return 1 + Sum(n-1)
	}
}

func main() {
	x := Sum(10)
	println(x)
}
