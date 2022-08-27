package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"
)

type hacker struct {
	Name, Birthday, Bio string
}

var hackers = []hacker{
	{"Ada Lovelace", "December 10, 1815", "Ada was a British mathematician and writer, chiefly known for her work on Charles Babbage's early mechanical general-purpose computer, the Analytical Engine"},
	{"Alan Turing", "June 23, 1912", "Alan was a British pioneering computer scientist, mathematician, logician, cryptanalyst and theoretical biologist"},
}

func main() {
	table := uitable.New()
	table.MaxColWidth = 50

	fmt.Println("==> List")
	table.AddRow("NAME", "BIRTHDAY", "BIO")
	for _, hacker := range hackers {
		table.AddRow(hacker.Name, hacker.Birthday, hacker.Bio)
	}
	fmt.Println(table)

	fmt.Print("\n==> Details\n")
	table = uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true
	for _, hacker := range hackers {
		table.AddRow("1. Name:", hacker.Name)
		table.AddRow("2. Birthday:", hacker.Birthday)
		table.AddRow("3. Bio:", hacker.Bio)
		table.AddRow("") // blank
	}
	fmt.Println(table)

	fmt.Print("\n==> Multicolor Support\n")
	table = uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true
	for _, hacker := range hackers {
		table.AddRow(color.RedString("1. Name:"), color.WhiteString(hacker.Name))
		table.AddRow(color.BlueString("2. Birthday:"), hacker.Birthday)
		table.AddRow(color.GreenString("3. Bio:"), hacker.Bio)
		table.AddRow("") // blank
	}
	fmt.Println(table)
}
