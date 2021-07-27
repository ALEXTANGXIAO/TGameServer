package tserver

import (
	GameProto "TGameServer/Gameproto"
	"TGameServer/config"
	"net"

	"github.com/gogo/protobuf/proto"
	"github.com/wonderivan/logger"
)

func (client *Client) StartUDP() {
	serviceUdp := config.UDPport
	logger.Info("Client Starting Udp Server:", serviceUdp)
	addr, err := net.ResolveUDPAddr("udp", serviceUdp)
	checkErr(err)
	conn, err := net.ListenUDP("udp", addr)
	checkErr(err)
	if conn == nil {
		return
	}
	// defer conn.Close()
	// client.UDPConn = conn
	// for {
	// 	client.handleClientUdp(conn)
	// }
}

func HandleClientUdp(conn *net.UDPConn) {
	buf := make([]byte, 1024)
	if conn == nil {
		return
	}
	n, remoteAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		logger.Info("failed to read UDP msg because of ", err.Error())
		return
	}
	errBuffer := handBufferUdp(conn, remoteAddr, buf[0:n])
	if checkNetErr(errBuffer, conn) {
		return
	}
}

func handBufferUdp(conn *net.UDPConn, remoteAddr *net.UDPAddr, buf []byte) error {
	mainPack := &GameProto.MainPack{}
	protoData := buf
	err := proto.Unmarshal(protoData, mainPack)
	if err != nil {
		return err
	}
	client, err := GetClient(mainPack.User)
	if err != nil {
		return err
	}
	if client.UDPConn == nil {
		client.UDPConn = conn
	}
	if client.UDPAddr == nil {
		client.UDPAddr = remoteAddr
	}
	handleReqUdp(client, mainPack, true)
	return nil
}
