package main

import (
	"fmt"
	"math"
)

var a, b float64
var m int64

func go_f(x float64) float64 {
	// fmt.Println(math.Pow(float64(1)+(float64(2)/math.Pi)*math.Abs(math.Atan(x)), 0.5))
	return math.Pow(float64(1)+(float64(2)/math.Pi)*math.Abs(math.Atan(x)), 0.5)
}
func my_f(x, eps float64) float64 {
	x = (math.Abs(arctg(x, eps))) * 2 / math.Pi
	result := float64(1)
	k := 1
	step := result
	for step >= eps {
		step *= -2 * float64(k) * x * (float64(3-2*k) / float64(4*math.Pow(float64(k), 2)*
			float64(1-2*k)))
		result += step
		k++
	}
	return result
}
func arctg(x, eps float64) float64 {
	new_x := x
	if math.Abs(x) > 1 {
		new_x = 1.0 / x
	}
	result := new_x
	k := 2
	step := x
	for step >= eps {
		step *= -1 * math.Pow(new_x, 2) * float64(2*k-3) / float64(2*k-1)
		k++
		result += step
	}
	if math.Abs(x) > 1 {
		return math.Pi/2 - result
	}
	return result
}
func main() {
	var eps = 0.0000001
	fmt.Println("Введите зачения a, b и m\nНе забудьте про ограничения на переменные из отчёта")
	fmt.Scan(&a, &b, &m)
	h := (b - a) / float64(m)
	for i := 0; int64(i) <= m; i++ {
		x := float64(a + float64(i)*h)
		// fmt.Println(x, " ", my_f(x, eps), " ", go_f(x))
		fmt.Println(x, arctg(x, eps), math.Atan(x))
	}
}
