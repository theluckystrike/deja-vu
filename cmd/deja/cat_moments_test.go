package main

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"
)

// #4624: the cat has a moment beside an empty search and beside a first recall
// that finds something, in a terminal only. The sprite renders for both, and
// each carries a line that says which moment it is. The terminal-only gating is
// tested separately in logo_devnull_test.go (logoWanted) and at the runSearch
// call sites; these tests pin the rendering that the gated call sites show.
func TestCatMomentShowsOnEmptySearch(t *testing.T) {
	var b bytes.Buffer
	mood, line := searchCatMoment(true) // empty search
	printLogoMood(&b, []string{line}, mood)
	out := b.String()
	// The sprite (half blocks) and the empty-search line are both present.
	if !strings.Contains(out, "█") && !strings.Contains(out, "▀") && !strings.Contains(out, "▄") {
		t.Fatalf("empty-search cat rendered no sprite: %q", out)
	}
	if !strings.Contains(out, "nothing") {
		t.Fatalf("empty-search cat should say nothing matched: %q", out)
	}
}

func TestCatMomentShowsOnFirstRecall(t *testing.T) {
	var b bytes.Buffer
	mood, line := searchCatMoment(false) // a recall that found something
	printLogoMood(&b, []string{line}, mood)
	out := b.String()
	if !strings.Contains(out, "█") && !strings.Contains(out, "▀") && !strings.Contains(out, "▄") {
		t.Fatalf("first-recall cat rendered no sprite: %q", out)
	}
	if !strings.Contains(out, "found") {
		t.Fatalf("first-recall cat should say something was found: %q", out)
	}
}

// The two moments carry different moods, so the empty search does not look like
// a win. (moodReady vs moodSurprised draw the sprite differently.)
func TestCatMomentsUseDistinctMoods(t *testing.T) {
	emptyMood, _ := searchCatMoment(true)
	recallMood, _ := searchCatMoment(false)
	if emptyMood == recallMood {
		t.Fatalf("empty search and first recall should use different moods, both %v", emptyMood)
	}
}

// The ready cat greets only the first recall that finds something on an index,
// not every result list after it.
func TestFirstRecallIsOncePerIndex(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "index")
	if !firstRecall(dir) {
		t.Fatal("the first recall on a fresh index should get the cat")
	}
	if firstRecall(dir) {
		t.Fatal("a second recall on the same index should not get the cat again")
	}
	if !firstRecall(filepath.Join(t.TempDir(), "other")) {
		t.Fatal("another index has its own first recall")
	}
}
