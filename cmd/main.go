package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/chriso345/gypsum"
)

func main() {
	args := ParseArgs()
	_ = args

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	chain := gypsum.New(rng)

	if args.InFile.Value != "" {
		// Load input file
		data, err := os.ReadFile(args.InFile.Value)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input file:", err)
			os.Exit(1)
		}
		chain.Learn(string(data))
	} else {
		// Use default lorem ipsum
		chain.Learn(gypsum.DefaultLorem)
	}

	if args.OutFile.Value != "" {
		// Redirect output to file
		f, err := os.Create(args.OutFile.Value)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file:", err)
			os.Exit(1)
		}
		defer f.Close()
		os.Stdout = f
	}

	for p := 0; p < args.Paragraphs.Value; p++ {
		out := chain.Generate(args.Words.Value)
		fmt.Println(out)
		fmt.Println()
	}
}
