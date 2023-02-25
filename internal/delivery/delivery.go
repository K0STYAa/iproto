package delivery

import (
	"log"

	"github.com/K0STYAa/vk_iproto/internal/service"
	"github.com/K0STYAa/vk_iproto/pkg/models"
	"github.com/vmihailenco/msgpack/v5"
)

type Delivery struct {
	service *service.Service
}

func NewDelivery(service service.Service) *Delivery {
	return &Delivery{service: &service}
}

func (d *Delivery) MainHandler(req models.Request) models.Response {
	{ // LOG REQUEST 
		var body_req interface{}
		msgpack.Unmarshal(req.Body, &body_req)
		log.Printf("[REQUEST]: %x(%v)", req.Header.FuncID, body_req)
	}

	// Find Handler function by FuncID
	handlerFuncMap := map[uint32] func(*Delivery, []byte) ([]byte, error) {
        0x00010001: ADM_STORAGE_SWITCH_READONLY,
        0x00010002: ADM_STORAGE_SWITCH_READWRITE,
        0x00010003: ADM_STORAGE_SWITCH_MAINTENANCE,
        0x00020001: STORAGE_REPLACE,
        0x00020002: STORAGE_READ,
    }
	delivery_func := handlerFuncMap[req.Header.FuncID]

	// Call handler for function
	resp, err := delivery_func(d, req.Body)

	var return_code uint32
	var return_body []byte
	if err != nil {
		return_code = 1
		{ // LOG ERROR
			log.Println("[ERROR]:", err.Error())
		}
		return_body, _ = msgpack.Marshal(err.Error())
	} else {
		{ // LOG RESPONSE 
			var body_resp interface{}
			msgpack.Unmarshal(resp, &body_resp)
			log.Println("[RESPONSE]:", body_resp)
		}
		return_body = resp
	}

	result := models.Response{
		Header: models.Header{
			FuncID: req.Header.FuncID,
			BodyLength: uint32(len(return_body)),
			RequestID: req.Header.RequestID,
		},
		ReturnCode: return_code,
		Body: return_body,
	}
	return result
}