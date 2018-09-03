package tenki

import (
    "fmt"
    "os"
    "encoding/json"
    "net/http"
)

type Weather struct {
    Id          int    `json:"id"`
    Main        string `json:main"`
    Description string `json:"description"`
}

type Data struct {
    Weather []Weather `json:"weather"`
}

func tenkiUrl() string {
    appId   := os.Getenv("APPID")
    baseUrl := os.Getenv("BASE_URL")
    city    := os.Getenv("CITY")
    units   := os.Getenv("UNITS")

  return fmt.Sprintf("%s?q=%s&units=%s&appid=%s", baseUrl, city, units, appId)
}

func (data *Data) Fetch() error {
    resp, err := http.Get(tenkiUrl())

    if err != nil { return err }

    defer resp.Body.Close()

    decoder := json.NewDecoder(resp.Body)
    return decoder.Decode(data)
}
