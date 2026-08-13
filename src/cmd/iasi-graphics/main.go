package main

import (
	"flag"
	"fmt"
	"github.com/iasi-org/iasi-graphics/internal/compiler"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func run(args []string) error {
	if len(args) == 0 || args[0] != "render" {
		return fmt.Errorf("usage: iasi-graphics render <input.ig> [-o output.svg]")
	}
	fs := flag.NewFlagSet("render", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	out := fs.String("o", "", "output SVG file")
	var input string
	// Accept -o both before and after the input file.
	var flags []string
	for i := 1; i < len(args); i++ {
		if args[i] == "-o" {
			if i+1 >= len(args) {
				return fmt.Errorf("-o requires a path")
			}
			flags = append(flags, args[i], args[i+1])
			i++
		} else if strings.HasPrefix(args[i], "-") {
			flags = append(flags, args[i])
		} else if input == "" {
			input = args[i]
		} else {
			return fmt.Errorf("unexpected argument %q", args[i])
		}
	}
	if err := fs.Parse(flags); err != nil {
		return err
	}
	if input == "" {
		return fmt.Errorf("render requires an input .ig file")
	}
	if filepath.Ext(input) != ".ig" {
		return fmt.Errorf("input must have .ig extension: %s", input)
	}
	data, err := os.ReadFile(input)
	if err != nil {
		return fmt.Errorf("read %s: %w", input, err)
	}
	rendered, err := compiler.Compile(input, data)
	if err != nil {
		return err
	}
	target := *out
	if target == "" {
		target = strings.TrimSuffix(input, filepath.Ext(input)) + ".svg"
	}
	if err := os.WriteFile(target, rendered, 0644); err != nil {
		return fmt.Errorf("write %s: %w", target, err)
	}
	return nil
}
