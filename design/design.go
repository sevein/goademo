package design

import . "goa.design/goa/v3/dsl"

var _ = API("goademo", func() {
	Title("goademo API")
	Server("goademo", func() {
		Services("demo")
	})
})

var _ = Service("demo", func() {
	HTTP(func() {
		Path("/demo")
	})

	Method("show", func() {
		Payload(func() {
			Attribute("id", UInt)
			Required("id")
		})
		Result(
			Thing, func() { View("extended") },
		)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	// Below is a list of methods where I've tried to show a collection of
	// Thing using its "default" view which should not include the "age" attr.

	// ↓ ✅ - age is not included as expected.

	Method("plain-list-0", func() {
		Result(Things, func() {
			View("default")
		})
		HTTP(func() {
			GET("/plain-list-0")
			Response(StatusOK)
		})
	})

	// ↓ 🚨 - age is included.

	Method("plain-list-1", func() {
		Result(ThingsDefaultView)
		HTTP(func() {
			GET("/plain-list-1")
			Response(StatusOK)
		})
	})

	// ↓ 🚨 - age is included.

	Method("list0", func() {
		Result(func() {
			Attribute("items", Things, func() { View("default") })
			Required("items")
		})
		HTTP(func() {
			GET("/list0")
			Response(StatusOK)
		})
	})

	// ↓ 🚨 - age is included.

	Method("list1", func() {
		Result(func() {
			Attribute("items", Things)
			Required("items")
		})
		HTTP(func() {
			GET("/list1")
			Response(StatusOK)
		})
	})

	// ↓ 🚨 - age is included.

	Method("list2", func() {
		Result(func() {
			Attribute("items", ThingsDefaultView, func() { View("default") })
			Required("items")
		})
		HTTP(func() {
			GET("/list2")
			Response(StatusOK)
		})
	})

	// ↓ 🚨 - age is included.

	Method("list3", func() {
		Result(func() {
			Attribute("items", ThingsExtendedView, func() { View("extended") })
			Required("items")
		})
		HTTP(func() {
			GET("/list3")
			Response(StatusOK)
		})
	})
})

var Things = CollectionOf(Thing)

// Using View("default") here doesn't seem to have an impact.
// Note: https://pkg.go.dev/goa.design/goa/v3@v3.11.1/dsl#CollectionOf shows an example with View("tiny") - see TinyMultiResults.
var ThingsDefaultView = CollectionOf(Thing, func() { View("default") })

// Using View("default") here doesn't seem to have an impact.
// Note: https://pkg.go.dev/goa.design/goa/v3@v3.11.1/dsl#CollectionOf shows an example with View("tiny") - see TinyMultiResults.
var ThingsExtendedView = CollectionOf(Thing, func() { View("extended") })

// A thing.
var Thing = ResultType("application/vnd.goademo.thing", func() {
	Attributes(func() {
		Attribute("id", Int)
		Attribute("name", String)
		Attribute("age", Int)
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
	View("extended", func() {
		Attribute("id")
		Attribute("name")
		Attribute("age")
	})
	Required("id")
})
