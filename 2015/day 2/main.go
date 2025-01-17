package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := ReadFile("input.txt")

	Part1(input)
	Part2(input)

}

type Gift struct {
	Length int
	Width  int
	Height int
}

func ReadFile(filename string) []Gift {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result []Gift = []Gift{}

	for scanner.Scan() {
		var tempo Gift = Gift{}
		numbs := strings.Split(scanner.Text(), "x")
		tempo.Length, _ = strconv.Atoi(numbs[0])
		tempo.Width, _ = strconv.Atoi(numbs[1])
		tempo.Height, _ = strconv.Atoi(numbs[2])

		result = append(result, tempo)
	}

	return result
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 1                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part1(input []Gift) {

	result := Resolve(input)

	fmt.Println("Resultat part 1 : ", result)
}

func Resolve(input []Gift) int {
	var result int = 0

	for _, char := range input {
		min := char.Length * char.Width

		if char.Width*char.Height < min {
			min = char.Width * char.Height
		}
		if char.Height*char.Length < min {
			min = char.Height * char.Length
		}

		result += 2*char.Length*char.Width + 2*char.Width*char.Height + 2*char.Height*char.Length + min
	}
	return result
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(input []Gift) {
	result := Resolve2(input)

	println("Resultat part 2 : ", result)
}

func Resolve2(input []Gift) int {
	var result int = 0

	for _, char := range input {
		min := 2*char.Length + 2*char.Width
		per2 := 2*char.Width + 2*char.Height
		per3 := 2*char.Height + 2*char.Length

		if per2 < min {
			min = per2
		}
		if per3 < min {
			min = per3
		}

		result += char.Length*char.Width*char.Height + min
	}

	return result
}
