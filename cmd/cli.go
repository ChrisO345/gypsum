package main

import (
	"fmt"
	"os"

	"github.com/chriso345/clifford"
)

// CLIArgs holds the parsed command-line arguments
type CLIArgs struct {
	clifford.Clifford `name:"gypsum"`
	clifford.Help
	clifford.Version `version:"0.0.1"`

	Paragraphs struct {
		Value             int
		clifford.Clifford `desc:"Number of paragraphs to generate (default: 3)"`
		clifford.ShortTag
	}

	Words struct {
		Value             int
		clifford.Clifford `desc:"Number of words per paragraph (default: 50)"`
		clifford.ShortTag
	}

	InFile struct {
		Value             string
		clifford.Clifford `desc:"Path to input text file (optional)"`
		clifford.ShortTag
	}

	OutFile struct {
		Value             string
		clifford.Clifford `desc:"Path to output file (optional)"`
		clifford.ShortTag
	}

	Seed struct {
		Value             int64
		clifford.Clifford `desc:"Random seed (default: current time)"`
		clifford.ShortTag
	}
}

// ParseArgs parses command-line flags using Clifford
func ParseArgs() *CLIArgs {
	args := &CLIArgs{}

	if err := clifford.Parse(args); err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing arguments:", err)
		os.Exit(1)
	}

	if args.Paragraphs.Value == 0 {
		args.Paragraphs.Value = 3
	}

	if args.Words.Value == 0 {
		args.Words.Value = 50
	}

	return args
}
