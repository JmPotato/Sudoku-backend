package generator

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

//Generate generates a question with unique solution
func Generate(level int) string {
	board := make(chan [9][9]byte, 10)
	for i := 0; i < 200; i++ {
		go func() {
			defer func() {
				recover()
			}()
			s, err := LasVegas(11)
			s.DigHole(81 - level)
			if err == nil {
				board <- s.board
			}
		}()
	}

	result := ""
	boardResult := <-board
	close(board)
	for _, row := range boardResult {
		for _, num := range row {
			if rune(num) == '.' {
				result += "0"
			} else {
				result += strconv.Itoa(int(num))
			}
		}
	}
	return result
}

//LasVegas algorithm generates a Sudoku solution, 99% success rate
func LasVegas(n int) (*Sudoku, error) {
	var b [9][9]byte
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b[i][j] = '.'
		}
	}
	s := NewSudoku(b)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; {
		x := rand.Intn(9)
		y := rand.Intn(9)
		if '.' == s.board[x][y] {
			val := byte(rand.Intn(9) + 1)
			if s.CanInsert(x, y, val) {
				s.Insert(x, y, val)
				i++
			}
		}
	}

	if s.SolveSudoku() {
		return s, nil
	}
	return s, errors.New("Unsolvable")
}

//DigHole dig holes on board
func (s *Sudoku) DigHole(holes int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < holes; {
		x := rand.Intn(9)
		y := rand.Intn(9)
		if s.board[x][y] != '.' {
			s.Remove(x, y)
			i++
		}
	}
}

/*CheckUnique check whether has unique solution after dig out (x, y)
By replacing the number on (x, y), if it has solution still, return false*/
// func (s *Sudoku) CheckUnique(x, y int) bool {
// 	temp := DeepCopy(s)
// 	temp.Remove(x, y)
// 	for _, val := range []byte{1, 2, 3, 4, 5, 6, 7, 8, 9} {
// 		if val != s.board[x][y] && temp.CanInsert(x, y, val) {
// 			temp.Insert(x, y, byte(val))
// 			if temp.SolveSudoku() {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// DeepCopy the Sudoku
// func DeepCopy(s *Sudoku) *Sudoku {
// 	copy := new(Sudoku)
// 	copy.board = s.board
// 	copy.rowMap = s.rowMap
// 	copy.colMap = s.rowMap
// 	copy.boxMap = s.boxMap
// 	copy.solved = s.solved
// 	return copy
// }
