package parser

import (
	"log"
	"regexp"
	"strconv"
	"web-spider/engine"
	"web-spider/model"
)

var nameRe = regexp.MustCompile(`<a class="name fs24">([^<]+)</a>`)
var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)

// ParserUserProfile 解析用户
func ParserUserProfile(contents []byte) engine.ParserResult {

	userProfile := model.UserProfile{}
	userProfile.Name = extractString(contents, nameRe)
	userProfile.Age, _ = strconv.Atoi(extractString(contents, ageRe))
	userProfile.Height, _ = strconv.Atoi(extractString(contents, heightRe))
	userProfile.Gender = extractString(contents, genderRe)
	result := engine.ParserResult{
		Items: []interface{}{userProfile},
	}
	log.Printf("get User: %s %s %d %d \n", userProfile.Name, userProfile.Gender, userProfile.Age, userProfile.Height)
	return result
}
func extractString(contents []byte, re *regexp.Regexp) string {

	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""

}
