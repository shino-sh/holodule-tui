package holodule

import (
  "strings"

  "github.com/PuerkitoBio/goquery"
)

const (
	url string = "https://schedule.hololive.tv/"
)

type Schedule struct {
  Date string
  Time string
  Name string
  Link string
  ThumbnailUrl string
}

func (s *Schedule) GetText() string {
  return strings.Join([]string{s.Date, s.Time, s.Name}, " ")
}

func GetSchedules() []*Schedule {
  doc, err := goquery.NewDocument(url)

  if err != nil {
    panic(err)
  }

  rep := strings.NewReplacer(
    " ", "", 
    "\n", "")

  var schedules []*Schedule
  var date string

  hololiveSelection := doc.Find("div#hololive")
  selections := hololiveSelection.Find("div.holodule, a.thumbnail")
  selections.Each(func(index int, s *goquery.Selection) {
    if s.HasClass("thumbnail") {
      time := strings.TrimSpace(s.Find("div.datetime").Text())
      name := strings.TrimSpace(s.Find("div.name").Text())
      link, _ := s.Attr("href")
      thumbnailUrl, _ := s.Find("img[src*='img.youtube.com']").Attr("src")
      schedules = append(schedules, &Schedule{date, time, name, link, thumbnailUrl})
    } else {
      date = rep.Replace(s.Text())
    }
  })

  return schedules
}
