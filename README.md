# gypsum

`gypsum` is a lightweight Lorem Ipsum generator written in Go. It uses a Markov chain to generate pseudo-Latin placeholder text, making it easy to create paragraphs and titles for testing, prototyping, or design purposes.

---

## Features

* Generate random Lorem Ipsum text
* Control the number of paragraphs and words per paragraph
* Deterministic or random output using a seedable RNG
* Lightweight and dependency-free Go implementation
* Simple command-line interface
* Supports title-case generation for headings

---

## Installation

Build `gypsum` from source:

```bash
git clone https://github.com/chriso345/gypsum.git
cd gypsum
go build ./cmd/gypsum

# Copy the binary to a directory in your PATH
cp gypsum $HOME/.local/bin/gypsum
```

Or install directly using `go install`:

```bash
go install github.com/chriso345/gypsum/cmd/gypsum@latest
```

---

## Usage

Generate Lorem Ipsum text using the CLI:

```bash
# Generate 3 paragraphs with 50 words each (default)
gypsum

# Generate 5 paragraphs with 30 words each
gypsum -p 5 -w 30

# Display help information
gypsum --help
```

Alternatively, use `gypsum` as a Go library:

```go
package main

import (
  "fmt"
  "math/rand"
  "time"
  "github.com/chriso345/gypsum"
)

func main() {
  // Create a new Markov chain with a random seed
  rng := rand.New(rand.NewSource(42))

  // Initialize the Markov chain generator
  mc := gypsum.New(rng)

  // Train the generator with sample text
  mc.Learn("one two three four five six seven eight nine ten")

  // Generate and print 5 words of Random text
  fmt.Println(mc.Generate(5))

  // Generate and print 3 paragraphs of Lorem Ipsum text with 10 words each
  fmt.Println(gypsum.LoremIpsum(3, 10, rng))
}

```

### Example Output (3 paragraphs, 10 words each)

```
Dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor.

Voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur.

Ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis.
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
