package main

import (
    "fmt"
    "strings"
)

func main()  {
    // Создание доски "крестики-нолики"
    board := [][]string{
        []string{"_","_","_"},
        []string{"_","_","_"},
        []string{"_","_","_"},
    }

    // Игроки делают ходы
    board[0][0] = "X"
    board[2][2] = "O"
    board[1][2] = "X"
    board[1][0] = "O"
    board[0][2] = "X"
    board[0][1] = "X"

    for i := 0; i < len(board); i++ {
        fmt.Printf("%s\n", strings.Join(board[i], " "))
    }
}
