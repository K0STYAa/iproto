package main

import (
	"bytes"
	"log"
	"net/rpc"
	"strings"

	"github.com/K0STYAa/iproto/pkg/iproto"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	AdmStorageSwitchReadOnly    = 0x00010001
	AdmStorageSwitchReadWrite   = 0x00010002
	AdmStorageSwitchMaintenance = 0x00010003
	StorageReplace              = 0x00020001
	StorageRead                 = 0x00020002
	BigID                       = 1000
	BigStringBytes              = 300
)

// Arrange.
type testTable struct {
	foo  uint32
	req  interface{}
	resp interface{}
}

var myTestTable = []testTable{ //nolint: gochecknoglobals
	{ // Read Empty Value
		foo:  StorageRead,
		req:  &iproto.ReqReadArgs{ID: 1},
		resp: &iproto.RespReadArgs{S: ""},
	},
	{ // Write Value
		foo:  StorageReplace,
		req:  &iproto.ReqReplaceArgs{ID: 1, S: "Kostya"},
		resp: &iproto.RespReplaceArgs{},
	},
	{ // Read New Value
		foo:  StorageRead,
		req:  &iproto.ReqReadArgs{ID: 1},
		resp: &iproto.RespReadArgs{S: "Kostya"},
	},
	{ // Replace Value
		foo:  StorageReplace,
		req:  &iproto.ReqReplaceArgs{ID: 1, S: "Max"},
		resp: &iproto.RespReplaceArgs{},
	},
	{ // Read New Value
		foo:  StorageRead,
		req:  &iproto.ReqReadArgs{ID: 1},
		resp: &iproto.RespReadArgs{S: "Max"},
	},
	{ // Change On State Read-Only
		foo:  AdmStorageSwitchReadOnly,
		req:  nil,
		resp: nil,
	},
	{ // Try To Write In State Read-Only
		foo:  StorageReplace,
		req:  &iproto.ReqReplaceArgs{ID: 1, S: "Serg"},
		resp: "can't replace at readOnly mode", // ERROR
	},
	{ // Read Same Value
		foo:  StorageRead,
		req:  &iproto.ReqReadArgs{ID: 1},
		resp: &iproto.RespReadArgs{S: "Max"},
	},
	{ // Change On State Maintenance
		foo:  AdmStorageSwitchMaintenance,
		req:  nil,
		resp: nil,
	},
	{ // Try To Write In State Maintenance
		foo:  StorageReplace,
		req:  &iproto.ReqReplaceArgs{ID: 1, S: "Serg"},
		resp: "can't replace at maintenance mode", // ERROR
	},
	{ // Try To Read In State Maintenance
		foo:  StorageRead,
		req:  &iproto.ReqReadArgs{ID: 1},
		resp: "can't read at maintenance mode", // ERROR
	},
	{ // Change On State Read-Write
		foo:  AdmStorageSwitchReadWrite,
		req:  nil,
		resp: nil,
	},
	{ // Write Big Value
		foo:  StorageReplace,
		req:  &iproto.ReqReplaceArgs{ID: 1, S: strings.Repeat("A", BigStringBytes)},
		resp: "incoming string cannot take up more than 256 bytes", // ERROR
	},
	{ // Read With Invalid ID
		foo:  StorageRead,
		req:  &iproto.ReqReadArgs{ID: -1},
		resp: "invalid ID. Valid value in range[0; 999]", // ERROR
	},
	{ // Read With Invalid ID
		foo:  StorageRead,
		req:  &iproto.ReqReadArgs{ID: BigID},
		resp: "invalid ID. Valid value in range[0; 999]", // ERROR
	},
	{ // Write With Invalid ID
		foo:  StorageReplace,
		req:  &iproto.ReqReplaceArgs{ID: -1, S: "a"},
		resp: "invalid ID. Valid value in range[0; 999]", // ERROR
	},
	{ // Write With Invalid ID
		foo:  StorageReplace,
		req:  &iproto.ReqReplaceArgs{ID: BigID, S: "a"},
		resp: "invalid ID. Valid value in range[0; 999]", // ERROR
	},
	{ // Return Storage
		foo:  StorageReplace,
		req:  &iproto.ReqReplaceArgs{ID: 1, S: ""},
		resp: &iproto.RespReplaceArgs{},
	},
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	// Act
	for requestID, testCase := range myTestTable {
		var resp iproto.Response

		reqBody, _ := msgpack.Marshal(testCase.req)
		testCaseRespBytes, _ := msgpack.Marshal(testCase.resp)
		req := iproto.Request{
			Header: iproto.Header{
				FuncID:     testCase.foo,
				BodyLength: uint32(len(reqBody)),
				RequestID:  uint32(requestID),
			},
			Body: reqBody,
		}

		err := client.Call("MyService.MainHandler", req, &resp)
		if err != nil {
			log.Fatal("Call error: ", err)
		}

		var bodyResp interface{}

		err = msgpack.Unmarshal(resp.Body, &bodyResp)
		if err != nil {
			log.Println(err.Error())
		}

		log.Printf("Calling %x(%v), result %v\n", testCase.foo, testCase.req, bodyResp)

		// Assert
		if bodyResp != nil && !bytes.Equal(resp.Body, testCaseRespBytes) ||
			bodyResp == nil && testCase.resp != bodyResp {
			log.Printf("Incorrect result. Expected %v, got %v",
				testCase.resp, bodyResp)
		}
	}
}
