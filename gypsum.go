package gypsum

import (
	"math/rand"
	"strings"
)

import _ "embed"

//go:embed lorem_ipsum.txt
var DefaultLorem string

// Bigram: two consecutive words.
type Bigram struct {
	A, B string
}

// MarkovChain is a simple order-2 markov chain
type MarkovChain struct {
	Map  map[Bigram][]string
	Keys []Bigram
	RNG  *rand.Rand
}

// New creates a new empty chain.
// If rng is nil, a deterministic RNG is used.
func New(rng *rand.Rand) *MarkovChain {
	if rng == nil {
		rng = rand.New(rand.NewSource(97))
	}
	return &MarkovChain{
		Map: make(map[Bigram][]string),
		RNG: rng,
	}
}

// Learn takes input text and feeds its word transitions to the chain.
func (m *MarkovChain) Learn(text string) {
	words := strings.Fields(text)
	if len(words) < 3 {
		return
	}

	for i := 0; i < len(words)-2; i++ {
		b := Bigram{words[i], words[i+1]}
		m.Map[b] = append(m.Map[b], words[i+2])
	}

	// Refresh keys
	m.Keys = m.Keys[:0]
	for k := range m.Map {
		m.Keys = append(m.Keys, k)
	}
}

// randomKey picks a random existing bigram
func (m *MarkovChain) randomKey() Bigram {
	return m.Keys[m.RNG.Intn(len(m.Keys))]
}

// Generate produces n words
func (m *MarkovChain) Generate(n int) string {
	if n <= 0 || len(m.Map) == 0 {
		return ""
	}

	// Pick starting bigram
	state := m.randomKey()
	result := []string{state.A, state.B}

	for len(result) < n {
		nextWords := m.Map[state]
		if len(nextWords) == 0 {
			// dead end, jump to a different random state
			state = m.randomKey()
			continue
		}
		next := nextWords[m.RNG.Intn(len(nextWords))]
		result = append(result, next)

		// shift
		state = Bigram{state.B, next}
	}

	return finalizeSentence(result)
}

// finalizeSentence capitalizes first word and ensures punctuation
func finalizeSentence(words []string) string {
	if len(words) == 0 {
		return ""
	}

	words[0] = strings.ToUpper(string(words[0][0])) + words[0][1:]
	text := strings.Join(words, " ")

	// If already ends with strong punctuation, keep
	if strings.HasSuffix(text, ".") ||
		strings.HasSuffix(text, "!") ||
		strings.HasSuffix(text, "?") {
		return text
	}

	// Otherwise trim trailing punctuation and add `.`
	text = strings.TrimRight(text, ",;:-_~")
	return text + "."
}

func LoremIpsum(paragraphs, wordsPerParagraph int, rng *rand.Rand) string {
	chain := New(rng)
	chain.Learn(DefaultLorem)

	var result strings.Builder
	for p := range paragraphs {
		out := chain.Generate(wordsPerParagraph)
		result.WriteString(out)
		if p < paragraphs-1 {
			result.WriteString("\n\n")
		}
	}
	return result.String()
}
