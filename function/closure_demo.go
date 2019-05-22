package funcstudy

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	f := closure(10)
	result := f(10)
}
