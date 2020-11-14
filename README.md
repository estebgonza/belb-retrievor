# belb-retrievor

# Examples

## Matches

You can parse results of all match from 2020-01 with

```
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
```

## Clubs

Yu can parse and export all clubs with

```
var c = ClubParse{}
var countries = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1"}
// Iterate trough all pages
for _, e := range countries {
    c.CurrentPage = e
    c.ParseAll()

}
c.ExportAsCSV()
```