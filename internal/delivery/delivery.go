package delivery

import (
	"github.com/sirupsen/logrus"

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

	var body_resp models.RespReadArgs
	if err == nil {
		err2 := msgpack.Unmarshal(resp, &body_resp)
		if err2 != nil {
			return_code = 1
			{ // LOG ERROR
				logrus.Warn("[ERROR]: ", err2.Error())
			}
			return_body, _ = msgpack.Marshal(err2.Error())
		} else {
			if body_resp.S != "" { // LOG RESPONSE
				logrus.Info("[RESPONSE]: ", body_resp.S)
			} else {
				logrus.Debug("[RESPONSE]: ", body_resp.S)
			}
			return_body = resp
		}
	} else {
		return_code = 1
		{ // LOG ERROR
			logrus.Warn("[ERROR]: ", err.Error())
		}
		return_body, _ = msgpack.Marshal(err.Error())
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