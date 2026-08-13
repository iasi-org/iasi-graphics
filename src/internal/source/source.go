package source

import "strings"

type File struct {
	Name string
	Text string
}

func New(name string, data []byte) File {
	text := strings.ReplaceAll(string(data), "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "\n")
	return File{Name: name, Text: text}
}

type Position struct{ Line, Column int }
