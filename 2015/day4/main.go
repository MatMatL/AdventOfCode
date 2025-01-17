package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := ReadFile("input.txt")

	Part2(input)
}

func ReadFile(filename string) []string {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []string = []string{}

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 1                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part1(input []string) {
	var result int = 0

	for x, line := range input {
		for y, char := range line {
			if char == 'X' {
				result += CheckFront(input, x, y)
				result += CheckBack(input, x, y)
				result += CheckUp(input, x, y)
				result += CheckDown(input, x, y)
				result += CheckDiagUpRight(input, x, y)
				result += CheckDiagUpLeft(input, x, y)
				result += CheckDiagDownRight(input, x, y)
				result += CheckDiagDownLeft(input, x, y)
				fmt.Println(result)
			}
		}
	}

	fmt.Println("Resultat part 1 : ", result)
}

func CheckFront(input []string, x int, y int) int {
	if len(input[x]) <= y+3 {
		return 0
	}

	if input[x][y+1] == 'M' && input[x][y+2] == 'A' && input[x][y+3] == 'S' {
		return 1
	}
	return 0
}

func CheckBack(input []string, x int, y int) int {
	if y-3 < 0 {
		return 0
	}

	if input[x][y-1] == 'M' && input[x][y-2] == 'A' && input[x][y-3] == 'S' {
		return 1
	}
	return 0
}

func CheckUp(input []string, x int, y int) int {
	if x-3 < 0 {
		return 0
	}

	if input[x-1][y] == 'M' && input[x-2][y] == 'A' && input[x-3][y] == 'S' {
		return 1
	}
	return 0
}

func CheckDown(input []string, x int, y int) int {
	if x+3 >= len(input) {
		return 0
	}

	if input[x+1][y] == 'M' && input[x+2][y] == 'A' && input[x+3][y] == 'S' {
		return 1
	}
	return 0
}

func CheckDiagUpRight(input []string, x int, y int) int {
	if (x-3 < 0) || (len(input[x]) <= y+3) {
		return 0
	}

	if input[x-1][y+1] == 'M' && input[x-2][y+2] == 'A' && input[x-3][y+3] == 'S' {
		return 1
	}
	return 0
}

func CheckDiagUpLeft(input []string, x int, y int) int {
	if (x-3 < 0) || (y-3 < 0) {
		return 0
	}

	if input[x-1][y-1] == 'M' && input[x-2][y-2] == 'A' && input[x-3][y-3] == 'S' {
		return 1
	}
	return 0
}

func CheckDiagDownRight(input []string, x int, y int) int {
	if (x+3 >= len(input)) || (len(input[x]) <= y+3) {
		return 0
	}

	if input[x+1][y+1] == 'M' && input[x+2][y+2] == 'A' && input[x+3][y+3] == 'S' {
		return 1
	}
	return 0
}

func CheckDiagDownLeft(input []string, x int, y int) int {
	if (x+3 >= len(input)) || (y-3 < 0) {
		return 0
	}

	if input[x+1][y-1] == 'M' && input[x+2][y-2] == 'A' && input[x+3][y-3] == 'S' {
		return 1
	}
	return 0
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(input []string) {
	var result int = 0

	for x, line := range input {
		for y, char := range line {
			if char == 'A' {
				result += Check2DiagUpRight(input, x, y)
				result += Check2DiagUpLeft(input, x, y)
				result += Check2DiagDownRight(input, x, y)
				result += Check2DiagDownLeft(input, x, y)
				fmt.Println(result)
			}
		}
	}

	fmt.Println("Resultat part 2 : ", result)
}

func Check2DiagUpRight(input []string, x int, y int) int {
	if (x-1 < 0) || (y-1 < 0) || (len(input[x]) <= x+1) || (len(input[y]) <= y+1) {
		return 0
	}

	if input[x-1][y-1] == 'M' && input[x+1][y+1] == 'S' && input[x-1][y+1] == 'M' && input[x+1][y-1] == 'S' {
		return 1
	}
	return 0
}

func Check2DiagUpLeft(input []string, x int, y int) int {
	if (x-1 < 0) || (y-1 < 0) || (len(input[x]) <= x+1) || (len(input[y]) <= y+1) {
		return 0
	}

	if input[x-1][y-1] == 'M' && input[x+1][y+1] == 'S' && input[x+1][y-1] == 'M' && input[x-1][y+1] == 'S' {
		return 1
	}
	return 0
}

func Check2DiagDownRight(input []string, x int, y int) int {
	if (x-1 < 0) || (y-1 < 0) || (len(input[x]) <= x+1) || (len(input[y]) <= y+1) {
		return 0
	}

	if input[x+1][y+1] == 'M' && input[x-1][y-1] == 'S' && input[x+1][y-1] == 'M' && input[x-1][y+1] == 'S' {
		return 1
	}
	return 0
}

func Check2DiagDownLeft(input []string, x int, y int) int {
	if (x-1 < 0) || (y-1 < 0) || (len(input[x]) <= x+1) || (len(input[y]) <= y+1) {
		return 0
	}

	if input[x+1][y+1] == 'M' && input[x-1][y-1] == 'S' && input[x-1][y+1] == 'M' && input[x+1][y-1] == 'S' {
		return 1
	}
	return 0
}
