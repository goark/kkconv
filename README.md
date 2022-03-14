# [kkconv] -- Hiragana-Katakana Conversion

[![check vulns](https://github.com/goark/kkconv/workflows/vulns/badge.svg)](https://github.com/goark/kkconv/actions)
[![lint status](https://github.com/goark/kkconv/workflows/lint/badge.svg)](https://github.com/goark/kkconv/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/goark/kkconv/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/kkconv.svg)](https://github.com/goark/kkconv/releases/latest)

This package is required Go 1.16 or later.

**Migrated repository to [github.com/goark/kkconv][kkconv]**

## Import package

```go
import "github.com/goark/kkconv"
```

## Usage

```go
import (
    "fmt"

    "github.com/goark/kkconv"
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

[kkconv]: https://github.com/goark/kkconv "goark/kkconv: Hiragana-Katakana Conversion"
