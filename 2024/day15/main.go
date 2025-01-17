package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	Map, Directions, player := ReadFile("input4.txt")

	fmt.Println("Map : ", Map, "\nDirections : ", Directions)

	//Part1(Map, Directions, player)
	Part2(Map, Directions, player)
}

type Player struct {
	x int
	y int
}

func ReadFile(filename string) ([][]rune, []int, Player) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return nil, nil, Player{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var Map [][]rune = [][]rune{}
	var Directions []int = []int{}
	var player Player = Player{}

	MapDone := false

	for scanner.Scan() {
		if scanner.Text() == "" {
			MapDone = true
		} else if !MapDone {
			var newLine []rune = []rune{}
			for i, char := range scanner.Text() {
				if char == '@' {
					player.x = i
					player.y = len(Map)
				}
				newLine = append(newLine, char)
			}
			Map = append(Map, newLine)
		} else {
			for _, direction := range scanner.Text() {
				switch direction {
				case '^':
					Directions = append(Directions, 0)
				case '>':
					Directions = append(Directions, 1)
				case 'v':
					Directions = append(Directions, 2)
				case '<':
					Directions = append(Directions, 3)
				}
			}
		}
	}
	return Map, Directions, player
}

// ####################################################################################################################################
// ####################################################################################################################################
// ########################################                  PART 1                   #################################################
// ####################################################################################################################################
// ####################################################################################################################################

