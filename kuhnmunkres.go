package kuhnMunkres

import (
	"math"
)

// Kuhn-Munkres algorithm for the Assignment Problem with an incomplete cost matrix.
// 带权二分图最大匹配
type Munkres struct {
	maxtrix   [][]int // 矩阵
	started   [][]int // 返回结果
	path      [][]int
	rowCover  []int
	colCover  []int
	nrow      int // 行数
	ncol      int // 列数
	pathCount int
	pathRow0  int
	pathCol0  int
	step      int
}

func NewMunkres(matrix [][]int) *Munkres {
	nrow := len(matrix)
	ncol := 0
	if nrow >= 1 {
		ncol = len(matrix[0])
	}
	m := &Munkres{
		maxtrix:  matrix,
		nrow:     nrow,
		ncol:     ncol,
		started:  make([][]int, nrow),
		rowCover: make([]int, nrow),
		colCover: make([]int, ncol),
		step:     1,
	}

	m.initBidArray()
	return m
}

func (self *Munkres) initBidArray() {
	m := make([][]int, self.nrow)
	for r := 0; r < self.nrow; r++ {
		c := make([]int, self.ncol)
		m[r] = c
	}
	self.started = m

	t := self.nrow + self.ncol + 1
	m = make([][]int, t)
	for i := 0; i < t; i++ {
		c := make([]int, 2)
		m[i] = c
	}
	self.path = m
}

//For each row of the cost matrix, find the smallest element and subtract
//it from every element in its row.  When finished, Go to Step 2.
func (self *Munkres) step1() {
	var min int
	for r, row := range self.maxtrix {
		min = self.maxtrix[r][0]
		for _, col := range row {
			if col < min {
				min = col
			}
		}
		for c, _ := range row {
			self.maxtrix[r][c] -= min
		}
	}
	self.step = 2
}

//Find a zero (Z) in the resulting matrix.  If there is no starred
//zero in its row or column, star Z. Repeat for each element in the
//matrix. Go to Step 3.
func (self *Munkres) step2() {
	for r, row := range self.maxtrix {
		for c, col := range row {
			if col == 0 && self.rowCover[r] == 0 && self.colCover[c] == 0 {
				self.started[r][c] = 1
				self.rowCover[r] = 1
				self.colCover[c] = 1
			}
		}
	}
	for r := 0; r < self.nrow; r++ {
		self.rowCover[r] = 0
	}
	for c := 0; c < self.ncol; c++ {
		self.colCover[c] = 0
	}
	self.step = 3
}

//Cover each column containing a starred zero.  If K columns are covered,
//the starred zeros describe a complete set of unique assignments.  In this
//case, Go to DONE, otherwise, Go to Step 4.
func (self *Munkres) step3() {
	for _, row := range self.started {
		for c, col := range row {
			if col == 1 {
				self.colCover[c] = 1
			}
		}
	}

	colCount := 0
	for c := 0; c < self.ncol; c++ {
		if self.colCover[c] == 1 {
			colCount++
		}
	}

	if colCount >= self.ncol || colCount >= self.nrow {
		self.step = 7
	} else {
		self.step = 4
	}
}

// FindAZero setp4用到的方法
func (self *Munkres) findAZero() (int, int) {
	rw, cl := -1, -1
	for r, row := range self.maxtrix {
		for c, col := range row {
			if col == 0 && self.rowCover[r] == 0 && self.colCover[c] == 0 {
				return r, c
			}
		}
	}
	return rw, cl
}

func (self *Munkres) findStarInRow(row int) int {
	cl := -1
	rw := self.started[row]
	for c, col := range rw {
		if col == 1 {
			cl = c
		}
	}
	return cl
}

//Find a noncovered zero and prime it.  If there is no starred zero
//in the row containing this primed zero, Go to Step 5.  Otherwise,
//cover this row and uncover the column containing the starred zero.
//Continue in this manner until there are no uncovered zeros left.
//Save the smallest uncovered value and Go to Step 6.
func (self *Munkres) step4() {
	row, col := -1, -1
	done := false
	for {
		if done {
			break
		}

		row, col = self.findAZero()
		if row == -1 {
			done = true
			self.step = 6
		} else {
			self.started[row][col] = 2
			if c := self.findStarInRow(row); c > -1 {
				self.rowCover[row] = 1
				self.colCover[c] = 0
			} else {
				done = true
				self.step = 5
				self.pathRow0 = row
				self.pathCol0 = col
			}
		}
	}
}

