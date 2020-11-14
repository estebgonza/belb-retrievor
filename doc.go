/*
Example main function to use this module for match parsing

func main() {
	fmt.Println("El Retrievor")
	// Initialize on first page
	scope := "monde"
	date := "2019-03"
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

Example main function to use this module for club parsing
func main() {
	fvar c = ClubParse{}
	var countries = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1"}
	// Iterate trough all pages
	for _, e := range countries {
		c.CurrentPage = e
		c.ParseAll()

	}
	c.ExportAsCSV()
}
*/

package retrievor
