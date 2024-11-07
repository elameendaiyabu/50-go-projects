package main

import (
	"bytes"
	"testing"
)

func TestSearch(t *testing.T) {
	tasks := Tasks{1: "sleep"}

	t.Run("Known word", func(t *testing.T) {
		got := tasks.Search(1)
		want := true

		assertBools(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		got := tasks.Search(2)
		want := false

		assertBools(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	tasks := Tasks{}
	tasks.Add("sleep")

	// search with index
	got := tasks.Search(1)
	want := true

	assertBools(t, got, want)
}

func TestUpdate(t *testing.T) {
	tasks := Tasks{}
	tasks.Add("sleep")
	tasks.Add("watch")

	got := tasks.Update(2, "eat")
	if got != nil {
		t.Error(got)
	}
}

func TestDelete(t *testing.T) {
	tasks := Tasks{}
	tasks.Add("sleep")
	tasks.Add("watch")

	got := tasks.Delete(2)
	if got != nil {
		t.Error(got)
	}
}

func TestGetTasks(t *testing.T) {
	buffer := bytes.Buffer{}
	tasks := Tasks{}
	tasks.Add("sleep")
	tasks.Add("watch")
	tasks.GetTasks(&buffer)

	got := buffer.String()
	want := `1. sleep
2. watch
`

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}

func assertBools(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
