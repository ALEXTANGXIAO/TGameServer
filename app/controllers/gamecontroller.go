package controllers

import (
	GameProto "TGameServer/Gameproto"
	server "TGameServer/app/tserver"
	"errors"
)

func InitGameController() {
	controller := server.InstanceController("Game", GameProto.RequestCode_Game)
	controller.Funcs = map[string]interface{}{}
	controller.Funcs, _ = controller.AddFunction("UpPos", UpPos)
	server.RegisterController(GameProto.RequestCode_Game, controller)
}

func UpPos(client *server.Client, mainpack *GameProto.MainPack, isUdp bool) (*GameProto.MainPack, error) {
	if client == nil {
		return nil, errors.New("client is nill")
	}
	if client.RoomInfo == nil {
		return nil, errors.New("client roomInfo is nill")
	}
	client.RoomInfo.BroadcastUDP(client, mainpack)
	client.UpPos(mainpack)
	return nil, nil
}
