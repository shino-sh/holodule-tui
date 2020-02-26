package imageloader

import (
  "image"
  "net/http"
)

func Load(url string) *image.Image {
  res, err := http.Get(url)
  if err != nil || res.StatusCode != 200 {
    panic(err)
  }
  defer res.Body.Close()

  img, _, err := image.Decode(res.Body)
  if err != nil {
    panic(err)
  }

  return &img
}
