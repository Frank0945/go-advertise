package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("advertise", func() {
	Title("AD Service")
	Description("This service provides the AD service")
	Server("advertise", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
		Services("advertise")
	})
})

var _ = Service("advertise", func() {
	Method("create", func() {
		Meta("openapi:summary", "Create a new AD")
		Description("Create a new edge")
		Payload(func() {
			Attribute("title", String, "Title of AD", func() {
				Example("AD 1")
			})
			Attribute("start_at", String, "Start time of AD", func() {
				Format(FormatDateTime)
				Example("2024-03-10T03:00:00.000Z")
			})
			Attribute("end_at", String, "End time of AD", func() {
				Format(FormatDateTime)
				Example("2024-12-10T03:00:00.000Z")
			})
			Attribute("conditions", func() {
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
				Attribute("country", String, "Country of target", func() {
					Enum("TW", "JP")
					Example("TW")
				})
				Attribute("platform", String, "Platform of target", func() {
					Enum("ios", "android", "web")
					Example("ios")
				})
			})

			Required("title", "start_at", "end_at")
		})
		Result(Empty)
		HTTP(func() {
			POST("ad")
			Response(StatusCreated)
		})
	})

	Method("list", func() {
		Meta("openapi:summary", "List all ADs by filter")
		Description("List all ADs by filter")

		Payload(AdList)
		Result(ArrayOf(Ads))
		HTTP(func() {
			GET("ad")

			Params(func() {
				Param("offset")
				Param("limit")
				Param("age_start")
				Param("age_end")
				Param("gender")
				Param("country")
				Param("platform")
			})

			Response(StatusOK)
		})
	})
})
