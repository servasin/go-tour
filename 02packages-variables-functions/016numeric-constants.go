package main

import "fmt"

const (
	// Создайте огромное число, сдвинув на 1 бит влево на 100 мест.
	// Другими словами, двоичное число, равное 1, сопровождается 100 нулями.
	Big = 1 << 1
	// Сдвиньте его еще раз на 99 мест, чтобы мы получили 1 << 1 или 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))

    fmt.Println(needInt(Big))
	fmt.Println(needFloat(Big))
}
