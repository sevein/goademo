// Code generated by goa v3.11.1, DO NOT EDIT.
//
// demo HTTP client types
//
// Command:
// $ goa gen github.com/sevein/goademo/design -o api

package client

import (
	demo "github.com/sevein/goademo/api/gen/demo"
	demoviews "github.com/sevein/goademo/api/gen/demo/views"
	goa "goa.design/goa/v3/pkg"
)

// ShowResponseBody is the type of the "demo" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	ID   *int    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Age  *int    `form:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
}

// PlainList0ResponseBody is the type of the "demo" service "plain-list-0"
// endpoint HTTP response body.
type PlainList0ResponseBody []*GoademoThingResponse

// PlainList1ResponseBody is the type of the "demo" service "plain-list-1"
// endpoint HTTP response body.
type PlainList1ResponseBody []*GoademoThingResponse

// List0ResponseBody is the type of the "demo" service "list0" endpoint HTTP
// response body.
type List0ResponseBody struct {
	Items GoademoThingCollectionResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
}

// List1ResponseBody is the type of the "demo" service "list1" endpoint HTTP
// response body.
type List1ResponseBody struct {
	Items GoademoThingCollectionResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
}

// List2ResponseBody is the type of the "demo" service "list2" endpoint HTTP
// response body.
type List2ResponseBody struct {
	Items GoademoThingCollectionResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
}

// List3ResponseBody is the type of the "demo" service "list3" endpoint HTTP
// response body.
type List3ResponseBody struct {
	Items GoademoThingCollectionResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
}

// GoademoThingResponse is used to define fields on response body types.
type GoademoThingResponse struct {
	ID   *int    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Age  *int    `form:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
}

// GoademoThingCollectionResponseBody is used to define fields on response body
// types.
type GoademoThingCollectionResponseBody []*GoademoThingResponseBody

// GoademoThingResponseBody is used to define fields on response body types.
type GoademoThingResponseBody struct {
	ID   *int    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Age  *int    `form:"age,omitempty" json:"age,omitempty" xml:"age,omitempty"`
}

// NewShowGoademoThingOK builds a "demo" service "show" endpoint result from a
// HTTP "OK" response.
func NewShowGoademoThingOK(body *ShowResponseBody) *demoviews.GoademoThingView {
	v := &demoviews.GoademoThingView{
		ID:   body.ID,
		Name: body.Name,
		Age:  body.Age,
	}

	return v
}

// NewPlainList0GoademoThingCollectionOK builds a "demo" service "plain-list-0"
// endpoint result from a HTTP "OK" response.
func NewPlainList0GoademoThingCollectionOK(body PlainList0ResponseBody) demoviews.GoademoThingCollectionView {
	v := make([]*demoviews.GoademoThingView, len(body))
	for i, val := range body {
		v[i] = unmarshalGoademoThingResponseToDemoviewsGoademoThingView(val)
	}

	return v
}

// NewPlainList1GoademoThingCollectionOK builds a "demo" service "plain-list-1"
// endpoint result from a HTTP "OK" response.
func NewPlainList1GoademoThingCollectionOK(body PlainList1ResponseBody) demoviews.GoademoThingCollectionView {
	v := make([]*demoviews.GoademoThingView, len(body))
	for i, val := range body {
		v[i] = unmarshalGoademoThingResponseToDemoviewsGoademoThingView(val)
	}

	return v
}

// NewList0ResultOK builds a "demo" service "list0" endpoint result from a HTTP
// "OK" response.
func NewList0ResultOK(body *List0ResponseBody) *demo.List0Result {
	v := &demo.List0Result{}
	v.Items = make([]*demo.GoademoThing, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalGoademoThingResponseBodyToDemoGoademoThing(val)
	}

	return v
}

// NewList1ResultOK builds a "demo" service "list1" endpoint result from a HTTP
// "OK" response.
func NewList1ResultOK(body *List1ResponseBody) *demo.List1Result {
	v := &demo.List1Result{}
	v.Items = make([]*demo.GoademoThing, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalGoademoThingResponseBodyToDemoGoademoThing(val)
	}

	return v
}

// NewList2ResultOK builds a "demo" service "list2" endpoint result from a HTTP
// "OK" response.
func NewList2ResultOK(body *List2ResponseBody) *demo.List2Result {
	v := &demo.List2Result{}
	v.Items = make([]*demo.GoademoThing, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalGoademoThingResponseBodyToDemoGoademoThing(val)
	}

	return v
}

// NewList3ResultOK builds a "demo" service "list3" endpoint result from a HTTP
// "OK" response.
func NewList3ResultOK(body *List3ResponseBody) *demo.List3Result {
	v := &demo.List3Result{}
	v.Items = make([]*demo.GoademoThing, len(body.Items))
	for i, val := range body.Items {
		v.Items[i] = unmarshalGoademoThingResponseBodyToDemoGoademoThing(val)
	}

	return v
}

// ValidateList0ResponseBody runs the validations defined on List0ResponseBody
func ValidateList0ResponseBody(body *List0ResponseBody) (err error) {
	if body.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "body"))
	}
	if body.Items != nil {
		if err2 := ValidateGoademoThingCollectionResponseBody(body.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateList1ResponseBody runs the validations defined on List1ResponseBody
func ValidateList1ResponseBody(body *List1ResponseBody) (err error) {
	if body.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "body"))
	}
	if body.Items != nil {
		if err2 := ValidateGoademoThingCollectionResponseBody(body.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateList2ResponseBody runs the validations defined on List2ResponseBody
func ValidateList2ResponseBody(body *List2ResponseBody) (err error) {
	if body.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "body"))
	}
	if body.Items != nil {
		if err2 := ValidateGoademoThingCollectionResponseBody(body.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateList3ResponseBody runs the validations defined on List3ResponseBody
func ValidateList3ResponseBody(body *List3ResponseBody) (err error) {
	if body.Items == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("items", "body"))
	}
	if body.Items != nil {
		if err2 := ValidateGoademoThingCollectionResponseBody(body.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateGoademoThingResponse runs the validations defined on
// GoademoThingResponse
func ValidateGoademoThingResponse(body *GoademoThingResponse) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}

// ValidateGoademoThingCollectionResponseBody runs the validations defined on
// GoademoThingCollectionResponseBody
func ValidateGoademoThingCollectionResponseBody(body GoademoThingCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateGoademoThingResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateGoademoThingResponseBody runs the validations defined on
// GoademoThingResponseBody
func ValidateGoademoThingResponseBody(body *GoademoThingResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	return
}
