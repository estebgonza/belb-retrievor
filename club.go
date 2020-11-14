package retrievor

// Retrieve all clubs on https://footballdatabase.com
// Ouput in CSV
//
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const clubURLPattern = "https://footballdatabase.com/clubs-list-letter/%s"

var countries = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1"}

// ClubParse Main struct for parsing soccers clubs in footballdatabase
type ClubParse struct {
	CurrentPage string
	Clubs       []Club
}

// Club Represent one club
type Club struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getURL(page string) string {
	return fmt.Sprintf(clubURLPattern, page)
}

// ExportAsCSV Used to export content of ClubParse in a csv file
func (r *ClubParse) ExportAsCSV() error {
	f, err := os.Create(fmt.Sprintf("clubs-%d.csv", time.Now().Second()))
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, e := range r.Clubs {
		line := fmt.Sprintf("\"%s\",\"%s\"\n", e.ID, e.Name)
		w.WriteString(line)
	}
	w.Flush()
	return nil
}

// ParseAll Main method to parse all clubs of current date
func (r *ClubParse) ParseAll() error {
	var url = getURL(r.CurrentPage)
	fmt.Print(fmt.Sprintf("> %s :", url))
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	r.parseClub(doc)
	fmt.Println(" OK!")

	return nil
}

func (r *ClubParse) parseClub(doc *goquery.Document) {
	doc.Find(".sm_logo-name").Each(func(i int, s *goquery.Selection) {
		c := Club{}
		c.Name = s.Text()
		v, exist := s.Attr("href")
		if exist {
			id := strings.Split(v, "/")[2]
			c.ID = id
		}
		r.Clubs = append(r.Clubs, c)
	})
}
