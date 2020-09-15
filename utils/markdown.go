package utils

import (
	"html/template"
	"regexp"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func MarkdownToHtml(mdText string) template.HTML {
	mdText = strings.Replace(mdText, "\r\n", "\n", -1)

	unsafe := blackfriday.Run([]byte(mdText))
	html := bluemonday.UGCPolicy().AllowAttrs("class").
		Matching(regexp.MustCompile("^language-[a-zA-Z0-9]+$")).OnElements("code").
		AllowAttrs("src").OnElements("img").
		AllowAttrs("class", "target", "id", "style", "align").Globally().
		AllowAttrs("src", "width", "height", "border", "marginwidth", "marginheight").OnElements("iframe").
		AllowAttrs("controls", "src").OnElements("audio").
		AllowAttrs("color").OnElements("font").
		AllowAttrs("controls", "src", "width", "height").OnElements("video").
		AllowAttrs("src", "media", "type").OnElements("source").
		AllowAttrs("width", "height", "data", "type").OnElements("object").
		AllowAttrs("name", "value").OnElements("param").
		AllowAttrs("src", "type", "width", "height", "wmode", "allowNetworking").OnElements("embed").
		SanitizeBytes(unsafe)
	return template.HTML(html)
}
