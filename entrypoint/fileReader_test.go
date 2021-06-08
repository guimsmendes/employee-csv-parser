package entrypoint

import (
	"testing"
)

func TestParseTable(t *testing.T) {
	var lines []string
	lines = append(lines, "Name,Email,Wage,Number")
	lines = append(lines, "John Doe,doe@test.com,$10.00,1")

	if len(parseTable(lines)) != 2 {
		t.Errorf("Fail: The Parsing Table length for the rows expected was 2, got %d", len(parseTable(lines)))
	} else {
		t.Logf("Success: The Parsing Table length for the rows expected was 2, got %d", len(parseTable(lines)))
	}

	if len(parseTable(lines)[0]) != 4 {
		t.Errorf("Fail: The Parsing Table length for the columns expected was 4, got %d", len(parseTable(lines)[0]))
	} else {
		t.Logf("Success: The Parsing Table length for the columns expected was 4, got %d", len(parseTable(lines)[0]))
	}

}
