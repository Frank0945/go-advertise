package design

import (
	. "goa.design/goa/v3/dsl"
)

var Ad = Type("Ad", func() {
	Description("List all ads by filter")

	Field(0, "title", String, "Title of AD", func() {
		MinLength(1)
		MaxLength(100)
		Example("AD 1")
	})
	Field(1, "end_at", String, "End time of AD", func() {
		Format(FormatDateTime)
		Example("2024-12-10T03:00:00.000Z")
	})

	Required("title", "end_at")
})

var AdOverview = Type("AdOverview", func() {
	Description("Search AD payload")

	Field(0, "offset", Int, "Offset of AD", func() {
		Default(0)
		Minimum(0)
		Example(0)
	})
	Field(1, "limit", Int, "Limit of AD", func() {
		Default(5)
		Minimum(1)
		Example(10)
	})
	Field(2, "age_start", Int, "Start age of target", func() {
		Minimum(1)
		Maximum(100)
		Example(18)
	})
	Field(3, "age_end", Int, "End age of target", func() {
		Minimum(1)
		Maximum(100)
		Example(60)
	})
	Field(4, "gender", String, "Gender of target", func() {
		Enum("M", "F")
		Example("M")
	})
	Field(5, "country", String, "Country of target", func() {
		Enum("TW", "JP")
		Example("TW")
	})
	Field(6, "platform", String, "Platform of target", func() {
		Enum("ios", "android", "web")
		Example("ios")
	})

	Required("offset", "limit")
})
