package table

import "testing"

func TestShow(t *testing.T) {
	//LineEndTag = "*"
	//LineBody = "="
	//LineDivider = "/"
	MaxCellWidth = 20
	Show([]string{"Module", "Type", "Path", "Author"}, [][]string{
		{"11111111111111111111111111111111111111", "2", "3", "4"},
		{"1", "2", "3", "4"},
		{"1", "2", "3", "4"},
		{"1", "2", "3", "4"},
	})
}
