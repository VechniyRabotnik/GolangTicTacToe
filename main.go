package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Game struct {
	board  [3][3]string
	player string
}

func (g *Game) printBoard() {
	fmt.Println("  0 1 2")
	for i, row := range g.board {
		fmt.Print(i, "| ")
		for _, cell := range row {
			if cell == "" {
				fmt.Print(". ")
			} else {
				if cell == "X" {
					fmt.Print(color.HiRedString("X "))
				} else {
					fmt.Print(color.HiBlueString("O "))
				}
			}
		}
		fmt.Println("|")
	}
	fmt.Println(" ------------")
}

func (g *Game) checkWinner() string {
	for i := 0; i < 3; i++ {
		if g.board[i][0] != "" && g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] {
			return g.board[i][0]
		}
		if g.board[0][i] != "" && g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i] {
			return g.board[0][i]
		}
	}
	if g.board[0][0] != "" && g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] {
		return g.board[0][0]
	}
	if g.board[0][2] != "" && g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] {
		return g.board[0][2]
	}
	return ""
}

func (g *Game) isBoardFull() bool {
	for _, row := range g.board {
		for _, cell := range row {
			if cell == "" {
				return false
			}
		}
	}
	return true
}

func (g *Game) play() {
	var row, col int
	for {
		g.printBoard()
		fmt.Printf("Ход игрока %s. Введите номер строки и столбца (например, '0 1'): ", g.player)
		_, err := fmt.Scan(&row, &col)
		if err != nil || row < 0 || row > 2 || col < 0 || col > 2 || g.board[row][col] != "" {
			fmt.Println("Некорректный ввод, попробуйте снова.")
			fmt.Scanln()
			continue
		}
		g.board[row][col] = g.player

		if winner := g.checkWinner(); winner != "" {
			g.printBoard()
			fmt.Printf("Игрок %s победил!\n", winner)
			break
		} else if g.isBoardFull() {
			g.printBoard()
			fmt.Println("Игра окончилась вничью.")
			break
		}

		if g.player == "X" {
			g.player = "O"
		} else {
			g.player = "X"
		}
	}
}

func main() {
	color.Set(color.Bold)
	fmt.Println("Добро пожаловать в игру Крестики-Нолики!")
	color.Unset()
	game := Game{player: "X"}
	game.play()
}
