package generator

//Sudoku quastion board
type Sudoku struct {
	board  [9][9]byte
	rowMap [90]bool  //rowMap[x*9+val] == true means row(x) exists number val
	colMap [90]bool  //colMap[y*9+val] == true means column(y) exists number val
	boxMap [230]bool //boxMap[(i/3)*100+(j/3)*10+val] == true means the box contain (i,j)exist number val
	solved bool
}

//NewSudoku initial a board
func NewSudoku(board [9][9]byte) *Sudoku {
	s := new(Sudoku)
	s.board = board
	for i, row := range board {
		for j, val := range row {
			if val != '.' {
				s.Insert(i, j, val)
			}
		}
	}
	return s
}

//SolveSudoku accept a 9*9 byte array and solve it
func (s *Sudoku) SolveSudoku() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.board[i][j] == '.' {
				s.BackTrace(i, j)
				if s.solved {
					return true
				}
			}
		}
	}
	return false
}

//CanInsert judge whether number inp can be inserted into (x, y)
func (s *Sudoku) CanInsert(x, y int, inp byte) bool {
	if s.board[x][y] != '.' {
		return false
	}
	val := int(inp)
	if s.rowMap[x*10+val] || s.colMap[y*10+val] || s.boxMap[(x/3)*100+(y/3)*10+val] {
		return false
	}
	return true
}

//Insert number inp into (x, y)
func (s *Sudoku) Insert(x, y int, inp byte) {
	s.board[x][y] = inp
	val := int(inp)
	s.rowMap[x*10+val] = true
	s.colMap[y*10+val] = true
	s.boxMap[(x/3)*100+(y/3)*10+val] = true
}

//Remove number inp from (x, y)
func (s *Sudoku) Remove(x, y int) {
	val := int(s.board[x][y])
	s.board[x][y] = '.'
	s.rowMap[x*10+val] = false
	s.colMap[y*10+val] = false
	s.boxMap[(x/3)*100+(y/3)*10+val] = false
}

//BackTrace use deep first search
func (s *Sudoku) BackTrace(x, y int) {
	println("back trace", x, y)
	if y == 9 {
		x = x + 1
		y = 0
	}
	if x == 9 {
		s.solved = true
		return
	}
	if s.board[x][y] == '.' {
		for opt := 1; opt < 9; opt++ {
			val := byte(opt)
			if !s.CanInsert(x, y, val) {
				continue
			}
			s.Insert(x, y, val)
			s.BackTrace(x, y+1)
			if s.solved {
				return
			}
			s.Remove(x, y)
		}
	} else {
		s.BackTrace(x, y+1)
	}
}