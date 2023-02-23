package delivery

import (
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
	handlerMap := map[uint32] func(*Delivery, []byte) ([]byte, error) {
        0x00010001: ADM_STORAGE_SWITCH_READONLY,
        0x00010002: ADM_STORAGE_SWITCH_READWRITE,
        0x00010003: ADM_STORAGE_SWITCH_MAINTENANCE,
        0x00020001: STORAGE_REPLACE,
        0x00020002: STORAGE_READ,
    }
	delivery_func := handlerMap[req.Header.FuncID]
	resp, err := delivery_func(d, req.Body)



	var return_code uint32
	var return_body []byte
	if err != nil {
		return_code = 1
		return_body, _ = msgpack.Marshal(err.Error())
	} else {
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
	// if req.Header.FuncID == 0x00010001 || req.Header.FuncID == 0x00010002 || req.Header.FuncID == 0x00010003 {
	// 	var tmp interface{}
	// 	msgpack.Unmarshal(result.Body, &tmp)
	// 	fmt.Println(req.Header.FuncID, tmp)
	// }
	return result
}