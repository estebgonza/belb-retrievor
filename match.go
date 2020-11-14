package retrievor

// Retrieve all match from www.matchendirect.fr
// Ouput in CSV
//
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Start date 2009-12-01
const matchURLPattern = "https://footballdatabase.com/scores/%s"

//const scope = "monde"
// const date = "2019-03"

// ResultParse Matches by date
type MatchesResult struct {
	currentDate time.Time
	Matches     []Match
}

// Match Data about a specific match
type Match struct {
	Date            time.Time `json:"time"`
	CompetitionId   string    `json:"competition_id"`
	CompetitionName string    `json:"competition_name"`
	T1Id            string    `json:"t1_id"`
	T2Id            string    `json:"t2_id"`
	T1Name          string    `json:"t1_name"`
	T2Name          string    `json:"t2_name"`
	T1Score         int       `json:"t1_score"`
	T2Score         int       `json:"t2_score"`
	WinnerId        string    `json:"winner_id"`
	Draw            bool      `json:"draw"`
}

// Returns formatted url for the date
func getMatchURL(date string) string {
	return fmt.Sprintf(matchURLPattern, date)
}

func (m *Match) checkWinner() {
	m.checkDraw()
	if !m.Draw {
		if m.T1Score > m.T2Score {
			m.WinnerId = m.T1Id
		} else {
			m.WinnerId = m.T2Id
		}
	}
}

func (m *Match) checkDraw() {
	m.Draw = m.T1Score == m.T2Score
}

// ExportAsCSV export results in a CSV file
func (r *MatchesResult) ExportAsCSV() error {
	f, err := os.Create(fmt.Sprintf("matches-%d.csv", time.Now().Second()))
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, m := range r.Matches {
		line := fmt.Sprintf("\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%d\",\"%d\",\"%s\",\"%t\"\n", m.Date.String(), m.CompetitionId, m.CompetitionName, m.T1Name, m.T2Name, m.T1Id, m.T2Id, m.T1Score, m.T2Score, m.WinnerId, m.Draw)
		w.WriteString(line)
	}
	w.Flush()
	return nil
}

func (r *MatchesResult) ParseAllWithStringRange(start string, end string) error {
	s, err := time.Parse("2006-01-02", start)
	if err != nil {
		return err
	}
	e, err := time.Parse("2006-01-02", end)
	if err != nil {
		return err
	}
	r.ParseAllWithTimeRange(s, e)
	return nil
}

func (r *MatchesResult) ParseAllWithTimeRange(start time.Time, end time.Time) error {
	for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		r.currentDate = d
		date := d.Format("2006-01-02")
		err := r.ParsePage(date)
		if err != nil {
			return err
		}
	}
	return nil
}

// ParseAll is the starting point of the module
func (r *MatchesResult) ParsePage(date string) error {
	var url = getMatchURL(date)
	fmt.Print(fmt.Sprintf("> %s : ", url))
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	r.parseMatches(doc)
	fmt.Println("OK.")

	return nil
}

func (r *MatchesResult) parseMatches(doc *goquery.Document) {
	var competitionId string
	var competitionName string
	doc.Find(".col-md-9").Children().Each(func(i int, s *goquery.Selection) {
		node := s.Nodes[0]
		divType := node.Data
		// Date
		switch divType {
		case "a":
			v, e := s.Attr("name")
			if e {
				competitionId = v
			}
			break
		case "h4":
			competitionName = s.Text()
			break
		case "div":
			// Create Match struct
			m := Match{}
			m.Date = r.currentDate
			m.CompetitionId = competitionId
			m.CompetitionName = competitionName
			// Check if it's a result match div
			v, e := s.Attr("class")
			if e && strings.Contains(v, "club-gamelist-match") {
				var t1Name, t2Name, t1Id, t2Id string
				scoreString := s.Find(".club-gamelist-match-score").Text()
				scores := strings.Split(strings.TrimSpace(scoreString), "-")
				s.Find(".club-gamelist-match-clubs").Each(func(i int, s *goquery.Selection) {
					if t1Id == "" {
						e, _ := s.Find("a").Attr("href")
						t1Id = strings.Split(e, "/")[2]
						t1Name = s.Text()
					} else {
						e, _ := s.Find("a").Attr("href")
						t2Id = strings.Split(e, "/")[2]
						t2Name = s.Text()
					}
				})
				s1, _ := strconv.Atoi(strings.TrimSpace(scores[0]))
				s2, _ := strconv.Atoi(strings.TrimSpace(scores[1]))
				// Fill results
				m.T1Id = t1Id
				m.T2Id = t2Id
				m.T1Name = t1Name
				m.T2Name = t2Name
				m.T1Score = s1
				m.T2Score = s2
				m.checkWinner()
				m.checkDraw()
				r.Matches = append(r.Matches, m)
			}
		}
	})
}
