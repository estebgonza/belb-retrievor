package retrievor

import (
	"testing"
)

func TestMatchBase(t *testing.T) {
	scope := "monde"
	date := "2020-01"
	var r = ResultParse{CurrentDate: date, Scope: scope}
	r.ParseAll()
	// Iterate trough all pages
	for ; r.NextDate != ""; r.ParseAll() {
		if r.NextDate <= r.CurrentDate {
			break
		}
		// Switch page
		r.CurrentDate = r.NextDate
		r.NextDate = ""
	}
	r.ExportAsCSV()
}

func TestMatchResult(t *testing.T) {
	scope := "monde"
	date := "2020-01"
	var r = ResultParse{CurrentDate: date, Scope: scope}
	r.ParseAll()
	if r.Matches[0].T2 != "Auxerre" {
		t.Errorf("Team 2 of match 0 should be Auxerre: is %s", r.Matches[0].T2)
	}
	if r.Matches[2].Draw {
		t.Errorf("Draw of match 2 should be false: is %t", r.Matches[2].Draw)
	}
	if !r.Matches[5].Draw {
		t.Errorf("Draw of match 5 should be true: is %t", r.Matches[2].Draw)
	}
	if r.Matches[9].Winner != "NEC" {
		t.Errorf("Winner of match 9 should be NEC: is %s", r.Matches[9].Winner)
	}
	// r.ExportAsCSV()
}
