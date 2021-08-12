package tserver

import (
	"net"

	GameProto "TGameServer/Gameproto"

	"github.com/gogo/protobuf/proto"
	"github.com/wonderivan/logger"
)

type Client struct {
	Addr     string
	Conn     net.Conn
	UDPConn  *net.UDPConn
	UDPAddr  *net.UDPAddr
	Username string
	RoomInfo *Room
	PosPack  *GameProto.PosPack
	Uniid    uint32
}

func InstanceClient(conn net.Conn, uniid uint32) *Client {

	rAddr := conn.RemoteAddr()

	client := Client{Addr: rAddr.String(), Conn: conn, Uniid: uniid}

	return &client
}

func (client *Client) SendTCP(mainpack *GameProto.MainPack) {
	if client == nil {
		return
	}
	if client.Conn == nil {
		return
	}
	data, err := proto.Marshal(mainpack)
	if err != nil {
		logger.Emer("marshal error: ", err.Error())
		return
	}
	bodylen := len(data)

	buff := make([]byte, 0)

	buff = append(buff, byte(bodylen), 0, 0, 0)

	for i := 0; i < bodylen; i++ {
		buff = append(buff, byte(data[i]))
	}
	logger.Debug("send mainpack: ", mainpack)
	logger.Debug("send buff: ", buff)

	_, err2 := client.Conn.Write(buff)
	//_, err2 := client.UDPConn.WriteToUDP(buff, client.UDPAddr)

	if err2 != nil {
		logger.Debug("marshal error: ", err2.Error())
	}
}

func (client *Client) SendUDP(mainpack *GameProto.MainPack) {
	if client == nil {
		return
	}
	if client.UDPConn == nil {
		return
	}
	data, err := proto.Marshal(mainpack)
	if err != nil {
		logger.Emer("marshal error: ", err.Error())
		return
	}
	_, err2 := client.UDPConn.WriteToUDP(data, client.UDPAddr)

	if err2 != nil {
		logger.Emer("UDPConn error: ", err2.Error())
	}
}

func (client *Client) UpPos(mainpack *GameProto.MainPack) {
	if client == nil {
		logger.Emer("Error client is nil")
		return
	}
	client.PosPack = mainpack.Playerpack[0].PosPack
}
