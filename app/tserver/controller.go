package tserver

import (
	"errors"
	"reflect"

	GameProto "TGameServer/Gameproto"

	"github.com/wonderivan/logger"
)

type Controller struct {
	ControllerName string
	RequestCode    GameProto.RequestCode
	Funcs          map[string]interface{}
	Action         func(client Client, mainpack *GameProto.MainPack, isUdp bool) (*GameProto.MainPack, error)
}

func (controller *Controller) AddFunction(funcName string, inter interface{}) (map[string]interface{}, error) {
	_, hadValue := controller.Funcs[funcName]
	if !hadValue {
		controller.Funcs[funcName] = inter
	} else {
		errorstr := "Had Func" + funcName
		return controller.Funcs, errors.New(errorstr)
	}
	return controller.Funcs, nil
}

func (controller *Controller) Call(name string, params ...interface{}) (result []reflect.Value, err error) {
	_, hadValue := controller.Funcs[name]
	if !hadValue {
		logger.Emer("Func called %d, is not registered", name)
		return nil, errors.New("name = nil")
	}

	f := reflect.ValueOf(controller.Funcs[name])
	if f.Type() == nil {
		return nil, errors.New("f = nil")
	}

	if len(params) != f.Type().NumIn() {
		return nil, errors.New("The number of params is not adapted.")
	}

	in := make([]reflect.Value, len(params))

	for key, param := range params {
		in[key] = reflect.ValueOf(param)
	}

	result = f.Call(in)
	return result, err
}

func InstanceController(name string, code GameProto.RequestCode) Controller {
	return Controller{ControllerName: name, RequestCode: code}
}

func RegisterController(code GameProto.RequestCode, ctr Controller) error {
	if controllerMap == nil {
		return errors.New("controllerMap is nil")
	}
	_, hadRegister := controllerMap[code]
	if !hadRegister {
		controllerMap[code] = ctr
	} else {
		errorstr := "Had Register" + code.String()
		return errors.New(errorstr)
	}
	logger.Debug("RegisterController", ctr)
	return nil
}

var controllerMap map[GameProto.RequestCode]Controller = map[GameProto.RequestCode]Controller{}

func handleReq(client *Client, mainpack *GameProto.MainPack, isUdp bool) error {
	conn := client.Conn
	if conn == nil {
		return errors.New("conn is nil")
	}
	if mainpack == nil {
		return errors.New("mainpack is nil")
	}

	requestType := mainpack.Requestcode

	ctr, ok := controllerMap[requestType]
	if ok {
		if isUdp {

		} else {
			valuePack, err := ctr.Call(mainpack.Actioncode.String(), client, mainpack, isUdp)
			if err != nil {
				return err
			}
			var sendpack *GameProto.MainPack = valuePack[0].Interface().(*GameProto.MainPack)

			// err = valuePack[1].Interface().(error)
			if sendpack == nil {
				return nil
			}
			if err != nil {
				return err
			}
			logger.Debug(sendpack)

			sendBuffer(client, sendpack)
		}
	} else {
		logger.Error("Unknown request", ctr, requestType)
	}

	return nil
}

func handleReqUdp(client *Client, mainpack *GameProto.MainPack, isUdp bool) error {
	if mainpack == nil {
		return errors.New("mainpack is nil")
	}

	requestType := mainpack.Requestcode

	ctr, ok := controllerMap[requestType]
	if ok {
		if isUdp {
			valuePack, err := ctr.Call(mainpack.Actioncode.String(), client, mainpack, isUdp)
			if err != nil {
				return err
			}
			var sendpack *GameProto.MainPack = valuePack[0].Interface().(*GameProto.MainPack)

			if err != nil {
				return err
			}
			if sendpack == nil {
				return errors.New("sendpack is nil")
			}
			logger.Debug(sendpack)

			sendBufferUdp(client, sendpack)

		} else {
			valuePack, err := ctr.Call(mainpack.Actioncode.String(), client, mainpack, isUdp)
			if err != nil {
				return err
			}
			var sendpack *GameProto.MainPack = valuePack[0].Interface().(*GameProto.MainPack)

			// err = valuePack[1].Interface().(error)

			if err != nil {
				// return err
			}
			logger.Debug(sendpack)

			sendBuffer(client, sendpack)
		}
	} else {
		logger.Error("Unknown request", ctr, requestType)
	}

	return nil
}