func Part1(Map [][]rune, Directions []int, player Player) {
	var result int = 0
	for i := 0; i < len(Directions); i++ {
		fmt.Println(Directions[i])
		if TryMove(Map, Directions[i], player.x, player.y) {
			switch Directions[i] {
			case 0:
				player.y--
			case 1:
				player.x++
			case 2:
				player.y++
			case 3:
				player.x--
			}
		}
	}

	PrintMap(Map)

	result = CalculeResult(Map)

	fmt.Println("Resultat part 1 : ", result)
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func TryMove(Map [][]rune, Direction int, objectX int, objectY int) bool {
	switch Direction {
	case 0:
		if Map[objectY-1][objectX] == '.' {
			Map[objectY-1][objectX] = Map[objectY][objectX]
			Map[objectY][objectX] = '.'
			return true
		} else if Map[objectY-1][objectX] == 'O' {
			if TryMove(Map, Direction, objectX, objectY-1) {
				Map[objectY-1][objectX] = Map[objectY][objectX]
				Map[objectY][objectX] = '.'
				return true
			}
		} else if Map[objectY-1][objectX] == '#' {
			return false
		}
	case 1:
		if Map[objectY][objectX+1] == '.' {
			Map[objectY][objectX+1] = Map[objectY][objectX]
			Map[objectY][objectX] = '.'
			return true
		} else if Map[objectY][objectX+1] == 'O' {
			if TryMove(Map, Direction, objectX+1, objectY) {
				Map[objectY][objectX+1] = Map[objectY][objectX]
				Map[objectY][objectX] = '.'
				return true
			}
		} else if Map[objectY+1][objectX] == '#' {
			return false
		}
	case 2:
		if Map[objectY+1][objectX] == '.' {
			Map[objectY+1][objectX] = Map[objectY][objectX]
			Map[objectY][objectX] = '.'
			return true
		} else if Map[objectY+1][objectX] == 'O' {
			if TryMove(Map, Direction, objectX, objectY+1) {
				Map[objectY+1][objectX] = Map[objectY][objectX]
				Map[objectY][objectX] = '.'
				return true
			}
		} else if Map[objectY+1][objectX] == '#' {
			return false
		}
	case 3:
		if Map[objectY][objectX-1] == '.' {
			Map[objectY][objectX-1] = Map[objectY][objectX]
			Map[objectY][objectX] = '.'
			return true
		} else if Map[objectY][objectX-1] == 'O' {
			if TryMove(Map, Direction, objectX-1, objectY) {
				Map[objectY][objectX-1] = Map[objectY][objectX]
				Map[objectY][objectX] = '.'
				return true
			}
		} else if Map[objectY][objectX-1] == '#' {
			return false
		}
	}

	return false
}

func PrintMap(Map [][]rune) {
	for i := 0; i < len(Map); i++ {
		for j := 0; j < len(Map[0]); j++ {
			if Map[i][j] != '@' {
				fmt.Printf("\033[31m%c\033[0m", Map[i][j])
			} else {
				fmt.Print(string(Map[i][j]))
			}
		}
		fmt.Println()
	}
}

func CalculeResult(Map [][]rune) int {
	result := 0

	for i := 0; i < len(Map); i++ {
		for j := 0; j < len(Map[0]); j++ {
			if Map[i][j] == 'O' {
				result += i*100 + j
			}
		}
	}

	return result
}

//####################################################################################################################################
//####################################################################################################################################
//########################################                  PART 2                   #################################################
//####################################################################################################################################
//####################################################################################################################################

func Part2(Map [][]rune, Directions []int, player Player) {
	var result int = 0

	var newMap [][]rune = ReSizeMap(Map)
	player.x = player.x * 2

	for i := 0; i < len(Directions); i++ {
		fmt.Println("Direction : ", Directions[i], player.y, player.x)
		if TryTowMove(newMap, Directions[i], player.x, player.y) {
			switch Directions[i] {
			case 0:
				player.y--
			case 1:
				player.x++
			case 2:
				player.y++
			case 3:
				player.x--
			}
		}
		time.Sleep(750 * (time.Millisecond))
		clearTerminal()
		PrintMap(newMap)
	}

	PrintMap(newMap)

	result = CalculeTowResult(newMap)

	fmt.Println("Resultat part 2 : ", result)
}

func ReSizeMap(Map [][]rune) [][]rune {
	var newMap [][]rune = [][]rune{}

	for i := 0; i < len(Map); i++ {
		NewLine := []rune{}
		for j := 0; j < len(Map[0]); j++ {
			if Map[i][j] == '.' {
				NewLine = append(NewLine, '.', '.')
			} else if Map[i][j] == '#' {
				NewLine = append(NewLine, '#', '#')
			} else if Map[i][j] == 'O' {
				NewLine = append(NewLine, '[', ']')
			} else if Map[i][j] == '@' {
				NewLine = append(NewLine, '@', '.')
			}
		}
		newMap = append(newMap, NewLine)
	}

	return newMap
}

func TryTowMove(Map [][]rune, Direction int, objectX int, objectY int) bool {
	if Map[objectY][objectX] == '.' {
		return true
	}
	if Map[objectY][objectX] == '#' {
		return false
	}

	switch Direction {
	case 0:
		if Map[objectY][objectX] == '[' && TryTowMove(Map, Direction, objectX, objectY-1) && TryTowMove(Map, Direction, objectX+1, objectY-1) {
			Map[objectY-1][objectX], Map[objectY][objectX] = Map[objectY][objectX], Map[objectY-1][objectX]
			Map[objectY-1][objectX+1], Map[objectY][objectX+1] = Map[objectY][objectX+1], Map[objectY-1][objectX+1]
			return true
		} else if Map[objectY][objectX] == ']' && TryTowMove(Map, Direction, objectX, objectY-1) && TryTowMove(Map, Direction, objectX-1, objectY-1) {
			Map[objectY-1][objectX], Map[objectY][objectX] = Map[objectY][objectX], Map[objectY-1][objectX]
			Map[objectY-1][objectX-1], Map[objectY][objectX-1] = Map[objectY][objectX-1], Map[objectY-1][objectX-1]
			return true
		} else if Map[objectY][objectX] == '@' && TryTowMove(Map, Direction, objectX, objectY-1) {
			Map[objectY-1][objectX], Map[objectY][objectX] = Map[objectY][objectX], Map[objectY-1][objectX]
			return true
		}
	case 1:
		if TryTowMove(Map, Direction, objectX+1, objectY) {
			Map[objectY][objectX+1], Map[objectY][objectX] = Map[objectY][objectX], Map[objectY][objectX+1]
			return true
		}
	case 2:
		if Map[objectY][objectX] == '[' && TryTowMove(Map, Direction, objectX, objectY+1) && TryTowMove(Map, Direction, objectX+1, objectY+1) {
			Map[objectY+1][objectX], Map[objectY][objectX] = Map[objectY][objectX], Map[objectY+1][objectX]
			Map[objectY+1][objectX+1], Map[objectY][objectX+1] = Map[objectY][objectX+1], Map[objectY+1][objectX+1]
			return true
		} else if Map[objectY][objectX] == ']' && TryTowMove(Map, Direction, objectX, objectY+1) && TryTowMove(Map, Direction, objectX-1, objectY+1) {
			Map[objectY+1][objectX], Map[objectY][objectX] = Map[objectY][objectX], Map[objectY+1][objectX]
			Map[objectY+1][objectX-1], Map[objectY][objectX-1] = Map[objectY][objectX-1], Map[objectY+1][objectX-1]
			return true
		} else if Map[objectY][objectX] == '@' && TryTowMove(Map, Direction, objectX, objectY+1) {
			Map[objectY+1][objectX], Map[objectY][objectX] = Map[objectY][objectX], Map[objectY+1][objectX]
			return true
		}
	case 3:
		if TryTowMove(Map, Direction, objectX-1, objectY) {
			Map[objectY][objectX-1], Map[objectY][objectX] = Map[objectY][objectX], Map[objectY][objectX-1]
			return true
		}
	}

	return false
}

func CalculeTowResult(Map [][]rune) int {
	result := 0

	for i := 0; i < len(Map); i++ {
		for j := 0; j < len(Map[0]); j++ {
			if Map[i][j] == '[' {
				result += i*100 + j
			}
		}
	}

	return result
}
