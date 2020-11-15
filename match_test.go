package retrievor

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestParseLittleRange_1(t *testing.T) {
	r := MatchesResult{}
	r.ParseAllWithStringRange("2009-01-31", "2009-02-01")
	content, _ := ioutil.ReadFile("asserts/matchs/assert_2009-01-31_2009-02-01.txt")
	expected := string(content)
	resultJSON, _ := json.Marshal(r.Matches)
	result := string(resultJSON)
	if expected != result {
		t.Errorf("Matches struct are differents than expected. Expected=%s, Result=%s", expected, result)
	}
}

func TestParseLittleRange_2(t *testing.T) {
	r := MatchesResult{}
	r.ParseAllWithStringRange("2012-01-31", "2012-02-01")
	content, _ := ioutil.ReadFile("asserts/matchs/assert_2012-01-31_2012-02-01.txt")
	expected := string(content)
	resultJSON, _ := json.Marshal(r.Matches)
	result := string(resultJSON)
	if expected != result {
		t.Errorf("Matches struct are differents than expected. Expected=%s, Result=%s", expected, result)
	}
}

func TestParseLittleRange_3(t *testing.T) {
	r := MatchesResult{}
	r.ParseAllWithStringRange("2020-07-15", "2020-07-20")
	content, _ := ioutil.ReadFile("asserts/matchs/assert_2020-07-15_2020-07-20.txt")
	expected := string(content)
	resultJSON, _ := json.Marshal(r.Matches)
	result := string(resultJSON)
	if expected != result {
		t.Errorf("Matches struct are differents than expected. Expected=%s, Result=%s", expected, result)
	}
}
