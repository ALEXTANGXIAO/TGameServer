package tserver

import (
	GameProto "TGameServer/Gameproto"

	"github.com/wonderivan/logger"
)

type Room struct {
	ClientList []*Client
	RoomPack   *GameProto.RoomPack
}

func InstanceRoom(roomPack *GameProto.RoomPack) Room {
	room := Room{RoomPack: roomPack}
	return room
}

func CreateRooms() {
	for i := 0; i < 1; i++ {
		mainPack := &GameProto.MainPack{}
		mainPack.Requestcode = GameProto.RequestCode_Room
		roompack := &GameProto.RoomPack{}
		roompack.Roomname = "1"
		roompack.Maxnum = 999
		room := InstanceRoom(roompack)
		RoomList = append(RoomList, &room)
		// logger.Info(&RoomList)
	}
}

func (room *Room) Join(client *Client) {
	client.RoomInfo = room
	room.ClientList = append(room.ClientList, client)
	room.starting(client)
}

func (room *Room) starting(client *Client) {
	mainPack := &GameProto.MainPack{}
	mainPack.Requestcode = GameProto.RequestCode_Room
	mainPack.Actioncode = GameProto.ActionCode_Starting
	for i := 0; i < len(room.ClientList); i++ {
		playerpack := &GameProto.PlayerPack{}
		playerpack.Playername = room.ClientList[i].Username
		mainPack.Playerpack = append(mainPack.Playerpack, playerpack)
	}
	room.Broadcast(mainPack)
}

func (room *Room) Starting() {
	mainPack := &GameProto.MainPack{}
	mainPack.Requestcode = GameProto.RequestCode_Room
	mainPack.Actioncode = GameProto.ActionCode_Starting

	for i := 0; i < len(room.ClientList); i++ {
		playerpack := &GameProto.PlayerPack{}
		playerpack.Playername = room.ClientList[i].Username
		mainPack.Playerpack = append(mainPack.Playerpack, playerpack)
	}
	room.Broadcast(mainPack)
}

func (room *Room) Broadcast(mainPack *GameProto.MainPack) {
	for i := 0; i < len(room.ClientList); i++ {
		room.ClientList[i].SendTCP(mainPack)
	}
}

func (room *Room) BroadcastTCP(client *Client, mainPack *GameProto.MainPack) {
	if room == nil {
		logger.Emer("Broadcast failed room is nil!!!")
		return
	}
	for i := 0; i < len(room.ClientList); i++ {
		if room.ClientList[i] == client {
			continue
		}
		if room.ClientList[i] == nil {
			continue
		}
		room.ClientList[i].SendTCP(mainPack)
	}
}

func (room *Room) BroadcastUDP(client *Client, mainPack *GameProto.MainPack) {
	if room == nil {
		logger.Emer("Broadcast failed room is nil!!!")
		return
	}
	for i := 0; i < len(room.ClientList); i++ {
		if room.ClientList[i] == client {
			continue
		}
		if room.ClientList[i] == nil {
			continue
		}
		room.ClientList[i].SendUDP(mainPack)
	}
}
