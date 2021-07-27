package controllers

import (
	GameProto "TGameServer/Gameproto"
	server "TGameServer/app/tserver"
	"errors"

	"github.com/wonderivan/logger"
)

func InitUserController() {
	controller := server.InstanceController("User", GameProto.RequestCode_User)
	controller.Funcs = map[string]interface{}{}
	controller.Funcs, _ = controller.AddFunction("Login", Login)
	controller.Funcs, _ = controller.AddFunction("Register", Register)
	server.RegisterController(GameProto.RequestCode_User, controller)
}

func Login(client *server.Client, mainpack *GameProto.MainPack, isUdp bool) (*GameProto.MainPack, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	if CheckLogin(mainpack) {
		mainpack.Returncode = GameProto.ReturnCode_Success
		client.Username = mainpack.LoginPack.Username

		logger.Info("Login Test ↓ UserName")
		for i := 0; i < len(server.ClientList); i++ {
			logger.Info(server.ClientList[i].Username)
		}
	} else {
		mainpack.Returncode = GameProto.ReturnCode_Fail
	}
	return mainpack, nil
}

func Register(client *server.Client, mainpack *GameProto.MainPack, isUdp bool) (*GameProto.MainPack, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	if CheckLogin(mainpack) {
		mainpack.Returncode = GameProto.ReturnCode_Success
	} else {
		mainpack.Returncode = GameProto.ReturnCode_Fail
	}
	return mainpack, nil
}

func CheckLogin(mainpack *GameProto.MainPack) bool {
	//todo check mysqlLogin
	return true
}
