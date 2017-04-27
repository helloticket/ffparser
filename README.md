# ffparser
The library is a Flat-File Parser

## Install
```bash
go get -v github.com/helderfarias/ffparser
```

## Using
```go
package main

import "github.com/helderfarias/ffparser"

type Pojo struct {
    Value1 time.Time `record:"start=1,end=21"`
    Value2 int       `record:"start=22,end=32"`
    Value3 int64     `record:"start=33,end=43"`
    Value4 float64   `record:"start=44,end=54"`
    Value5 string    `record:"start=55,end=65"`
}

func main() {
    ffp := ffparser.NewSimpleParser()

    result, _ := parser.ParseToText(&Pojo{
        Value1: time.Date(2017, 5, 10, 0, 0, 0, 0, time.UTC),
        Value2: 12402,
        Value3: 4567822222,
        Value4: 4567833.22,
        Value5: "be happy",
    })

    log.Println(result)
}
```


## Inspired
https://github.com/ffpojo/ffpojo


