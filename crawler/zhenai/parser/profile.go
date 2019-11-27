package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div data-v-8b1eac0c class="m-btn purple">([\d]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div data-v-8b1eac0c class="m-btn purple">(.婚)</div>`)

func ParserProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		// user age
		profile.Age = age
	}

	profile.Marriage = extractString(contents, marriageRe)

}

// 将重复的代码抽离出来
func extractString(contents []byte, re *regexp.Regexp) string {
	// 在函数内部使用MustCompile比较浪费时间,提到外部
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
