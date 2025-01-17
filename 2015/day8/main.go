package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := ReadFile("input.txt")

	Part1(input)
}

func ReadFile(filename string) [][]rune {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result [][]rune = [][]rune{}

	for scanner.Scan() {
		tempo := []rune{}
		for _, char := range scanner.Text() {
			tempo = append(tempo, rune(char))
		}
		result = append(result, tempo)
	}

	return result
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 1                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part1(input [][]rune) {
	resultTab := FindProblemes(input)

	result := Count(resultTab)

	PrintTab(resultTab)
	fmt.Println("Resultat part 1 : ", result)
}

func FindProblemes(input [][]rune) [][]bool {
	var result [][]bool = make([][]bool, len(input))
	for i := range result {
		result[i] = make([]bool, len(input[0]))
	}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] != '.' {
				fmt.Println(" found : ", input[i][j], "at", i, j)
				Search(input, i, j, result)
			}
		}
	}

	return result
}

func Search(input [][]rune, i int, j int, result [][]bool) {
	xDistance := 0
	yDistance := 0

	for k := i + 1; k < len(input); k++ {
		for l := 0; l < len(input[0]); l++ {
			if input[i][j] == input[k][l] {
				xDistance = k - i
				yDistance = l - j
				for w := 0; i-(xDistance*w) >= 0 && j-(yDistance*w) >= 0 && i-(xDistance*w) < len(input) && j-(yDistance*w) < len(input[0]); w++ {
					result[i-(xDistance*w)][j-(yDistance*w)] = true
				}
				for z := 0; k+(xDistance*z) >= 0 && l+(yDistance*z) >= 0 && k+(xDistance*z) < len(input) && l+(yDistance*z) < len(input[0]); z++ {
					result[k+(xDistance*z)][l+(yDistance*z)] = true
				}
			}
		}
	}
}

func Count(input [][]bool) int {
	result := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] {
				result++
			}
		}
	}
	return result
}

func PrintTab(input [][]bool) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if input[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(input [][]rune) {
	var result int = 0

	fmt.Println("Resultat part 2 : ", result)
}
