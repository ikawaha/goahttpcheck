package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")
	Method("multiply", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})
		Result(Int)
		HTTP(func() {
			GET("/multiply/{a}/{b}")
			Response(StatusOK)
		})
	})

	Method("divide", func() {
		Error("zero_division", ErrorResult)
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})
		Result(Int)
		HTTP(func() {
			GET("/divide/{a}/{b}")
			Response(StatusOK)
			Response("zero_division", StatusBadRequest)
		})
	})

	Method("redirect", func() {
		HTTP(func() {
			GET("/redirect")
			Redirect("https://localhost/redirect-test", StatusTemporaryRedirect)
		})
	})
})
