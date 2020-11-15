package retrievor

import (
	"testing"
)

func TestClubBase(t *testing.T) {
	var c = ClubParse{}
	var countries = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1"}
	// Iterate trough all pages
	for _, e := range countries {
		c.CurrentPage = e
		c.ParseAll()

	}
	// c.ExportAsCSV()
}

func TestClubP(t *testing.T) {
	c := ClubParse{}
	country := "P"
	c.CurrentPage = country
	c.ParseAll()
	if c.Clubs[1].ID != "pacifico" {
		t.Errorf("Club at index 1 should be pacifico: is %s", c.Clubs[1].ID)
	}
	if c.Clubs[25].ID != "paris-saint-germain" {
		t.Errorf("Club at index 25 should be paris-saint-germain: is %s", c.Clubs[25].ID)
	}
}
