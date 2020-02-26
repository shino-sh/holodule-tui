package main

import (
  "github.com/shino-sh/holodule-tui/holodule"
  "github.com/shino-sh/holodule-tui/imageloader"

  "github.com/gdamore/tcell"
  "github.com/rivo/tview"
  "github.com/nfnt/resize"
  "github.com/kevin-cantwell/dotmatrix"
)

func main() {
  schedules := holodule.GetSchedules()

  app := tview.NewApplication()

  list := tview.NewList().ShowSecondaryText(false)
  list.SetBorder(true).SetTitle("[holodule[]")

  linkView := tview.NewTextView()
  linkView.SetBorder(true).SetTitle("[YouTube[]")

  thumbnailView := tview.NewTextView()
  thumbnailView.SetBorder(true).SetTitle("[thumbnail[]")

  for _, s := range schedules {
    schedule := s
    list.AddItem(s.GetText(), "", 0, func() {
      linkView.Clear()
      linkView.SetTextColor(tcell.ColorRed).
          SetText(schedule.Link)

      thumbnailView.Clear()
      _, _, width, _ := thumbnailView.Box.GetRect()
      img := imageloader.Load(schedule.ThumbnailUrl)
      resizeImg := resize.Resize(uint((width - 2) * 2), 0, *img, resize.Lanczos3)
      dotmatrix.Print(thumbnailView, resizeImg)
    })
  }

  list.AddItem("Quit", "Press to exit", 'q', func() {
    app.Stop()
  })

  flex := tview.NewFlex().
      AddItem(list, 0, 1, true).
      AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
          AddItem(linkView, 3, 1, false).
          AddItem(thumbnailView, 0, 5, false), 0, 2, false)
    
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}

