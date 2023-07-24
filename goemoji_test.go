package goemoji

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func cleanUpAllTempDirs() {
	_ = os.Remove("testdatatemp/codepoints.txt")
	_ = os.Remove("testdatatemp")
	_ = os.Remove("emojidata/codepoints.txt")
	_ = os.Remove("emojidata")
}

type emojiPadTestData struct {
	raw                   string
	pad                   string
	separated             string
	partsLen              int
	removeExtraWhitespace bool
}

var emojiPadTestDataList = []emojiPadTestData{
	{
		raw:                   "hello😊World!😄🌎🏴󠁧󠁢󠁷󠁬󠁳󠁿🏳️‍🌈",
		pad:                   "hello 😊 World! 😄 🌎 🏴󠁧󠁢󠁷󠁬󠁳󠁿 🏳️‍🌈 ",
		separated:             "hello|😊|World!|😄|🌎|🏴󠁧󠁢󠁷󠁬󠁳󠁿|🏳️‍🌈|",
		partsLen:              8,
		removeExtraWhitespace: true,
	},
	{
		raw:                   "I ❤️ coding!👍Let's build something amazing!🚀     🌟",
		pad:                   "I ❤️ coding! 👍 Let's build something amazing! 🚀 🌟 ",
		separated:             "I|❤️|coding!|👍|Let's|build|something|amazing!|🚀|🌟|",
		partsLen:              11,
		removeExtraWhitespace: true,
	},
	{
		raw:                   "Good morning!  ☀️   It's a new day!🎉Let's make the most of it!💪😃  ",
		pad:                   "Good morning!   ☀️    It's a new day! 🎉 Let's make the most of it! 💪  😃   ",
		separated:             "Good|morning!|||☀️||||It's|a|new|day!|🎉|Let's|make|the|most|of|it!|💪||😃|||",
		partsLen:              25,
		removeExtraWhitespace: false,
	},
	{
		raw:                   "That joke was hilarious!😂😂😂 Bravo!👏👏👏",
		pad:                   "That joke was hilarious! 😂 😂 😂 Bravo! 👏 👏 👏 ",
		separated:             "That|joke|was|hilarious!|😂|😂|😂|Bravo!|👏|👏|👏|",
		partsLen:              12,
		removeExtraWhitespace: true,
	},
	{
		raw:                   "😄🌎🏴󠁧󠁢󠁷󠁬󠁳󠁿🏳️‍🌈😂😂😂❤️🚀🌟",
		pad:                   " 😄 🌎 🏴󠁧󠁢󠁷󠁬󠁳󠁿 🏳️‍🌈 😂 😂 😂 ❤️ 🚀 🌟 ",
		separated:             "|😄|🌎|🏴󠁧󠁢󠁷󠁬󠁳󠁿|🏳️‍🌈|😂|😂|😂|❤️|🚀|🌟|",
		partsLen:              12,
		removeExtraWhitespace: true,
	},
	{
		raw:                   "🏳‍⚧️🏴‍☠️🇦🇫🏳️‍🌈🎌",
		pad:                   " 🏳‍⚧️ 🏴‍☠️ 🇦🇫 🏳️‍🌈 🎌 ",
		separated:             "|🏳‍⚧️|🏴‍☠️|🇦🇫|🏳️‍🌈|🎌|",
		partsLen:              7,
		removeExtraWhitespace: true,
	},
	{
		raw:                   "👨🏾‍❤️‍💋‍👨🏻👨🏽‍❤‍💋‍👨🏿🏳️‍🌈🏳‍⚧️",
		pad:                   " 👨🏾‍❤️‍💋‍👨🏻 👨🏽‍❤‍💋‍👨🏿 🏳️‍🌈 🏳‍⚧️ ",
		separated:             "|👨🏾‍❤️‍💋‍👨🏻|👨🏽‍❤‍💋‍👨🏿|🏳️‍🌈|🏳‍⚧️|",
		partsLen:              6,
		removeExtraWhitespace: true,
	},
}

func TestPad(t *testing.T) {
	goe, _ := NewDefault(true)
	for _, testData := range emojiPadTestDataList {
		pad := goe.Pad(testData.raw, testData.removeExtraWhitespace)
		assert.Equal(t, testData.pad, pad, "wrong pad value")
		assert.Equal(t, testData.partsLen, len(strings.Split(pad, " ")), "wrong parts len")
		assert.Equal(t, testData.separated, strings.Join(strings.Split(pad, " "), "|"), "wrong separated value")
	}
}

type emojiWordsTestData struct {
	raw   string
	split []string
}

var emojiWordsTestDataList = []emojiWordsTestData{
	{
		raw:   "hello😊World!😄🌎🏴󠁧󠁢󠁷󠁬󠁳󠁿🏳️‍🌈",
		split: []string{"hello", "World!"},
	},
	{
		raw:   "I ❤️ coding!👍  Let's build something  amazing!🚀     🌟",
		split: []string{"I", "coding!", "Let's build something  amazing!"},
	},
	{
		raw:   "  Good morning!  ☀️     It's a new day!🎉Let's make the most of it!💪😃",
		split: []string{"Good morning!", "It's a new day!", "Let's make the most of it!"},
	},
	{
		raw:   "That joke was hilarious!😂😂😂 Bravo!👏  👏👏",
		split: []string{"That joke was hilarious!", "Bravo!"},
	},
}

