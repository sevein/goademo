// Code generated by goa v3.11.1, DO NOT EDIT.
//
// demo client
//
// Command:
// $ goa gen github.com/sevein/goademo/design -o api

package demo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "demo" service client.
type Client struct {
	ShowEndpoint       goa.Endpoint
	PlainList0Endpoint goa.Endpoint
	PlainList1Endpoint goa.Endpoint
	List0Endpoint      goa.Endpoint
	List1Endpoint      goa.Endpoint
	List2Endpoint      goa.Endpoint
	List3Endpoint      goa.Endpoint
}

// NewClient initializes a "demo" service client given the endpoints.
func NewClient(show, plainList0, plainList1, list0, list1, list2, list3 goa.Endpoint) *Client {
	return &Client{
		ShowEndpoint:       show,
		PlainList0Endpoint: plainList0,
		PlainList1Endpoint: plainList1,
		List0Endpoint:      list0,
		List1Endpoint:      list1,
		List2Endpoint:      list2,
		List3Endpoint:      list3,
	}
}

// Show calls the "show" endpoint of the "demo" service.
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *GoademoThing, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*GoademoThing), nil
}

// PlainList0 calls the "plain-list-0" endpoint of the "demo" service.
func (c *Client) PlainList0(ctx context.Context) (res GoademoThingCollection, err error) {
	var ires interface{}
	ires, err = c.PlainList0Endpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(GoademoThingCollection), nil
}

// PlainList1 calls the "plain-list-1" endpoint of the "demo" service.
func (c *Client) PlainList1(ctx context.Context) (res GoademoThingCollection, err error) {
	var ires interface{}
	ires, err = c.PlainList1Endpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(GoademoThingCollection), nil
}

// List0 calls the "list0" endpoint of the "demo" service.
func (c *Client) List0(ctx context.Context) (res *List0Result, err error) {
	var ires interface{}
	ires, err = c.List0Endpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*List0Result), nil
}

// List1 calls the "list1" endpoint of the "demo" service.
func (c *Client) List1(ctx context.Context) (res *List1Result, err error) {
	var ires interface{}
	ires, err = c.List1Endpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*List1Result), nil
}

// List2 calls the "list2" endpoint of the "demo" service.
func (c *Client) List2(ctx context.Context) (res *List2Result, err error) {
	var ires interface{}
	ires, err = c.List2Endpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*List2Result), nil
}

// List3 calls the "list3" endpoint of the "demo" service.
func (c *Client) List3(ctx context.Context) (res *List3Result, err error) {
	var ires interface{}
	ires, err = c.List3Endpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*List3Result), nil
}
