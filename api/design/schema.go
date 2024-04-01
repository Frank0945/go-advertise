package design

import (
	. "goa.design/goa/v3/dsl"
)

var Genders = ArrayOf(String, func() {
	Pattern("M|F")
})

var Countries = ArrayOf(String, func() {
	Pattern("TW|JP")
})

var Platforms = ArrayOf(String, func() {
	Pattern("ios|android|web")
})

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

var AdQuery = Type("AdQuery", func() {
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
	Field(0, "gender", Genders, "Gender of target", func() {
		Example([]string{"M", "F"})
	})
	Field(1, "country", Countries, "Country of target", func() {
		Example([]string{"TW", "JP"})
	})
	Field(2, "platform", Platforms, "Platform of target", func() {
		Example([]string{"ios", "android", "web"})
	})

	Required("offset", "limit")
})
