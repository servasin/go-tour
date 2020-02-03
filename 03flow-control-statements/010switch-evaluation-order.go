package main

import (
    "fmt"
    "time"
)

func main()  {
    fmt.Println("Когда Воскресенье?")
    today := time.Now().Weekday()
    switch time.Saturday {
    case today + 0:
        fmt.Println("Сегодня!")
    case today + 1:
        fmt.Println("Завтра.")
    case today + 2:
        fmt.Println("Через 2 дня.")
    default:
        fmt.Println("Не скоро.")
    }
}
