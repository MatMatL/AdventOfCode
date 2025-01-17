package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	keys, locks := ReadFile("input.txt")

	fmt.Println("keys : ", keys, "locks : ", locks)

	Part1(keys, locks)
	//Part2(input)
}

func ReadFile(filename string) ([][]int, [][]int) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var keys [][]int = [][]int{}
	var locks [][]int = [][]int{}

	currentObject := []string{}

	for scanner.Scan() {
		if scanner.Text() != "" {
			currentObject = append(currentObject, scanner.Text())
		} else {
			if currentObject[6] == "#####" {
				var key []int = []int{0, 0, 0, 0, 0}
				for i := 0; i < len(currentObject)-1; i++ {
					for i, char := range currentObject[i] {
						if char == '#' {
							key[i]++
						}
					}
				}
				keys = append(keys, key)
			} else if currentObject[0] == "#####" {
				var lock []int = []int{0, 0, 0, 0, 0}
				for i := 1; i < len(currentObject); i++ {
					for i, char := range currentObject[i] {
						if char == '#' {
							lock[i]++
						}
					}
				}
				locks = append(locks, lock)
			}
			currentObject = []string{}
		}
	}

	return keys, locks
}

// ####################################################################################################################################
// ####################################################################################################################################
// ########################################                  PART 1                   #################################################
// ####################################################################################################################################
// ####################################################################################################################################

func Part1(keys [][]int, locks [][]int) {
	var result int = 0

	for i := 0; i < len(locks); i++ {
		for j := 0; j < len(keys); j++ {
			result += TryUnlock(locks[i], keys[j])
		}
	}

	fmt.Println("Resultat part 1 : ", result)
}

func TryUnlock(lock []int, key []int) int {
	for i := 0; i < len(lock); i++ {
		if lock[i]+key[i] > 5 {
			return 0
		}
	}
	return 1
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
