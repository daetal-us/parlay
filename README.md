# Parlay

A [golang][go] command line utility and library to generate a data URL, as specified by [RFC 2397](rfc2397), from a binary string, file path or remote URL.

## Installation

```sh
go get github.com/daetal-us/parlay/cmd/parlay
```

## Usage


### Command Line
Convert an local image to a data url and pipe to clipboard:

```sh
parlay path/to/my/photo.png
```

Convert a remote resource to a data url:

```sh
parlay https://secure.gravatar.com/avatar/245101069f239ec4a678f7eafd022045?s=1
```

### Golang

```go
package main

import (
	"log"

	"github.com/daetal-us/parlay"
)

func main() {
	getDataURLFromPath()
	getDataURLFromURL()
}

func getDataURLFromPath() {
	url, err := parlay.FromPath("parlay_test.bmp")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(url)
}

func getDataURLFromURL() {
	url, err := parlay.FromURL("https://secure.gravatar.com/avatar/245101069f239ec4a678f7eafd022045?s=1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(url)
}

```

[go]:https://golang.org
[rfc2397]://tools.ietf.org/html/rfc2397
