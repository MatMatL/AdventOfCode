package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := ReadFile("input2.txt")

	// fmt.Print(input)

	Part1(input)
	Part2(input)
}

type coords struct {
	x int
	y int
}

func ReadFile(filename string) [][]int {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result [][]int = [][]int{}

	for scanner.Scan() {
		var temp []int = []int{}
		for _, numb := range scanner.Text() {
			temp = append(temp, int(numb-48))
		}
		result = append(result, temp)
	}

	return result
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 1                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part1(input [][]int) {
	var result int = 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 0 {
				var NineFound [][]bool = make([][]bool, len(input))
				for i := 0; i < len(input); i++ {
					NineFound[i] = make([]bool, len(input[i]))
				}

				for i := 0; i < len(NineFound); i++ {
					for j := 0; j < len(NineFound[i]); j++ {
						NineFound[i][j] = false
					}
				}

				result += TrySnake(i, j, input, -1, NineFound)
			}
		}
	}

	fmt.Println("Resultat part 1 : ", result)
}

func TrySnake(i int, j int, input [][]int, old int, NineFound [][]bool) int {
	if i < 0 || i >= len(input) || j < 0 || j >= len(input[i]) || input[i][j] != old+1 {
		return 0
	}

	if input[i][j] == 9 && !NineFound[i][j] {
		NineFound[i][j] = true
		return 1
	}
	return TrySnake(i+1, j, input, input[i][j], NineFound) + TrySnake(i-1, j, input, input[i][j], NineFound) + TrySnake(i, j+1, input, input[i][j], NineFound) + TrySnake(i, j-1, input, input[i][j], NineFound)
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(input [][]int) {
	var result int = 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 0 {
				result += TrySnake2(i, j, input, -1)
			}
		}
	}

	fmt.Println("Resultat part 1 : ", result)
}

func TrySnake2(i int, j int, input [][]int, old int) int {
	if i < 0 || i >= len(input) || j < 0 || j >= len(input[i]) || input[i][j] != old+1 {
		return 0
	}

	if input[i][j] == 9 {
		return 1
	}
	return TrySnake2(i+1, j, input, input[i][j]) + TrySnake2(i-1, j, input, input[i][j]) + TrySnake2(i, j+1, input, input[i][j]) + TrySnake2(i, j-1, input, input[i][j])
}
