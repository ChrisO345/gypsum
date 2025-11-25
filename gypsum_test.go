package gypsum

import (
	"math/rand"
	"testing"
)

func TestGenerate(t *testing.T) {
	mc := New(nil)
	mc.Learn("one two three four")
	out := mc.Generate(5)
	t.Logf("Generated output: %q", out)
	if out == "" {
		t.Fatal("expected non-empty output")
	}
}

func TestGenerateNonNilRand(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	mc := New(rng)
	mc.Learn("one two three four five six seven eight nine ten")
	out := mc.Generate(5)
	t.Logf("Generated output with non-nil RNG: %q", out)
	if out == "" {
		t.Fatal("expected non-empty output")
	}
}

func TestLearnEmpty(t *testing.T) {
	mc := New(nil)
	mc.Learn("")
	if len(mc.Map) != 0 {
		t.Fatalf("expected empty map, got %d entries", len(mc.Map))
	}
}

func TestLearnShortInput(t *testing.T) {
	mc := New(nil)
	mc.Learn("one two")
	if len(mc.Map) != 0 {
		t.Fatalf("expected empty map for short input, got %d entries", len(mc.Map))
	}
}

func TestRandomKey(t *testing.T) {
	mc := New(nil)
	mc.Learn("one two three four five six")
	key := mc.randomKey()
	t.Logf("Random key: %+v", key)
	if key.A == "" || key.B == "" {
		t.Fatal("expected non-empty bigram")
	}
}

func TestGenerateNoLearn(t *testing.T) {
	mc := New(nil)
	out := mc.Generate(5)
	if out != "" {
		t.Fatalf("expected empty output when no learning done, got %q", out)
	}
}

func TestGenerateZeroWords(t *testing.T) {
	mc := New(nil)
	mc.Learn("one two three four five")
	out := mc.Generate(0)
	if out != "" {
		t.Fatalf("expected empty output for zero words, got %q", out)
	}
}

func TestGenerateSingleWord(t *testing.T) {
	mc := New(nil)
	mc.Learn("one two three four five")
	out := mc.Generate(1)
	t.Logf("Generated single word output: %q", out)
	if out == "" {
		t.Fatal("expected non-empty output for single word")
	}
}

func TestGenerateMoreWordsThanLearned(t *testing.T) {
	mc := New(nil)
	mc.Learn("one two three")
	out := mc.Generate(10)
	t.Logf("Generated output with more words than learned: %q", out)
	if out == "" {
		t.Fatal("expected non-empty output")
	}
}

func TestGenerateExactWords(t *testing.T) {
	mc := New(nil)
	mc.Learn("one two three four five six seven eight nine ten")
	out := mc.Generate(10)
	t.Logf("Generated output with exact words learned: %q", out)
	if out == "" {
		t.Fatal("expected non-empty output")
	}
}

func TestGenerateRepeatedCalls(t *testing.T) {
	mc := New(nil)
	mc.Learn("one two three four five six")
	for i := range 5 {
		out := mc.Generate(5)
		t.Logf("Generated output call %d: %q", i+1, out)
		if out == "" {
			t.Fatalf("expected non-empty output on call %d", i+1)
		}
	}
}

func TestLearnMultipleTimes(t *testing.T) {
	mc := New(nil)
	mc.Learn("one two three")
	mc.Learn("four five six")
	out := mc.Generate(6)
	t.Logf("Generated output after multiple learns: %q", out)
	if out == "" {
		t.Fatal("expected non-empty output")
	}
}

func TestGenerateLoremIpsum(t *testing.T) {
	mc := New(nil)
	mc.Learn(DefaultLorem)
	out := mc.Generate(50)
	t.Logf("Generated Lorem Ipsum output: %q", out)
	if out == "" {
		t.Fatal("expected non-empty output")
	}
}

func TestLoremIpsumLearning(t *testing.T) {
	out := LoremIpsum(3, 50, nil)
	t.Logf("Generated Lorem Ipsum paragraphs: %q", out)
	if out == "" {
		t.Fatal("expected non-empty output")
	}
}
