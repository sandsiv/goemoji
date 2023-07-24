package goemoji

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/Alliera/logging"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (goe *GoEmoji) hasFile(name string) bool {
	if _, err := os.Stat(name); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func (goe *GoEmoji) init() error {
	goe.codepointsFullFilePath = filepath.Join(goe.dataPath, goe.codepointsFileName) + ".txt"

	if !goe.downloadData && !goe.hasFile(goe.codepointsFullFilePath) {
		return logging.Trace(fmt.Errorf("file with emoji codepoints not found, please place it manually, or pass downloadData param as true"))
	}
	if goe.downloadData && !goe.hasFile(goe.codepointsFullFilePath) {
		if _, err := os.Stat(goe.dataPath); os.IsNotExist(err) {
			if err = os.MkdirAll(goe.dataPath, 0750); err != nil {
				return logging.Trace(err)
			}
		}
		body, err := goe.loadReference()
		if err != nil {
			return logging.Trace(err)
		}

		parsedReference, err := goe.parseReference(body)
		if err != nil {
			return logging.Trace(err)
		}

		f, err := os.Create(goe.codepointsFullFilePath)
		if err != nil {
			return logging.Trace(err)
		}
		defer func() { _ = f.Close() }()

		if _, err = f.Write([]byte(parsedReference)); err != nil {
			return logging.Trace(err)
		}
	}

	if err := goe.initRe(); err != nil {
		return logging.Trace(err)
	}

	return nil
}

func (goe *GoEmoji) loadReference() ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("http://unicode.org/Public/emoji/%s/emoji-test.txt", goe.referenceVersion))
	if err != nil {
		return nil, logging.Trace(err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return nil, logging.Trace(fmt.Errorf("failed to download emoji reference (version = %s): %s", goe.referenceVersion, resp.Status))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, logging.Trace(fmt.Errorf("failed to read emoji reference (version = %s): %s", goe.referenceVersion, err))
	}
	defer func() { _ = resp.Body.Close() }()

	return body, nil
}

func (goe *GoEmoji) parseReference(referenceData []byte) (string, error) {

	fileScanner := bufio.NewScanner(bytes.NewBuffer(referenceData))
	fileScanner.Split(bufio.ScanLines)

	var emojis []string
	lineN := 0
	for fileScanner.Scan() {
		lineN++
		l := fileScanner.Text()
		if len(l) == 0 || l[0] == '#' {
			continue
		}
		if split := strings.Split(l, ";"); len(split) > 0 {
			emoji, err := convRawCodepointsToEmoji(split[0], lineN)
			if err != nil {
				return "", logging.Trace(err)
			}
			emojis = append(emojis, emoji)
		}

	}
	return strings.Join(emojis, "\n"), nil
}

func convRawCodepointsToEmoji(rawCodePoints string, lineN int) (string, error) {
	rawCodePoints = strings.TrimSpace(rawCodePoints)

	codePoints := strings.Split(rawCodePoints, " ")
	var emojis []rune
	for _, code := range codePoints {
		val, err := strconv.ParseInt(code, 16, 32)
		if err != nil {
			return "", fmt.Errorf("parser error: failed to parse reference, bad raw codepoints, line:%d -> (%s): %s", lineN, strings.Join(codePoints, " "), err)
		}
		emojis = append(emojis, rune(val))
	}
	return string(emojis), nil
}
