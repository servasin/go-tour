package main

import "fmt"

func main()  {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    // Даем слайсу нулевую длину
    s = s[:0]
    printSlice(s)

    // Увеличиваем длину
    s = s[:4]
    printSlice(s)

    // Сбрасываем первые два значения
    s = s[2:]
    printSlice(s)
}

func printSlice(s[]int)  {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
