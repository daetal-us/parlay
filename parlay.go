package parlay

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

func FromURL(u string) (string, error) {
	r, err := http.Get(u)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return "", errors.New(r.Status)
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	m := r.Header.Get("Content-Type")
	if m == "" {
		return "", errors.New("Unspecified Content-Type.")
	}
	return FromBytes(b, m), nil
}

func FromPath(path string) (string, error) {
	if path == "" {
		return "", errors.New("No path specified.")
	}
	b, err := ioutil.ReadFile(path)
	if err == nil {
		m := strings.Replace(
			mime.TypeByExtension(filepath.Ext(path)), " ", "", -1,
		)
		return FromBytes(b, m), nil
	}
	return "", err
}

func FromBytes(data []byte, mimeType string) string {
	encoded := base64.StdEncoding.EncodeToString(data)
	return fmt.Sprintf(`data:%s;base64,%s`, mimeType, encoded)
}
