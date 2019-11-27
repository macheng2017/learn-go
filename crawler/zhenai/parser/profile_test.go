package parser

import (
	"io/ioutil"
	"learngo/crawler/model"
	"testing"
)

func TestParserProfile(t *testing.T) {
	contents, e := ioutil.ReadFile("profile_test_data.html")

	if e != nil {
		panic(e)
	}
	result := ParserProfile(contents)
	if len(result.Items) != 1 {
		t.Errorf("items should contain 1 element; but was %v ", result.Items)
	}

	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
		Name:       "小凌子",
		Id:         "1379198021",
		Gender:     "",
		Age:        25,
		Height:     160,
		Weight:     0,
		Income:     "3-5千",
		Marriage:   "未婚",
		Education:  "",
		Occupation: "",
		Hukou:      "江西吉安",
		Xingzuo:    "魔羯座",
		House:      "",
		Car:        "未买车",
		BodyType:   "",
	}
	if profile != expected {
		t.Errorf("expected %v; but was %v", expected, profile)
	}

}