func TestWords(t *testing.T) {
	goe, _ := NewDefault(true)
	for _, testData := range emojiWordsTestDataList {
		assert.Equal(t, testData.split, goe.Words(testData.raw))
	}
}

type emojiReplaceTestData struct {
	raw         string
	replacement string
	replaced    string
}

var emojiReplaceTestDataList = []emojiReplaceTestData{
	{
		raw:         "hello😊World!😄🌎🏴󠁧󠁢󠁷󠁬󠁳󠁿🏳️‍🌈",
		replacement: "",
		replaced:    "helloWorld!",
	},
	{
		raw:         "I ❤️ coding!👍  Let's build something  amazing!🚀     🌟",
		replacement: "$",
		replaced:    "I $ coding!$  Let's build something  amazing!$     $",
	},
	{
		raw:         "  Good morning!  ☀️     It's a new day!🎉Let's make the most of it!💪😃",
		replacement: "XXX",
		replaced:    "  Good morning!  XXX     It's a new day!XXXLet's make the most of it!XXXXXX",
	},
	{
		raw:         "That joke was hilarious!😂😂😂 Bravo!👏  👏👏",
		replacement: "    ",
		replaced:    "That joke was hilarious!             Bravo!              ",
	},
}

func TestReplace(t *testing.T) {
	goe, _ := NewDefault(true)

	for _, testData := range emojiReplaceTestDataList {
		assert.Equal(t, testData.replaced, goe.Replace(testData.raw, testData.replacement))
	}
}

func Test_buildPattern(t *testing.T) {
	goe := &GoEmoji{codepointsFullFilePath: "testdata/codepoints_selection_unsorted.txt"}
	pattern, _ := goe.buildPattern()
	assert.NotEqual(t, "(😀|😃|😄|😁|😆|😅|🙎‍♂️)", pattern, "emoji was not sorted by descending order, it must be sorted to build correct regexp pattern")
	assert.Equal(t, "(🙎‍♂️|😀|😃|😄|😁|😆|😅)", pattern, "bad regexp pattern")
}

type rawCodePointsTetData struct {
	raw     string
	badPart string
	emoji   string
	lineN   int
	isError bool
}

var rawCodePointsTestDataList = []rawCodePointsTetData{
	{
		"   \n1F600\n\r",
		"",
		"😀",
		0,
		false,
	},
	{
		"1F- f_34w5603",
		"1F-",
		"",
		2,
		true,
	},
	{
		"1F471 1F3FB 200D 2640 FE0F",
		"",
		"👱🏻‍♀️",
		0,
		false,
	},
	{
		"\n\n           1F64D 1F3FB 200D 2640 FE0F\n\t\r",
		"",
		"🙍🏻‍♀️",
		0,
		false,
	},
	{
		"1F471 1F3FB_f2640 FE0F",
		"1F3FB_f2640",
		"",
		5,
		true,
	},
	{
		"someLiteral",
		"someLiteral",
		"",
		100,
		true,
	},
	{
		"1F469 200D 2764 FE0F 200D 1F48B 200D 1F468",
		"",
		"👩‍❤️‍💋‍👨",
		0,
		false,
	},
	{
		"1F1FA 1F1E6",
		"",
		"🇺🇦",
		0,
		false,
	},
	{
		"1F1FA 1F1E6_",
		"1F1E6_",
		"",
		5,
		true,
	},
}

func Test_convRawCodepointsToEmoji(t *testing.T) {
	for _, testData := range rawCodePointsTestDataList {
		emoji, err := convRawCodepointsToEmoji(testData.raw, testData.lineN)
		if testData.isError {
			assert.Equal(t, fmt.Sprintf("parser error: failed to parse reference, bad raw codepoints, line:%d -> (%s): strconv.ParseInt: parsing \"%s\": invalid syntax", testData.lineN, testData.raw, testData.badPart), err.Error())
		} else {
			assert.Equal(t, testData.emoji, emoji)
		}
	}
}

func Test_parseReferenceErr(t *testing.T) {
	referenceData := `# emoji-test.txt
# Date: 2020-01-21, 13:40:25 GMT
# © 2020 Unicode®, Inc.
# Unicode and the Unicode Logo are registered trademarks of Unicode, Inc. in the U.S. and other countries.
# For terms of use, see http://www.unicode.org/terms_of_use.html
#
# Emoji Keyboard/Display Test Data for UTS #51
# Version: 13.0
#

# group: Smileys & Emotion

# subgroup: face-smiling
1F600                                      ; fully-qualified     # 😀 E1.0 grinning face
1F-9sd+_34w5603                                      ; fully-qualified     # 😃 E0.6 grinning face with big eyes
1F601                                      ; fully-qualified     # 😁 E0.6 beaming face with smiling eyes
`
	goe := &GoEmoji{}
	_, err := goe.parseReference([]byte(referenceData))
	assert.Equal(t, "parser error: failed to parse reference, bad raw codepoints, line:15 -> (1F-9sd+_34w5603): strconv.ParseInt: parsing \"1F-9sd+_34w5603\": invalid syntax", err.Error())
}

