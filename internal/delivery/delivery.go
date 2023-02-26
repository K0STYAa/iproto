package delivery

import (
	"github.com/K0STYAa/vk_iproto/internal/service"
	"github.com/K0STYAa/vk_iproto/pkg/models"
	"github.com/sirupsen/logrus"
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
	handlerFuncMap := map[uint32]func(*Delivery, []byte) ([]byte, error){
		0x00010001: ADM_STORAGE_SWITCH_READONLY,    //nolint: nosnakecase
		0x00010002: ADM_STORAGE_SWITCH_READWRITE,   //nolint: nosnakecase
		0x00010003: ADM_STORAGE_SWITCH_MAINTENANCE, //nolint: nosnakecase
		0x00020001: STORAGE_REPLACE,                //nolint: nosnakecase
		0x00020002: STORAGE_READ,                   //nolint: nosnakecase
	}
	deliveryFunc := handlerFuncMap[req.Header.FuncID]

	// Call handler for function
	resp, err := deliveryFunc(d, req.Body)

	var (
		returnCode uint32
		returnBody []byte
		bodyResp   models.RespReadArgs
	)

	if err != nil {
		returnCode = 1
		{ // LOG ERROR
			logrus.Warn("[ERROR]: ", err.Error())
		}

		returnBody, _ = msgpack.Marshal(err.Error())
	} else if err2 := msgpack.Unmarshal(resp, &bodyResp); err2 != nil {
		returnCode = 1
		{ // LOG ERROR
			logrus.Warn("[ERROR]: ", err2.Error())
		}
		returnBody, _ = msgpack.Marshal(err2.Error())
	} else {
		if bodyResp.S != "" { // LOG RESPONSE
			logrus.Info("[RESPONSE]: ", bodyResp.S)
		} else {
			logrus.Debug("[RESPONSE]: ", bodyResp.S)
		}
		returnBody = resp
	}

	result := models.Response{
		Header: models.Header{
			FuncID:     req.Header.FuncID,
			BodyLength: uint32(len(returnBody)),
			RequestID:  req.Header.RequestID,
		},
		ReturnCode: returnCode,
		Body:       returnBody,
	}

	return result
}
