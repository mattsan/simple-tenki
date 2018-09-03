package main

import (
    "fmt"
    "tenki"

    "github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
    Message string `json:"message"`
}

func Handler() (Response, error) {
    var data tenki.Data
    err := data.Fetch()

    if err != nil {
        fmt.Printf("ERROR: %v\n", err)
    } else {
        for i, weather := range data.Weather {
            fmt.Printf(`No.%d
           id: %d
         main: %s
  description: %s
`, i, weather.Id, weather.Main, weather.Description)
        }
    }

    return Response{
        Message: "Go Serverless v1.0! Your function executed successfully!",
    }, nil
}

func main() {
    lambda.Start(Handler)
}
