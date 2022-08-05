package table

import "testing"

func TestShow(t *testing.T) {
	LineEndTag = "*"
	LineBody = "="
	LineDivider = "/"
	Show([]string{"Module", "Type", "Path", "Author"}, [][]string{
		{"1111", "2", "3", "4"},
		{"1", "2", "3", "4"},
		{"1", "2", "3", "4"},
		{"1", "2", "3", "4"},
	})
}
