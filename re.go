package goemoji

import (
	"github.com/Alliera/logging"
	"os"
	"regexp"
	"sort"
	"strings"
)

func (goe *GoEmoji) initRe() error {
	if goe.re == nil {
		var err error
		pattern, err := goe.buildPattern()
		if err != nil {
			return logging.Trace(err)
		}
		goe.re, err = regexp.Compile(pattern)
		if err != nil {
			return logging.Trace(err)
		}
	}
	return nil
}

func (goe *GoEmoji) buildPattern() (string, error) {
	emojiFileData, err := os.ReadFile(goe.codepointsFullFilePath)
	if err != nil {
		return "", logging.Trace(err)
	}

	codepoints := strings.Fields(string(emojiFileData))

	sort.Slice(codepoints, func(i, j int) bool {
		return len(codepoints[i]) > len(codepoints[j])
	})
	return `(` + strings.Join(escapeStrings(codepoints), "|") + `)`, nil
}

func escapeStrings(strings []string) []string {
	escapedStrings := make([]string, len(strings))
	for i, s := range strings {
		escapedStrings[i] = regexp.QuoteMeta(s)
	}
	return escapedStrings
}
