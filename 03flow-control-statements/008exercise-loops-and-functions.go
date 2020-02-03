package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64  {
    z       := 1.0
    check   := 0.0

    for i := 0; i <= 100; i++{
        z = z - (z*z - x) / (2*z)
        fmt.Println("Попытка", i, "Квадратный корень равен =", z)
        if i != 0 {
            if check - z < 1e-8 {
                fmt.Println("Это значение довольно точное!")
                break
            }
        }
        check = z
    }
    return z
}

func main()  {
    fmt.Println("Квадратный корень по Ньютону",Sqrt(2))
    fmt.Println("Значение функции math.Sqrt(2)",math.Sqrt(2))
}
