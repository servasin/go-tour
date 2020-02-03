package main

import "fmt"

func main()  {
    c := 0.867 + 0.5i // тип c complex128
    fmt.Printf("тип c %T\n", c)

    i := 1          // тип i int
    fmt.Printf("тип i %T\n", i)

    f := 1.0       // тип f float64
    fmt.Printf("тип f %T\n", f)

    s := "s"       // тип s string
    fmt.Printf("тип s %T\n", s)
}
