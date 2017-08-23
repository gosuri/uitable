package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/svenax/uitable"
)

type hacker struct {
	Name, Birthday, Bio string
}

var hackers = []hacker{
	{"Ada Lovelace", "December 10, 1815", "Ada was a British mathematician and writer, chiefly known for her work on Charles Babbage's early mechanical general-purpose computer, the Analytical Engine"},
	{"Alan Turing", "June 23, 1912", "Alan was a British pioneering computer scientist, mathematician, logician, cryptanalyst and theoretical biologist"},
}

var (
	headerFormat = color.New(color.FgBlack, color.Bold).SprintFunc()
	nameFormat   = color.New(color.FgGreen).SprintFunc()
)

func main() {
	table := uitable.New()
	table.CellFormatter = func(x, y int) uitable.Formatter {
		if y == 0 {
			return headerFormat
		}
		if x == 0 {
			return nameFormat
		}
		return uitable.DefaultFormatter
	}
	table.MaxColWidth = 50

	fmt.Println("==> List")
	table.AddRow("NAME", "BIRTHDAY", "BIO")
	for _, hacker := range hackers {
		table.AddRow(hacker.Name, hacker.Birthday, hacker.Bio)
	}
	fmt.Println(table)

	fmt.Print("\n==> Details\n")
	table = uitable.New()
	table.CellFormatter = func(x, y int) uitable.Formatter {
		if x == 0 {
			return headerFormat
		}
		if x == 1 && y%4 == 0 {
			return nameFormat
		}
		return uitable.DefaultFormatter
	}
	table.MaxColWidth = 80
	table.Wrap = true
	for _, hacker := range hackers {
		table.AddRow("Name:", hacker.Name)
		table.AddRow("Birthday:", hacker.Birthday)
		table.AddRow("Bio:", hacker.Bio)
		table.AddRow("") // blank
	}
	fmt.Println(table)
}
