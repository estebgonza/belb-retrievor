package main

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

const urlPattern = "https://footballdatabase.com/clubs-list-letter/%s"

var countries = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1"}

type ResultParse struct {
	currentPage string
	Clubs       []Club
}

type Club struct {
	id   string
	name string
}

func main() {
	fmt.Println("El Retrievor")
	// Iterate trough all pages
	r := ResultParse{}
	for _, e := range countries {
		r.currentPage = e
		r.parseAll()

	}
	r.exportAsCSV()
}

func getUrl(page string) string {
	return fmt.Sprintf(urlPattern, page)
}

func (r *ResultParse) exportAsCSV() error {
	f, err := os.Create(fmt.Sprintf("clubs-%d.csv", time.Now().Second()))
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, e := range r.Clubs {
		line := fmt.Sprintf("\"%s\",\"%s\"\n", e.id, e.name)
		w.WriteString(line)
	}
	w.Flush()
	return nil
}

func (r *ResultParse) parseAll() error {
	var url = getUrl(r.currentPage)
	fmt.Print(fmt.Sprintf("> %s :", url))
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	r.parseClub(doc)
	fmt.Println(" OK!")

	return nil
}

func (r *ResultParse) parseClub(doc *goquery.Document) {
	doc.Find(".sm_logo-name").Each(func(i int, s *goquery.Selection) {
		c := Club{}
		c.name = s.Text()
		v, exist := s.Attr("href")
		if exist {
			id := strings.Split(v, "/")[2]
			c.id = id
		}
		r.Clubs = append(r.Clubs, c)
	})
}
