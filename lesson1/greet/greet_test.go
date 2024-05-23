package greet

import "testing"

func TestGreet(t *testing.T) {
	expected := "Hello, Sultan! Your table is 24.\n"
	text := Greet("Sultan", 24)
	if text != expected {
		t.Errorf("Greet(%s, %d), got: %s, expected: %s", "Sultan", 24, text, expected)
	}
}

func TestTableGreet(t *testing.T) {
	table := []struct {
		name     string
		table    int
		expected string
	}{
		{"Sultan", 24, "Hello, Sultan! Your table is 24."},
		{"Zhandos", 47, "Hello, Zhandos! Your table is 47."},
		{"Rosen", 98, "Hello, Rosen! Your table is 98."},
	}

	for _, row := range table {
		text := Greet(row.name, row.table)
		if text != row.expected {
			t.Errorf("Greet(%s, %d), got: %s, expected: %s", row.name, row.table)
		}
	}
}
