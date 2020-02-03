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

// функция Sqrt2(x) не работает
// сделать, чтобы она работала

// func Sqrt2(x float64) float64  {
//     l := 0.0
//     r := float64(1e100)
//     var m float64
//
//     for r - l > 1e-8 {
//         m = l + (r - 1)/2
//         if m*m > n {
//             l = m
//         } else {
//             r = m
//         }
//     }
//     return l
// }

// func Sqrt2(x float64) float64  {
//     x := 1.0
//     for {
//         var nx float64 = (x + n/2)
//     }
// }

func main()  {
    fmt.Println("Квадратный корень по Ньютону",Sqrt(2))
    // fmt.Println("Методом половинного деления",Sqrt2(2))
    fmt.Println("Значение функции math.Sqrt(2)",math.Sqrt(2))
}
