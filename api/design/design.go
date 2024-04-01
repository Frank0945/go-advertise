package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("manager", func() {
	Title("Manager Service")
	Description("This service manages the ADs")
	Server("manager", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
		Services("manager")
	})
})

var _ = Service("manager", func() {
	Method("create_ad", func() {
		Meta("openapi:summary", "Create a new AD")
		Description("Create a new edge")
		Payload(func() {
			Attribute("title", String, "Title of AD", func() {
				MinLength(1)
				MaxLength(100)
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
					Maximum(100)
					Example(18)
				})
				Attribute("age_end", Int, "End age of target", func() {
					Minimum(1)
					Maximum(100)
					Example(60)
				})
				Attribute("gender", Genders, "Gender of target", func() {
					Example([]string{"M", "F"})
				})
				Attribute("country", Countries, "Country of target", func() {
					Example([]string{"TW", "JP"})
				})
				Attribute("platform", Platforms, "Platform of target", func() {
					Example([]string{"ios", "android", "web"})
				})
			})

			Required("title", "start_at", "end_at")
		})
		Result(func() {
			Field(1, "id", String, "ID of the AD", func() {
				Example("34")
			})

			Required("id")
		})

		HTTP(func() {
			POST("ad")
			Response(StatusCreated)
		})
	})

	Method("list_ads", func() {
		Meta("openapi:summary", "List all ADs by filter")
		Description("List all ADs by filter")

		Payload(AdQuery)
		Result(ArrayOf(Ad))
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
