package main

type MyNumber int

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Soma[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	m2 := map[string]int{"four": 4, "five": 5, "six": 6}
	m3 := map[string]int{"seven": 7, "eight": 8, "nine": 9}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10.0))
}
