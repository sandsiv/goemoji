package main

import (
	"fmt"
	"github.com/Alliera/logging"
	"io"
	"net/http"
	"os"
)

const (
	defaultCodePointsFilePath = "emoji_data/codepoints.txt"
	defaultReferenceFilePath  = "emoji_data/reference.txt"
	defaultReferenceVersion   = "latest"
)

func SaveReferenceFile(version, filePath string) error {
	if version != "" {
		version = defaultReferenceVersion
	}
	if filePath != "" {
		filePath = defaultReferenceFilePath
	}
	f, err := os.Create(filePath)
	if err != nil {
		return logging.Trace(err)
	}
	defer func() { _ = f.Close() }()
	resp, err := http.Get(fmt.Sprintf("http://unicode.org/Public/emoji/%s/emoji-test.txt", version))
	if err != nil {
		return logging.Trace(err)
	}
	defer func() { _ = resp.Body.Close() }()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return logging.Trace(err)
	}
	return nil
}
