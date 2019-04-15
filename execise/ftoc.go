package exe

import "fmt"

func main() {
	var strss = " "
	fmt.Println(strss)
	var i, j = 2, 3
	i, j = j, i
	fmt.Println(i, j)
	var x, y int
	fmt.Println(&x == &y, &x == &x, &x == nil)
	xx := 1
	k := &xx
	fmt.Println(xx)
	*k = 3
	fmt.Println(*k)
	/*const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))*/
}

func inc(p *int) int {
	*p++
	return *p
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
