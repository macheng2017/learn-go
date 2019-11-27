package parser

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"regexp"
	"strconv"
)

const ageRe = `<div data-v-8b1eac0c class="m-btn purple">([\d]+)岁</div>`
const marriageRe = `<div data-v-8b1eac0c class="m-btn purple">(.婚)</div>`

func ParserProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	re := regexp.MustCompile(ageRe)
	match := re.FindSubmatch(contents)
	if match != nil {
		age, err := strconv.Atoi(string(match[1]))
		if err != nil {
			// user age
			profile.Age = age
		}

	}

	re = regexp.MustCompile(marriageRe)
	match = re.FindSubmatch(contents)
	if match != nil {
		marriage := match[1]
		profile.Marriage = string(marriage)

	}

}
