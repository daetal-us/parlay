<img src="http://daetal.us/static/media/dataurl.png" align="right" width="35%">

# Parley
_A command line utility to generate data URLs for images_

This golang powered command line utility generates a data URL, as specified by [RFC 2397](//tools.ietf.org/html/rfc2397), for a given file.

## Installation

```bash
go get github.com/daetal-us/parley
```

## Usage

Convert an image to a data url and pipe to clipboard:

```bash
parlay path/to/my/photo.png | pbcopy
```
