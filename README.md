# uitable [![GoDoc](https://godoc.org/github.com/gosuri/uitable?status.svg)](https://godoc.org/github.com/gosuri/uitable) [![Build Status](https://travis-ci.org/gosuri/uitable.svg?branch=master)](https://travis-ci.org/gosuri/uitable)

`uitable` is a go library for representing data as tables for terminal applications. It provides primitives for sizing and wrapping columns to improve readability. It also makes it possible to attach a formatter to each table cell, typically used to colorize terminal output.

## Example Usage

Full source code for the example is available at [example/main.go](example/main.go)

```go
table := uitable.New()
table.CellFormatter = func(x, y int) uitable.Formatter {
  if y == 0 { return headerFormat }
  if x == 0 { return nameFormat }
  return uitable.DefaultFormatter
}
table.MaxColWidth = 50

table.AddRow("NAME", "BIRTHDAY", "BIO")
for _, hacker := range hackers {
  table.AddRow(hacker.Name, hacker.Birthday, hacker.Bio)
}
fmt.Println(table)
```

Will render the data like:

<table>
  <tr>
    <td><b>NAME</b></td>
    <td><b>BIRTHDAY</b></td>
    <td><b>BIO</b></td>
  </tr>
  <tr>
    <td style="color: green">Ada Lovelace</td>
    <td>December 10, 1815</td>
    <td>Ada was a British mathematician and writer, chi...</td>
  </tr>
  <tr>
    <td  style="color: green">Alan Turing</td>
    <td>June 23, 1912</td>
    <td>Alan was a British pioneering computer scientis...</td>
  </tr>
</tr>

</table>

For wrapping in two columns:

```go
table = uitable.New()
table.CellFormatter = func(x, y int) uitable.Formatter {
  if x == 0 { return headerFormat }
  if x == 1 && y%4 == 0 { return nameFormat }
  return uitable.DefaultFormatter
}
table.MaxColWidth = 80
table.Wrap = true // wrap columns

for _, hacker := range hackers {
  table.AddRow("Name:", hacker.Name)
  table.AddRow("Birthday:", hacker.Birthday)
  table.AddRow("Bio:", hacker.Bio)
  table.AddRow("") // blank
}
fmt.Println(table)
```

Will render the data like:

<table align="top">
  <tr>
    <td><b>Name:</b><td>
    <td style="color: green">Ada Lovelace</td>
  </tr>
  <tr>
    <td><b>Birthday:</b><td>
    <td>December 10, 1815</td>
  </tr>
  <tr>
    <td><b>Bio:</b><td>
    <td>Ada was a British mathematician and writer, chiefly known for her work on Charles Babbage's early mechanical general-purpose computer, the Analytical Engine</td>
  </tr>
  <tr>
    <td><b>Name:</b><td>
    <td style="color: green">Alan Turing</td>
  </tr>
  <tr>
    <td><b>Birthday:</b><td>
    <td>June 23, 1912</td>
  </tr>
  <tr>
    <td><b>Bio:</b><td>
    <td>Alan was a British pioneering computer scientist, mathematician, logician, cryptanalyst and theoretical biologist</td>
  </tr>
</table>

## Installation

```
$ go get -v github.com/gosuri/uitable
```


[![Bitdeli Badge](https://d2weczhvl823v0.cloudfront.net/gosuri/uitable/trend.png)](https://bitdeli.com/free "Bitdeli Badge")

