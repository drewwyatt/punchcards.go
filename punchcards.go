// [0]      ________________________________________________________________
// [1]     /&-0123456789ABCDEFGHIJKLMNOPQR/STUVWXYZ:#@'="[.<(+^!$*);\],%_>?
// [2] 12 / O           OOOOOOOOO                        OOOOOO
// [3] 11|   O                   OOOOOOOOO                     OOOOOO
// [4]  0|    O                           OOOOOOOOO                  OOOOOO
// [5]  1|     O        O        O        O
// [6]  2|      O        O        O        O       O     O     O     O
// [7]  3|       O        O        O        O       O     O     O     O
// [8]  4|        O        O        O        O       O     O     O     O
// [9]  5|         O        O        O        O       O     O     O     O
// [10] 6|          O        O        O        O       O     O     O     O
// [11] 7|           O        O        O        O       O     O     O     O
// [12] 8|            O        O        O        O OOOOOOOOOOOOOOOOOOOOOOOO
// [13] 9|             O        O        O        O
// [14]  |__________________________________________________________________

package main

import (
	"fmt"
	"strings"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	idxs := [][]int{
		[]int{2},
		[]int{3},
		[]int{4},
		[]int{5},
		[]int{6},
		[]int{7},
		[]int{8},
		[]int{9},
		[]int{10},
		[]int{11},
		[]int{12},
		[]int{13},
		[]int{2, 5},
		[]int{2, 6},
		[]int{2, 7},
		[]int{2, 8},
		[]int{2, 9},
		[]int{2, 10},
		[]int{2, 11},
		[]int{2, 12},
		[]int{2, 13},
		[]int{3, 5},
		[]int{3, 6},
		[]int{3, 7},
		[]int{3, 8},
		[]int{3, 9},
		[]int{3, 10},
		[]int{3, 11},
		[]int{3, 12},
		[]int{3, 13},
		[]int{4, 5},
		[]int{4, 6},
		[]int{4, 7},
		[]int{4, 8},
		[]int{4, 9},
		[]int{4, 10},
		[]int{4, 11},
		[]int{4, 12},
		[]int{4, 13},

		[]int{12, 6},
		[]int{12, 7},
		[]int{12, 8},
		[]int{12, 9},
		[]int{12, 10},
		[]int{12, 11},
		[]int{12, 6, 2},
		[]int{12, 7, 2},
		[]int{12, 8, 2},
		[]int{12, 9, 2},
		[]int{12, 10, 2},
		[]int{12, 11, 2},
		[]int{12, 6, 3},
		[]int{12, 7, 3},
		[]int{12, 8, 3},
		[]int{12, 9, 3},
		[]int{12, 10, 3},
		[]int{12, 11, 3},
		[]int{12, 6, 4},
		[]int{12, 7, 4},
		[]int{12, 8, 4},
		[]int{12, 9, 4},
		[]int{12, 10, 4},
		[]int{12, 11, 4},
	}

	reserved := []int{0, 1, 14}

	template := []string{
		"     ",
		"    /",
		"12 / ",
		"11|  ",
		" 0|  ",
		" 1|  ",
		" 2|  ",
		" 3|  ",
		" 4|  ",
		" 5|  ",
		" 6|  ",
		" 7|  ",
		" 8|  ",
		" 9|  ",
		"  |__",
	}

	alphabet := "&-0123456789ABCDEFGHIJKLMNOPQR/STUVWXYZ:#@'=\"[.<(+^!$*);\\],%_>?"
	m := make(map[string][]int)
	for i := 0; i < len(alphabet); i++ {
		m[string(alphabet[i])] = idxs[i]
	}

	// rawInput := "3"
	// rawInput := "Hello, world!"
	rawInput := "This is Reddit's r/dailyprogrammer challenge."
	// rawInput := "WRITE (6,7) FORMAT(13H HELLO, WORLD) STOP END"
	input := strings.ToUpper(rawInput)

	for i := 0; i < len(input); i++ {
		letter := string(input[i])
		indexes := m[letter]

		fmt.Println(m[letter])

		// add punches
		for _, idx := range indexes {
			template[idx] += "0"
		}

		template[0] += "-"
		template[1] += letter
		template[14] += "_"

		// add spaces
		for j := range template {
			if !contains(indexes, j) && !contains(reserved, j) {
				template[j] += " "
			}
		}
	}

	for _, v := range template {
		fmt.Println(v)
	}
}
