package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strconv"
)

var nameRe = regexp.MustCompile(`<h1 class="nickName" [^>]*>([^<]+)</h1>`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)岁</div>`)
var idRe = regexp.MustCompile(`<div class="id" [^>]*>ID：\s*([\d]+)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>(.婚)</div>`)
var xingzuoRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>(.*座)[^<]+</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)cm</div>`)
var workPlaceRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>工作地:([^<]+)</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>月收入:([^<]+)</div>`)
var workRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>在校学生</div>`)
var hukouRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>籍贯:([^<]+)</div>`)
var carRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>(.买车)</div>`)
var houseRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>(.购房)</div>`)
var bodyTypeRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>体型:([^<])</div>`)

func ParserProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		// user age
		profile.Age = age
	}

	profile.Name = extractString(contents, nameRe)
	profile.Id = extractString(contents, idRe)
	profile.Marriage = extractString(contents, marriageRe)
	profile.Xingzuo = extractString(contents, xingzuoRe)
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	profile.Income = extractString(contents, incomeRe)
	profile.Hukou = extractString(contents, hukouRe)
	profile.Car = extractString(contents, carRe)
	profile.House = extractString(contents, houseRe)
	profile.BodyType = extractString(contents, bodyTypeRe)

	return engine.ParseResult{Items: []interface{}{profile}}
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
