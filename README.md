# kanaconv [![Build Status](https://travis-ci.org/miiton/jpconv.svg?branch=master)](https://travis-ci.org/miiton/jpconv) [![GoDoc](https://godoc.org/github.com/miiton/kanaconv?status.svg)](https://godoc.org/github.com/miiton/kanaconv)
日本語のひらがなカタカナ、全角半角を変換するパッケージ written in Go

## Installation

```
go get -u github.com/miiton/kanaconv
```

## Usage

```
import (
    "fmt"
    "github.com/miiton/kanaconv"
)

func main() {
    fmt.Println(kanaconv.HiraganaToKatakana("ひらがなをカタカナに"))
    // ヒラガナヲカタカナニ
    fmt.Println(kanaconv.KatakanaToHiragana("カタカナをひらがなに"))
    // かたかなをひらがなに
    fmt.Println(kanaconv.HankakuToZenkaku("ﾊﾝｶｸｦゼンカクﾆ"))
    // ハンカクヲゼンカクニ
    fmt.Println(kanaconv.ZenkakuToHankaku("ゼンカクヲﾊﾝｶｸニ"))
    // ｾﾞﾝｶｸｦﾊﾝｶｸﾆ
}
```

## Benchmarks

* MacBook Air (11-inch, Mid 2011)

```
% go test -bench .                                                                                                                                                                (git)-[master]
PASS
BenchmarkHiraganaToKatakana-4    1000000              1548 ns/op
BenchmarkKatakanaToHiragana-4    1000000              1501 ns/op
BenchmarkHankakuToZenkaku-4      1000000              1889 ns/op
BenchmarkZenkakuToHankaku-4      1000000              1826 ns/op
ok      github.com/miiton/kanaconv      7.797s
```

## Referenced

* [Go言語で文字列の変換(全角・半角、ひらがな・カタカナ)をする : Serendip - Webデザイン・プログラミング](http://www.serendip.ws/archives/6307)
