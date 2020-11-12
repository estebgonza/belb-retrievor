package retrievor

// Retrieve all match from www.matchendirect.fr
// Ouput in CSV
//
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const urlPattern = "https://www.matchendirect.fr/%s/%s"
const scope = "monde"
const date = "2019-03"

// ResultParse Matches by date
type ResultParse struct {
	currentDate string  // current date
	nextDate    string  // next date
	Matches     []Match // list of matches for current date
}

// Match Data about a specific match
type Match struct {
	Date        time.Time `json:"time"`
	Competition string    `json:"competition"`
	T1          string    `json:"t1"`
	T2          string    `json:"t2"`
	Score       string    `json:"score"`
	Status      string    `json:"status"`
}

func main() {
	fmt.Println("El Retrievor")
	// Initialize on first page
	var r = ResultParse{currentDate: date}
	r.parseAll()
	// Iterate trough all pages
	for ; r.nextDate != ""; r.parseAll() {
		if r.nextDate <= r.currentDate {
			break
		}
		// Switch page
		r.currentDate = r.nextDate
		r.nextDate = ""
	}
	r.exportAsCSV()
}

// Returns formatted url with scope and date
func getURL(scope string, date string) string {
	return fmt.Sprintf(urlPattern, scope, date)
}

func (r *ResultParse) exportAsCSV() error {
	f, err := os.Create(fmt.Sprintf("matches-%d.csv", time.Now().Second()))
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, m := range r.Matches {
		line := fmt.Sprintf("\"%s\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\n", m.Date.String(), m.Competition, m.T1, m.T2, m.Score, m.Status)
		w.WriteString(line)
	}
	w.Flush()
	return nil
}

func (r *ResultParse) parseAll() error {
	var url = getURL(scope, r.currentDate)
	fmt.Print(fmt.Sprintf("> %s :", url))
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	r.parseNextDate(doc)
	r.parseMatches(doc)
	fmt.Println(" OK!")

	return nil
}

func (r *ResultParse) parseNextDate(doc *goquery.Document) {
	doc.Find(".objselect_prevnext").Each(func(i int, s *goquery.Selection) {
		valueHref, exist := s.Attr("href")
		if exist {
			r.nextDate = strings.Split(valueHref, "/")[2]
		}
	})
}

func (r *ResultParse) parseMatches(doc *goquery.Document) {
	doc.Find("div.panel-info").Each(func(i int, s *goquery.Selection) {
		// Find competition name
		currentCompetition, _ := s.Find("div.lienCompetition a").Html()
		if currentCompetition == "" {
			// Not in panel for competition
			return
		}
		var currentDate string
		s.Find("div.panel-body > table.table").Children().Each(func(i int, s *goquery.Selection) {
			node := s.Nodes[0]
			divType := node.Data
			// Date

			if divType == "thead" {
				currentDate, _ = s.Find("tr > th").Html()
			} else if divType == "tbody" {
				s.Find("tr").Each(func(i int, s *goquery.Selection) {
					var t1 = s.Find(".lm3_eq1").Text()
					var t2 = s.Find(".lm3_eq2").Text()
					var hours = s.Find(".lm1").Text()
					var score = s.Find(".lm3_score").Text()
					var status = s.Find(".lm2_0").Text()
					score = strings.TrimSpace(score)
					status = parseStatus(status)
					date := convToDate(currentDate, hours)
					m := Match{
						Competition: currentCompetition,
						Date:        date,
						T1:          t1,
						T2:          t2,
						Score:       score,
						Status:      status,
					}
					r.Matches = append(r.Matches, m)
				})
			}
		})
	})
}

func convToDate(frDate string, hours string) time.Time {
	elem := strings.Split(frDate, " ")
	day := elem[1]
	month := monthToNumber(elem[2])
	year := elem[3]
	formattedDateString := fmt.Sprintf("%s-%s-%s %s", year, month, day, hours)
	date, _ := time.Parse("2006-01-02 15:04", formattedDateString)
	return date
}

func monthToNumber(frMonth string) string {
	frMonth = strings.ToUpper(frMonth)
	switch frMonth {
	case "JANVIER":
		return "01"
	case "FÉVRIER":
		return "02"
	case "MARS":
		return "03"
	case "AVRIL":
		return "04"
	case "MAI":
		return "05"
	case "JUIN":
		return "06"
	case "JUILLET":
		return "07"
	case "AOÛT":
		return "08"
	case "SEPTEMBRE":
		return "09"
	case "OCTOBRE":
		return "10"
	case "NOVEMBRE":
		return "11"
	case "DÉCEMBRE":
		return "12"
	}
	log.Fatalf("Unrecognize month: %s", frMonth)
	return "01"
}

func parseStatus(statusText string) string {
	statusText = strings.ReplaceAll(statusText, "-- : --", "")
	if strings.Contains(statusText, ":") {
		statusText = statusText[5:len(statusText)]
	}
	return statusText
}
