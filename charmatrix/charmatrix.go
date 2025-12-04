package charmatrix

import "strings"

type Pos struct{ R, C int }

var (
	FourDirs  = []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	EightDirs = []Pos{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
)

type Matrix struct {
	data [][]rune
}

func FromStrings(rows []string) Matrix {
	out := make([][]rune, len(rows))
	for i, row := range rows {
		out[i] = []rune(row)
	}
	return Matrix{data: out}
}

func (m Matrix) Rows() int { return len(m.data) }
func (m Matrix) Cols() int {
	if len(m.data) == 0 {
		return 0
	}
	return len(m.data[0])
}

func (m Matrix) InBounds(p Pos) bool {
	return 0 <= p.R && p.R < len(m.data) &&
		0 <= p.C && p.C < len(m.data[p.R])
}

func (m Matrix) At(p Pos) rune {
	return m.data[p.R][p.C]
}

func (m Matrix) Set(p Pos, r rune) {
	m.data[p.R][p.C] = r
}

func (m Matrix) Clone() Matrix {
	dup := make([][]rune, len(m.data))
	for i := range m.data {
		dup[i] = append([]rune(nil), m.data[i]...)
	}
	return Matrix{data: dup}
}

func (m Matrix) Neighbours(p Pos, dirs []Pos) []Pos {
	res := make([]Pos, 0, len(dirs))
	for _, d := range dirs {
		q := Pos{p.R + d.R, p.C + d.C}
		if m.InBounds(q) {
			res = append(res, q)
		}
	}
	return res
}

func (m Matrix) String() string {
	b := strings.Builder{}
	for i, row := range m.data {
		b.WriteString(string(row))
		if i+1 < len(m.data) {
			b.WriteByte('\n')
		}
	}
	return b.String()
}
