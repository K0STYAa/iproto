package main

import (
	"bytes"
	"log"
	"net/rpc"

	"github.com/K0STYAa/vk_iproto/pkg/models"
	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	// Arrange
	testTable := []struct {
		foo uint32
		req interface{}
		resp interface{}
	} {
		{ 	// Read Empty Value
			foo: 0x00020002,
			req: &models.ReqReadArgs{Id: 1},
			resp: &models.RespReadArgs{S: ""},
		},
		{	// Write Value
			foo: 0x00020001,
			req: &models.ReqReplaceArgs{Id: 1, S: "Kostya"},
			resp: &models.RespReplaceArgs{},
		},
		{	// Read New Value
			foo: 0x00020002,
			req: &models.ReqReadArgs{Id: 1},
			resp: &models.RespReadArgs{S: "Kostya"},
		},
		{	// Replace Value
			foo: 0x00020001,
			req: &models.ReqReplaceArgs{Id: 1, S: "Max"},
			resp: &models.RespReplaceArgs{},
		},
		{	// Read New Value
			foo: 0x00020002,
			req: &models.ReqReadArgs{Id: 1},
			resp: &models.RespReadArgs{S: "Max"},
		},
		{	// Change On State Read-Only 
			foo: 0x00010001,
			req: nil,
			resp: nil,
		},
		{	// Try To Write In State Read-Only
			foo: 0x00020001,
			req: &models.ReqReplaceArgs{Id: 1, S: "Serg"},
			resp: "Can't Replace At ReadOnly Mode", // ERROR
		},
		{	// Read Same Value
			foo: 0x00020002,
			req: &models.ReqReadArgs{Id: 1},
			resp: &models.RespReadArgs{S: "Max"},
		},
		{	// Change On State Maintenance 
			foo: 0x00010003,
			req: nil,
			resp: nil,
		},
		{	// Try To Write In State Maintenance
			foo: 0x00020001,
			req: &models.ReqReplaceArgs{Id: 1, S: "Serg"},
			resp: "Can't Replace At Maintenance Mode", // ERROR
		},
		{	// Try To Read In State Maintenance
			foo: 0x00020002,
			req: &models.ReqReadArgs{Id: 1},
			resp: "Can't Read At Maintenance Mode", // ERROR
		},
		{	// Change On State Read-Write 
			foo: 0x00010002,
			req: nil,
			resp: nil,
		},
		{	// Return Storage
			foo: 0x00020001,
			req: &models.ReqReplaceArgs{Id: 1, S: ""},
			resp: &models.RespReplaceArgs{},
		},

	}

	var resp models.Response

	// Act
	for i, testCase := range testTable {

		req_body, _ := msgpack.Marshal(testCase.req)
		testCase_resp_bytes, _ := msgpack.Marshal(testCase.resp)
		req := models.Request {
			Header: models.Header{
				FuncID: testCase.foo,
				BodyLength: uint32(len(req_body)),
				RequestID: uint32(i),
			},
			Body: req_body,
		}

		err := client.Call("MyService.MainHandler", req, &resp)
		if err != nil {
			log.Fatal("Call error:", err)
		}

		var body_resp interface{}
		msgpack.Unmarshal(resp.Body, &body_resp)
		log.Printf("Calling %x(%v), result %v\n", testCase.foo, testCase.req, body_resp)

		// Assert
		if bytes.Compare(resp.Body, testCase_resp_bytes) != 0 {
			log.Printf("Incorrect result. Expected %v, got %s", 
			testCase.resp, body_resp)
		}
	}
}