func Test_parseReferenceOk(t *testing.T) {
	referenceData := `# emoji-test.txt
# Date: 2020-01-21, 13:40:25 GMT
# © 2020 Unicode®, Inc.
# Unicode and the Unicode Logo are registered trademarks of Unicode, Inc. in the U.S. and other countries.
# For terms of use, see http://www.unicode.org/terms_of_use.html
#
# Emoji Keyboard/Display Test Data for UTS #51
# Version: 13.0
#

# group: Smileys & Emotion

# subgroup: face-smiling
1F600                                      ; fully-qualified     # 😀 E1.0 grinning face
1F603                                      ; fully-qualified     # 😃 E0.6 grinning face with big eyes
1F604                                      ; fully-qualified     # 😄 E0.6 grinning face with smiling eyes
1F601                                      ; fully-qualified     # 😁 E0.6 beaming face with smiling eyes
1F606                                      ; fully-qualified     # 😆 E0.6 grinning squinting face
1F605                                      ; fully-qualified     # 😅 E0.6 grinning face with sweat
`
	goe := &GoEmoji{}
	parsedReference, _ := goe.parseReference([]byte(referenceData))
	assert.Equal(t, "😀\n😃\n😄\n😁\n😆\n😅", parsedReference, "reference data was parsed incorrectly")
}

func Test_hasFile(t *testing.T) {
	goe := &GoEmoji{}
	assert.Equal(t, nil, os.MkdirAll("somepath", 0750), "failed to create test dir test")
	assert.Equal(t, nil, os.WriteFile("somepath/somefile.txt", []byte("hi!"), 0666), "failed to create test file")

	assert.True(t, goe.hasFile("somepath/somefile.txt"))

	assert.Equal(t, nil, os.Remove("somepath/somefile.txt"), "failed to remove test file")
	assert.Equal(t, nil, os.Remove("somepath"), "failed to remove test dir")

	assert.False(t, goe.hasFile("somepath/somefile.txt"))
}

func TestNewDefault(t *testing.T) {
	defer cleanUpAllTempDirs()
	goe, _ := NewDefault(true)
	assert.Equal(t, "latest", goe.referenceVersion)
	assert.Equal(t, "emojidata", goe.dataPath)
	assert.Equal(t, "codepoints", goe.codepointsFileName)
	assert.Equal(t, "emojidata/codepoints.txt", goe.codepointsFullFilePath)
	assert.NotEqual(t, nil, goe.re)
}

func TestBadReferenceURLErr(t *testing.T) {
	defer cleanUpAllTempDirs()
	_, err := New("unknown", "testdatatemp", "codepoints", true)
	assert.Equal(t, "failed to download emoji reference (version = unknown): 404 Not Found", err.Error(), "missed error in case of bad reference URL")
}

func TestNoDataErr(t *testing.T) {
	_, err := NewDefault(false)
	assert.EqualError(t, err, "file with emoji codepoints not found, please place it manually, or pass downloadData param as true", "wrong error")
}

func TestDownloadDataOnlyOnce(t *testing.T) {
	defer cleanUpAllTempDirs()
	goe, err := NewDefault(true)
	fi, _ := os.Stat(goe.codepointsFullFilePath)
	assert.Equal(t, false, os.IsNotExist(err), "data was not downloaded")
	timeStamp := fi.ModTime()

	// simulate restart of the application
	goe, err = NewDefault(true)
	fi, _ = os.Stat(goe.codepointsFullFilePath)
	assert.Equal(t, fi.ModTime(), timeStamp, "data was updated and downloaded again, it must be downloaded only once")
}

func TestNew(t *testing.T) {
	defer cleanUpAllTempDirs()

	_, err := os.Stat("testdatatemp/codepoints.txt")
	assert.True(t, os.IsNotExist(err), "test data already exist, it must be deleted before test")

	goe, err := New("13.0", "testdatatemp", "codepoints", true)
	assert.Equal(t, nil, err, "failed to init goemoji")

	assert.Equal(t, "13.0", goe.referenceVersion)
	assert.Equal(t, "testdatatemp", goe.dataPath)
	assert.Equal(t, "codepoints", goe.codepointsFileName)
	assert.Equal(t, "testdatatemp/codepoints.txt", goe.codepointsFullFilePath)
	assert.NotEqual(t, nil, goe.re)

	_, err = os.Stat(goe.codepointsFullFilePath)
	assert.False(t, os.IsNotExist(err), "test data was not created by goemoji")

	codepointsData, _ := os.ReadFile(goe.codepointsFullFilePath)
	data, _ := os.ReadFile("testdata/v13.0_codepoints.txt")
	assert.Equal(t, string(data), string(codepointsData), "corrupted codepoints in file")
}
