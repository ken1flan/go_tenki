# go_tenki
go_tenki is a CLI application of weather forecast.

## How to use
### Get source and install
```
$ go get github.com/ken1flan/go_tenki
$ cd $GOPATH/github.com/ken1flan/go_tenki
$ go install
```

### [Get OpenWeatherMap API Key](http://openweathermap.org/appid)

### Copy and edit .envrc with your API key and location info
```
$ cp dot.envrc .envrc
$ vi .envrc
$ source .envrc
```

### Run
```
$ go_tenki
☀️ 2.0℃$
```
