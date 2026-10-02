package uitable

import (
	"strings"
	"sync"
	"testing"

	"github.com/fatih/color"
	"github.com/gosuri/uitable/util/strutil"
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
			{Data: "foo", Width: 3, Wrap: true},
			{Data: "bar baz", Width: 3, Wrap: true},
		},
	}
	got := row.String()
	need := "foo\tbar\n   \tbaz"

	if got != need {
		t.Fatalf("need: %q | got: %q ", need, got)
	}
}

func TestAlign(t *testing.T) {
	table := New()
	table.AddRow("foo", "bar baz")
	table.Rows = []*Row{{
		Separator: "\t",
		Cells: []*Cell{
			{Data: "foo", Width: 5, Wrap: true},
			{Data: "bar baz", Width: 10, Wrap: true},
		},
	}}
	table.RightAlign(1)
	got := table.String()
	need := "foo  \t   bar baz"

	if got != need {
		t.Fatalf("need: %q | got: %q ", need, got)
	}
}

func TestAddRow(t *testing.T) {
	var wg sync.WaitGroup
	table := New()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			table.AddRow("foo")
		}()
	}
	wg.Wait()
	if len(table.Rows) != 100 {
		t.Fatal("want", 100, "got", len(table.Rows))
	}
}

func TestCellColorWrap(t *testing.T) {
	c := &Cell{
		Data:  color.RedString("foo bar"),
		Width: 10,
		Wrap:  true,
	}

	got := c.String()
	need := color.RedString("foo bar") + "   "
	if got != need {
		t.Fatalf("need %q, got %q", need, got)
	}
	if w := strutil.StringWidth(got); w != 10 {
		t.Fatalf("need visible width 10, got %d", w)
	}

	table := New()
	table.Wrap = true
	table.AddRow(color.RedString("1. Name:"), "Ada Lovelace")
	table.AddRow(color.BlueString("2. Birthday:"), "December 10, 1815")
	tgot := table.String()
	tneed := color.RedString("1. Name:") + "    \tAda Lovelace     \n" + color.BlueString("2. Birthday:") + "\tDecember 10, 1815"
	if tgot != tneed {
		t.Fatalf("need:\n%q\ngot:\n%q", tneed, tgot)
	}

	c2 := &Cell{
		Data:  color.RedString("foo bar"),
		Width: 5,
		Wrap:  true,
	}
	got2 := c2.String()
	if c2.LineWidth() != 3 {
		t.Fatalf("need line width 3, got %d", c2.LineWidth())
	}
	if lines := strings.Split(got2, "\n"); len(lines) != 2 {
		t.Fatalf("need 2 lines, got %d", len(lines))
	}
}
