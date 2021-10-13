# [kkconv] -- Hiragana-Katakana Conversion

[![check vulns](https://github.com/spiegel-im-spiegel/kkconv/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/kkconv/actions)
[![lint status](https://github.com/spiegel-im-spiegel/kkconv/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/kkconv/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/kkconv/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/kkconv.svg)](https://github.com/spiegel-im-spiegel/kkconv/releases/latest)

This package is required Go 1.16 or later.

## Import package

```go
import "github.com/spiegel-im-spiegel/kkconv"
```

## Usage

```go
import (
    "fmt"

    "github.com/spiegel-im-spiegel/kkconv"
)

func ExampleHiragana() {
    txt := "こんにちは ｾｶｲ"
    fmt.Println(kkconv.Hiragana(txt, true))
    // Output:
    // こんにちは せかい
}

func ExampleKatakana() {
    txt := "こんにちは ｾｶｲ"
    fmt.Println(kkconv.Katakana(txt, true))
    // Output:
    // コンニチハ セカイ
}
```

## Modules Requirement Graph

[![dependency.png](./dependency.png)](./dependency.png)

[kkconv]: https://github.com/spiegel-im-spiegel/kkconv "spiegel-im-spiegel/kkconv: Hiragana-Katakana Conversion"
