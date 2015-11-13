package uitable

import (
	"testing"
)

func TestCell(t *testing.T) {
	c := &Cell{
		Data:  "foo bar",
		Width: 5,
	}

	got := c.String()
	if got != "fo..." {
		t.Fatal("need", "fo...", "got", got)
	}
	if c.LineWidth() != 5 {
		t.Fatal("need", 5, "got", c.LineWidth())
	}

	c.Wrap = true
	got = c.String()
	if got != "foo\nbar" {
		t.Fatal("need", "foo\nbar", "got", got)
	}
	if c.LineWidth() != 3 {
		t.Fatal("need", 3, "got", c.LineWidth())
	}
}

func TestRow(t *testing.T) {
	row := &Row{
		Separator: "\t",
		Cells: []*Cell{
			&Cell{Data: "foo", Width: 3, Wrap: true},
			&Cell{Data: "bar baz", Width: 3, Wrap: true},
		},
	}
	got := row.String()
	need := "foo\tbar\n   \tbaz"

	if got != need {
		t.Fatalf("need: %q | got: %q ", need, got)
	}
}