// setp5 支持函数
func (self *Munkres) findStarInCol(c int) int {
	row := -1
	for r := 0; r < self.nrow; r++ {
		if self.started[r][c] == 1 {
			row = r
		}
	}
	return row
}

func (self *Munkres) findPrimeInRow(r int) int {
	col := -1
	for j := 0; j < self.ncol; j++ {
		if self.started[r][j] == 2 {
			col = j
		}
	}
	return col
}

func (self *Munkres) augmentPath() {
	for p := 0; p < self.pathCount; p++ {
		if self.started[self.path[p][0]][self.path[p][1]] == 1 {
			self.started[self.path[p][0]][self.path[p][1]] = 0
		} else {
			self.started[self.path[p][0]][self.path[p][1]] = 1
		}
	}
}

func (self *Munkres) clearCovers() {
	for r := 0; r < self.nrow; r++ {
		self.rowCover[r] = 0
	}
	for c := 0; c < self.ncol; c++ {
		self.colCover[c] = 0
	}
}

func (self *Munkres) earasePrimes() {
	for r, row := range self.started {
		for c, col := range row {
			if col == 2 {
				self.started[r][c] = 0
			}
		}
	}
}

//Construct a series of alternating primed and starred zeros as follows.
//Let Z0 represent the uncovered primed zero found in Step 4.  Let Z1 denote
//the starred zero in the column of Z0 (if any). Let Z2 denote the primed zero
//in the row of Z1 (there will always be one).  Continue until the series
//terminates at a primed zero that has no starred zero in its column.
//Unstar each starred zero of the series, star each primed zero of the series,
//erase all primes and uncover every line in the matrix.  Return to Step 3.
func (self *Munkres) step5() {
	done := false
	r, c := -1, -1
	self.pathCount = 1
	self.path[self.pathCount-1][0] = self.pathRow0
	self.path[self.pathCount-1][1] = self.pathCol0
	for {
		if done {
			break
		}
		r = self.findStarInCol(self.path[self.pathCount-1][1])
		if r > -1 {
			self.pathCount++
			self.path[self.pathCount-1][0] = r
			self.path[self.pathCount-1][1] = self.path[self.pathCount-2][1]
		} else {
			done = true
		}
		if !done {
			c = self.findPrimeInRow(self.path[self.pathCount-1][0])
			self.pathCount++
			self.path[self.pathCount-1][0] = self.path[self.pathCount-2][0]
			self.path[self.pathCount-1][1] = c
		}
	}
	self.augmentPath()
	self.clearCovers()
	self.earasePrimes()
	self.step = 3
}

func (self *Munkres) findSmallest(minval int) int {
	rf := false
	for r, row := range self.maxtrix {
		rf = self.rowCover[r] == 0
		for c, col := range row {
			if rf && self.colCover[c] == 0 {
				if minval > col {
					minval = col
				}
			}
		}
	}
	return minval
}

//Add the value found in Step 4 to every element of each covered row, and subtract
//it from every element of each uncovered column.  Return to Step 4 without
//altering any stars, primes, or covered lines.
func (self *Munkres) step6() {
	rc := -1
	minval := math.MaxInt64
	minval = self.findSmallest(minval)
	for r, row := range self.maxtrix {
		rc = self.rowCover[r]
		for c, _ := range row {
			if rc == 1 {
				self.maxtrix[r][c] += minval
			}
			if self.colCover[c] == 0 {
				self.maxtrix[r][c] -= minval
			}
		}
	}
	self.step = 4
}

func (self *Munkres) step7() {

}

func (self *Munkres) RunMunkres() [][]int {
	done := false
	for {
		if done {
			break
		}
		//fmt.Println(self.step)
		//fmt.Println(self.started)
		switch self.step {
		case 1:
			self.step1()
		case 2:
			self.step2()
		case 3:
			self.step3()
		case 4:
			self.step4()
		case 5:
			self.step5()
		case 6:
			self.step6()
		case 7:
			self.step7()
			done = true
		}
	}
	return self.started
}
