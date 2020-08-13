package design

import . "goa.design/goa/v3/dsl"

var _ = API("goademo", func() {
	Title("goademo API")
	Server("goademo", func() {
		Services("demo")
	})
})

func PaginatedCollectionOf(v interface{}, adsl ...func()) interface{} {
	return func() {
		Attribute("items", CollectionOf(v, adsl...))
		Attribute("next_cursor", String)
		Required("items")
	}
}

var _ = Service("demo", func() {
	HTTP(func() {
		Path("/demo")
	})
	Method("show", func() {
		Payload(func() {
			Attribute("id", UInt)
			Required("id")
		})
		Result(StoredDemo)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})
	Method("list", func() {
		Payload(func() {
			Attribute("name", String)
		})
		Result(PaginatedCollectionOf(StoredDemo))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
			Params(func() {
				Param("name")
			})
		})
	})
})

var Demo = Type("Demo", func() {
	Attribute("name", String, "Name of the demo")
	Required("name")
})

var StoredDemo = ResultType("application/vnd.goademo.stored-demo", func() {
	Reference(Demo)
	Attributes(func() {
		Attribute("id", UInt)
		Attribute("name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
	})
	Required("id", "name")
})
