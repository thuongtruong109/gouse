# Getting Started

## 📦 Installation

> Compatibility with Go version >= **`1.18`**

```go
go get github.com/thuongtruong109/gouse
```

## 🕯️ Quick Start

```go
package main

import "github.com/thuongtruong109/gouse"

func main() {
    gouse.Stater()
}
```

## 🦄 Usage

- Using package directly in your module as ultra-lightweight utility functions.
- Import as normal package

```go
package main

import (
    "fmt"
    "github.com/thuongtruong109/gouse/math"
)

func main() {
    fmt.Println(math.Add(1, 2))
}
```

- Please refer to the [`receipts`](/receipts/array) for more examples.