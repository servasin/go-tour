package main

import "fmt"

const Pi = 3.14

func main()  {
    const World = "世界"
//  World = "World"   нельзя поменять значение - это константа
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)
//  Константы не могут быть объявлены с помощью синтаксиса :=
}
