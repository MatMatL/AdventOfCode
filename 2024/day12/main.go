package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, checked := ReadFile("input2.txt")

	//Part1(input, checked)
	Part2(input, checked)
}

func ReadFile(filename string) ([][]rune, [][]bool) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result [][]rune = [][]rune{}
	var checked [][]bool = [][]bool{}

	for scanner.Scan() {
		tempo := []rune{}
		tempo2 := []bool{}
		for _, char := range scanner.Text() {
			tempo = append(tempo, rune(char))
			tempo2 = append(tempo2, false)
		}
		result = append(result, tempo)
		checked = append(checked, tempo2)
	}

	return result, checked
}

func PrintTab(tab [][]bool) {
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			fmt.Print(tab[i][j], " ")
		}
		fmt.Println()
	}
}

// ####################################################################################################################################
// ####################################################################################################################################
// ########################################                  PART 1                   #################################################
// ####################################################################################################################################
// ####################################################################################################################################

func Part1(input [][]rune, checked [][]bool) {
	var result int = 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if !checked[i][j] {
				peremiter, fences := FindPlot(input, i, j, checked, input[i][j])
				result += peremiter * fences
			}
		}
	}

	fmt.Println("Resultat part 1 : ", result)
}

func FindPlot(input [][]rune, i int, j int, checked [][]bool, char rune) (int, int) {
	if i < 0 || j < 0 || i >= len(input) || j >= len(input[0]) || (checked[i][j] == true && input[i][j] != char) || input[i][j] != char {
		return 0, 1
	}

	if checked[i][j] == true {
		return 0, 0
	}

	checked[i][j] = true

	peremiter, fences := FindPlot(input, i+1, j, checked, char)
	peremiter2, fences2 := FindPlot(input, i-1, j, checked, char)
	peremiter3, fences3 := FindPlot(input, i, j+1, checked, char)
	peremiter4, fences4 := FindPlot(input, i, j-1, checked, char)

	return peremiter + peremiter2 + peremiter3 + peremiter4 + 1, fences + fences2 + fences3 + fences4
}

// ####################################################################################################################################
// ####################################################################################################################################
// ########################################                  PART 2                   #################################################
// ####################################################################################################################################
// ####################################################################################################################################

func Part2(input [][]rune, checked [][]bool) {
	var result int = 0

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			if !checked[i][j] {
				var currentShape [][]bool = initBool(input)
				peremiter, _ := FindPlot2(input, i, j, checked, input[i][j], currentShape)
				fences := CalCulFences(currentShape)
				result += peremiter * fences
				PrintTab(currentShape)
				fmt.Println("peremiter : ", peremiter, " fences : ", fences)
			}
		}
	}

	fmt.Println("Resultat part 2 : ", result)
}

func FindPlot2(input [][]rune, i int, j int, checked [][]bool, char rune, currentShape [][]bool) (int, int) {
	if i < 0 || j < 0 || i >= len(input) || j >= len(input[0]) || (checked[i][j] == true && input[i][j] != char) || input[i][j] != char {
		return 0, 1
	}

	if checked[i][j] == true {
		return 0, 0
	}

	checked[i][j] = true
	currentShape[i][j] = true

	peremiter, fences := FindPlot2(input, i+1, j, checked, char, currentShape)
	peremiter2, fences2 := FindPlot2(input, i-1, j, checked, char, currentShape)
	peremiter3, fences3 := FindPlot2(input, i, j+1, checked, char, currentShape)
	peremiter4, fences4 := FindPlot2(input, i, j-1, checked, char, currentShape)

	return peremiter + peremiter2 + peremiter3 + peremiter4 + 1, fences + fences2 + fences3 + fences4
}

func initBool(input [][]rune) [][]bool {
	var result [][]bool = [][]bool{}

	for i := 0; i < len(input); i++ {
		tempo := []bool{}
		for j := 0; j < len(input[0]); j++ {
			tempo = append(tempo, false)
		}
		result = append(result, tempo)
	}

	return result
}

func CalCulFences(currentShape [][]bool) int {
	var result int = 0

	for k := 0; k < len(currentShape); k++ {
		for l := 0; l < len(currentShape[0]); l++ {
			if currentShape[k][l] {
				// Check left side
				if k > 0 && l > 0 {
					if currentShape[k-1][l] == false && currentShape[k][l-1] == false {
						result++
					} else if l > 0 && currentShape[k-1][l-1] == true && currentShape[k][l-1] == false {
						result++
					}
				} else if k == 0 {
					result++
				}

				// Check up side
				if k > 0 && l > 0 {
					if currentShape[k-1][l] == false && currentShape[k][l-1] == false {
						result++
					} else if l > 0 && currentShape[k-1][l-1] == true && currentShape[k-1][l] == false {
						result++

					}
				} else {
					result++
				}

				// Check right side
				if k > 0 && l+1 < len(currentShape[0]) {
					if currentShape[k-1][l] == false && currentShape[k][l+1] == false {
						result++
					} else if currentShape[k-1][l-1] == true && currentShape[k-1][l] == false {
						result++

					}
				} else {
					result++
				}

				// Check down side
				if l > 0 {
					if currentShape[k][l-1] == false {
						result++
					} else if k+1 < len(currentShape) && currentShape[k+1][l-1] == true {
						result++

					}
				} else {
					result++
				}
			}
		}
	}
	return result
}
