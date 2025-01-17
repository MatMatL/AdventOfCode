package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := ReadFile("exemple.txt")

	Part2(input)
}

func ReadFile(filename string) [][]string {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result [][]string = [][]string{}

	for scanner.Scan() {
		tempo := []string{}
		for _, char := range scanner.Text() {
			tempo = append(tempo, string(char))
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

func Part1(maps [][]string) {
	var result int = 1

	guardX, guardY := FindGuard(maps)

	var finish bool = false
	direction := "up"

	for !finish {
		fmt.Println(guardX, guardY)
		if direction == "up" {
			if guardX-1 < 0 || guardY > len(maps[0])-1 || guardY < 0 || guardX > len(maps)-1 {
				fmt.Println("out of bound")
				finish = true
				break
			} else if maps[guardX-1][guardY] == "#" {
				direction = "right"
			} else {
				if maps[guardX][guardY] != "X" {
					result++
					maps[guardX][guardY] = "X"
				}
				guardX--
			}
		} else if direction == "right" {
			if guardX < 0 || guardY+1 > len(maps[0])-1 || guardY < 0 || guardX > len(maps)-1 {
				finish = true
				break
			} else if maps[guardX][guardY+1] == "#" {
				direction = "down"
			} else {
				if maps[guardX][guardY] != "X" {
					result++
					maps[guardX][guardY] = "X"
				}
				guardY++
			}
		} else if direction == "down" {
			if guardX < 0 || guardY > len(maps[0])-1 || guardY < 0 || guardX+1 > len(maps)-1 {
				finish = true
				break
			} else if maps[guardX+1][guardY] == "#" {
				direction = "left"
			} else {
				if maps[guardX][guardY] != "X" {
					result++
					maps[guardX][guardY] = "X"
				}
				guardX++
			}
		} else if direction == "left" {
			if guardX < 0 || guardY > len(maps[0])-1 || guardY-1 < 0 || guardX > len(maps)-1 {
				finish = true
				break
			} else if maps[guardX][guardY-1] == "#" {
				direction = "up"
			} else {
				if maps[guardX][guardY] != "X" {
					result++
					maps[guardX][guardY] = "X"
				}
				guardY--
			}
		}
	}

	fmt.Println("Resultat part 1 : ", result)
}

func FindGuard(maps [][]string) (x, y int) {
	for x, line := range maps {
		for y, char := range line {
			if char == "^" {
				return x, y
			}
		}
	}
	return 0, 0
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(maps [][]string) {
	var result int = 1

	guardX, guardY := FindGuard(maps)

	var finish bool = false
	direction := "up"

	for !finish {
		fmt.Println(guardX, guardY)
		if direction == "up" {
			if guardX-1 < 0 || guardY > len(maps[0])-1 || guardY < 0 || guardX > len(maps)-1 {
				fmt.Println("out of bound")
				finish = true
				break
			} else if maps[guardX-1][guardY] == "#" {
				direction = "right"
			} else {
				if maps[guardX][guardY] != "X" {
					result++
					maps[guardX][guardY] = "X"
				}
				guardX--
			}
		} else if direction == "right" {
			if guardX < 0 || guardY+1 > len(maps[0])-1 || guardY < 0 || guardX > len(maps)-1 {
				finish = true
				break
			} else if maps[guardX][guardY+1] == "#" {
				direction = "down"
			} else {
				if maps[guardX][guardY] != "X" {
					result++
					maps[guardX][guardY] = "X"
				}
				guardY++
			}
		} else if direction == "down" {
			if guardX < 0 || guardY > len(maps[0])-1 || guardY < 0 || guardX+1 > len(maps)-1 {
				finish = true
				break
			} else if maps[guardX+1][guardY] == "#" {
				direction = "left"
			} else {
				if maps[guardX][guardY] != "X" {
					result++
					maps[guardX][guardY] = "X"
				}
				guardX++
			}
		} else if direction == "left" {
			if guardX < 0 || guardY > len(maps[0])-1 || guardY-1 < 0 || guardX > len(maps)-1 {
				finish = true
				break
			} else if maps[guardX][guardY-1] == "#" {
				direction = "up"
			} else {
				if maps[guardX][guardY] != "X" {
					result++
					maps[guardX][guardY] = "X"
				}
				guardY--
			}
		}
	}

	fmt.Println("Resultat part 2 : ", result)
}

func printMap(tab [][]string) {
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			fmt.Print(tab[i][j], " ")
		}
		fmt.Println()
	}
}
