package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No filepath specified.")
		os.Exit(1)
	}
	path := os.Args[1]
	if path == "" {
		fmt.Println("No filepath specified.")
		os.Exit(1)
	}
	img, err := ioutil.ReadFile(os.Args[1])
	if err == nil {
		mimeType := strings.Replace(
			mime.TypeByExtension(filepath.Ext(os.Args[1])), " ", "", -1,
		)
		encodedData := base64.StdEncoding.EncodeToString(img)
		fmt.Println(fmt.Sprintf(`data:%s;base64,%s`, mimeType, encodedData))
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
