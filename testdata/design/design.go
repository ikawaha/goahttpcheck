package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("calc", func() {
	Description("The calc service performs operations on numbers.")
	Method("add", func() {
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})
		Result(Int)
		HTTP(func() {
			GET("/add/{a}/{b}")
			Response(StatusOK)
		})
	})

	Method("div", func() {
		Error("zero_division", ErrorResult)
		Payload(func() {
			Field(1, "a", Int, "Left operand")
			Field(2, "b", Int, "Right operand")
			Required("a", "b")
		})
		Result(Int)
		HTTP(func() {
			GET("/div/{a}/{b}")
			Response(StatusOK)
			Response("zero_division", StatusBadRequest)
		})
	})
})
