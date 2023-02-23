package main

import (
	"log"

	"github.com/K0STYAa/vk_iproto/pkg/models"
)

func main() {
	clientCodec := models.GetClientCodec()

	// Arrange
	testTable := []struct {
		foo string
		req interface{}
		resp interface{}
	} {
		{ 	// Read Empty Value
			foo: "Service.STORAGE_READ",
			req: &models.ReqReadArgs{Id: 1},
			resp: &models.RespReadArgs{S: ""},
		},
		{	// Write Value
			foo: "Service.STORAGE_REPLACE",
			req: &models.ReqReplaceArgs{Id: 1, S: "Kostya"},
			resp: &models.RespReplaceArgs{},
		},
		{	// Read New Value
			foo: "Service.STORAGE_READ",
			req: &models.ReqReadArgs{Id: 1},
			resp: &models.RespReadArgs{S: "Kostya"},
		},
		{	// Replace Value
			foo: "Service.STORAGE_REPLACE",
			req: &models.ReqReplaceArgs{Id: 1, S: "Max"},
			resp: &models.RespReplaceArgs{},
		},
		{	// Read New Value
			foo: "Service.STORAGE_READ",
			req: &models.ReqReadArgs{Id: 1},
			resp: &models.RespReadArgs{S: "Max"},
		},
		{	// Change On State Read-Only 
			foo: "Service.ADM_STORAGE_SWITCH_READONLY",
			req: nil,
			resp: nil,
		},
		{	// Try To Write In State Read-Only
			foo: "Service.STORAGE_REPLACE",
			req: &models.ReqReplaceArgs{Id: 1, S: "Serg"},
			resp: nil, // ERROR
		},
		{	// Read Same Value
			foo: "Service.STORAGE_READ",
			req: &models.ReqReadArgs{Id: 1},
			resp: &models.RespReadArgs{S: "Max"},
		},
		{	// Change On State Maintenance 
			foo: "Service.ADM_STORAGE_SWITCH_MAINTENANCE",
			req: nil,
			resp: nil,
		},

	}

	var result interface{}

	// Act
	for _, testCase := range testTable {
		err := clientCodec.Call(testCase.foo, testCase.req, &result)
		if err != nil {
			log.Fatal("Call error:", err)
		}

		log.Printf("Calling %s(%s), result %s\n", 
		testCase.foo, testCase.req, result)

		// Assert
		if result != testCase.resp {
			log.Printf("Incorrect result. Expected %s, got %s", 
				testCase.resp, result)
		}
	}
}

