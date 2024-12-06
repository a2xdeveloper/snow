package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"golang.org/x/exp/rand"
	"golang.org/x/term"
)

const (
	newSnowflakesPerTick = 25
	timePerTick = 250 * time.Millisecond
)

var (
	width  int
	height  int
	grid [][]rune
)

func init() {
	y, x, err := term.GetSize(int(os.Stdout.Fd()))
    if err != nil {
        fmt.Printf("Error getting terminal size: %v\n", err)
        return
    }
	width = y-1
	height = x-1


	grid = make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
	}
}

func main() {
	letItSnow()
}	

func letItSnow() {
	for i := 0; i < newSnowflakesPerTick; i++ {
		// spawn a snowflake	
		grid[0][rand.Intn(width)] = '*'
	}
	clearScreen()
	print()
	updateGrid()
	
	time.Sleep(timePerTick)
	letItSnow()
}

func updateGrid() {
	for i := len(grid)-2; i >= 0; i-- {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '*'  {
				grid[i][j] = ' '
				grid[i+1][j] = '*'
			}
		}
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func print() {
	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", cell)
			}
		}
		fmt.Println()
	}
}