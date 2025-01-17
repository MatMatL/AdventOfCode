package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := ReadFile("input.txt")

	Part1(input)
	Part2(input)

}

func ReadFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 1                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part1(input string) {

	result := Resolve(input)

	fmt.Println("Resultat part 1 : ", result)
}

func Resolve(input string) int {
	var OpenCount, CloseCount int = 0, 0

	for _, char := range input {
		if string(char) == "(" {
			OpenCount++
		} else if string(char) == ")" {
			CloseCount++
		}
	}
	return OpenCount - CloseCount
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(input string) {
	result := Resolve2(input)

	println("Resultat part 2 : ", result)
}

func Resolve2(input string) int {
	var OpenCount, CloseCount int = 0, 0

	for i, char := range input {
		if char == '(' {
			OpenCount++
		} else if char == ')' {
			CloseCount++
		}
		if OpenCount-CloseCount == -1 {
			return i
		}
	}
	return 0
}
