package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64  {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

func main()  {
    fmt.Println(pow(3, 2, 10))
    fmt.Println(pow(3, 3, 20))
    fmt.Println(math.Pow(3, 3))

    fmt.Println(pow(2, 1, 10))
    fmt.Println(pow(2, 2, 10))

    fmt.Println(pow(4, 4, 10))
    fmt.Println(math.Pow(4, 4))

    fmt.Println(math.Pow(2, 2))
}
// степень - power
