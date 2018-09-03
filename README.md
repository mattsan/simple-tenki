# 天気をログに出力する

[OpenWeatherMap](https://openweathermap.org) の API を利用して天気を取得し、CloudWatch Logs に出力するサンプル。

## 使い方

git clone します。

```sh
$ git clone git@github.com:mattsan/simple-tenki.git
$ cd simple-tenki
```

必要であれば `serverless.yml` の環境変数の設定を編集します。


デプロイします。このとき環境変数 `APPID` に OpenWeatherMap の API Key を指定します。
デプロイにはビルド、ビルドにはパッケージ取得を依存関係として設定しているので、デプロイを実行するとパッケージ取得とビルドも実行されます。

```
$ APPID=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX make deploy
GOPATH=`pwd` go get github.com/aws/aws-lambda-go/lambda
GOPATH=`pwd` GOOS=linux go build -o bin/main
sls deploy
...
Service Information
service: simple-tenki
stage: dev
region: ap-northeast-1
stack: simple-tenki-dev
api keys:
  None
endpoints:
  None
functions:
  weather: simple-tenki-dev-weather
```

ログを確認します。

```sh
$ aws logs filter-log-events --log-group-name /aws/lambda/simple-tenki-dev-weather | jq .events[].message -r
START RequestId: f78bd5ba-af6e-11e8-82db-b7fd5f22122d Version: $LATEST

No.0

           id: 801

         main: Clouds

  description: few clouds

END RequestId: f78bd5ba-af6e-11e8-82db-b7fd5f22122d

REPORT RequestId: f78bd5ba-af6e-11e8-82db-b7fd5f22122d	Duration: 145.13 ms	Billed Duration: 200 ms 	Memory Size: 1024 MB	Max Memory Used: 28 MB

...
```

削除します。

```sh
$ make remove
```
