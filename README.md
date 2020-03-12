# Nanoleaf Go API

Nanoleaf API made in Go (GoLang)

## Contents

- [Installation](#installation)
- [Usage](#usage)
- [Dependencies](#dependencies)

## Installation

```sh
$ go get github.com/adnanbrq/nanoleaf
```

## Usage

```go
package main

import (
  "github.com/adnanbrq/nanoleaf"
  "os"
)

func main() {
  url := os.Getenv("nanoleaf_url")
  nano := nanoleaf.NewNanoleaf(url)

  if err := nano.Identity.Flash(); err != nil {
		panic(err)
  }
  
  if err := nano.Effects.Set("Flames"); err != nil {
    panic(err)
  }
}
```

## Dependencies

- [github.com/go-resty](https://github.com/go-resty/resty)
To Communicate with the Nanoleaf API