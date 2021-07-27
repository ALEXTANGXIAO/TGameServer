package controllers

import (
	GameProto "TGameServer/Gameproto"
	server "TGameServer/app/tserver"
	"errors"

	"github.com/wonderivan/logger"
)

func InitRoomController() {
	controller := server.InstanceController("Room", GameProto.RequestCode_Room)
	controller.Funcs = map[string]interface{}{}
	controller.Funcs, _ = controller.AddFunction("JoinRoom", JoinRoom)
	controller.Funcs, _ = controller.AddFunction("Chat", Chat)
	server.RegisterController(GameProto.RequestCode_Room, controller)
}

func Chat(client *server.Client, mainpack *GameProto.MainPack, isUdp bool) (*GameProto.MainPack, error) {
	if client == nil {
		return nil, errors.New("client is nill")
	}
	if client.RoomInfo == nil {
		return nil, errors.New("client roomInfo is nill")
	}
	mainpack.User = client.Username
	if CheckLogin(mainpack) {
		mainpack.Returncode = GameProto.ReturnCode_Success
		client.RoomInfo.BroadcastTCP(client, mainpack)
	} else {
		mainpack.Returncode = GameProto.ReturnCode_Fail
	}
	return nil, nil
}

func CreateRoom(client *server.Client, mainpack *GameProto.MainPack, isUdp bool) (*GameProto.MainPack, error) {
	if CheckLogin(mainpack) {
		mainpack.Returncode = GameProto.ReturnCode_Success
	} else {
		mainpack.Returncode = GameProto.ReturnCode_Fail
	}
	return mainpack, nil
}

func JoinRoom(client *server.Client, mainpack *GameProto.MainPack, isUdp bool) (*GameProto.MainPack, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}

	if len(server.RoomList) == 0 {
		logger.Error("RoomList count is empty")
	} else {
		room := server.RoomList[0]
		client.RoomInfo = room
		room.ClientList = append(room.ClientList, client)
		room.Starting()
	}

	if CheckLogin(mainpack) {
		mainpack.Returncode = GameProto.ReturnCode_Success
	} else {
		mainpack.Returncode = GameProto.ReturnCode_Fail
	}
	return mainpack, nil
}
