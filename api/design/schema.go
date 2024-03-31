package design

import (
	. "goa.design/goa/v3/dsl"
)

var Ads = Type("Ads", func() {
	Description("List all ads by filter")

	Field(0, "title", String, "Title of AD", func() {
		Example("AD 1")
	})
	Field(1, "end_at", String, "End time of AD", func() {
		Example("2024-10-01 00:00:00")
	})

	Required("title", "end_at")
})

var AdList = Type("AdList", func() {
	Description("Search AD payload")

	Field(0, "offset", Int, "Offset of AD", func() {
		Example(0)
	})
	Attribute("limit", Int, "Limit of AD", func() {
		Example(10)
	})
	Attribute("age_start", Int, "Start age of target", func() {
		Minimum(1)
		Example(18)
	})
	Attribute("age_end", Int, "End age of target", func() {
		Maximum(100)
		Example(60)
	})
	Attribute("gender", String, "Gender of target", func() {
		Enum("M", "F")
		Example("M")
	})
	Attribute("Country", String, "Country of target", func() {
		Enum("TW", "JP")
		Example("TW")
	})
	Attribute("platform", String, "Platform of target", func() {
		Enum("ios", "android", "web")
		Example("ios")
	})

	Required("offset", "limit")
})
