package goemoji

import (
	"github.com/Alliera/logging"
	"regexp"
	"strings"
)

var wsRe = regexp.MustCompile(`\s+`)

type GoEmoji struct {
	re                     *regexp.Regexp
	downloadData           bool
	dataPath               string
	referenceVersion       string
	codepointsFileName     string
	codepointsFullFilePath string
}

func New(version, dataPath, codepointsFileName string, downloadData bool) (*GoEmoji, error) {
	goe := &GoEmoji{
		downloadData:       downloadData,
		dataPath:           dataPath,
		codepointsFileName: codepointsFileName,
		referenceVersion:   version,
	}

	if err := goe.init(); err != nil {
		return nil, logging.Trace(err)
	}

	return goe, nil
}

func NewDefault(downloadData bool) (*GoEmoji, error) {
	return New("latest", "emojidata", "codepoints", downloadData)
}

func (goe *GoEmoji) Pad(emojis string, cleanExtraWhitespace bool) string {
	pad := goe.re.ReplaceAllString(emojis, ` $1 `)
	if cleanExtraWhitespace {
		pad = wsRe.ReplaceAllString(pad, " ")
	}
	return pad
}

func (goe *GoEmoji) Replace(emojis, replacementPattern string) string {
	return goe.re.ReplaceAllString(emojis, replacementPattern)
}

func (goe *GoEmoji) Words(emojis string) []string {
	var result []string
	for _, word := range goe.re.Split(emojis, -1) {
		word = strings.TrimSpace(word)
		if word != "" {
			result = append(result, word)
		}
	}
	return result
}
