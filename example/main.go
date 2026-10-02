package main

import (
	"fmt"

	"github.com/gosuri/uitable"
)

const (
	ansiReset = "\033[0m"
	ansiRed   = "\033[31m"
	ansiGreen = "\033[32m"
	ansiBlue  = "\033[34m"
	ansiWhite = "\033[37m"
)

func red(s string) string   { return ansiRed + s + ansiReset }
func green(s string) string { return ansiGreen + s + ansiReset }
func blue(s string) string  { return ansiBlue + s + ansiReset }
func white(s string) string { return ansiWhite + s + ansiReset }

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
		table.AddRow("Name:", hacker.Name)
		table.AddRow("Birthday:", hacker.Birthday)
		table.AddRow("Bio:", hacker.Bio)
		table.AddRow("") // blank
	}
	fmt.Println(table)

	fmt.Print("\n==> Multicolor Support\n")
	table = uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true
	for _, hacker := range hackers {
		table.AddRow(red("Name:"), white(hacker.Name))
		table.AddRow(blue("Birthday:"), hacker.Birthday)
		table.AddRow(green("Bio:"), hacker.Bio)
		table.AddRow("") // blank
	}
	fmt.Println(table)
}
