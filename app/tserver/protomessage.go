package tserver

import (
	"errors"
	"net"

	GameProto "TGameServer/Gameproto"

	"github.com/gogo/protobuf/proto"
	"github.com/wonderivan/logger"
)

var bufferHead int = 4

var buf []byte = make([]byte, 1024)

var startIndex int = 0

var remsize int = len(buf) - startIndex

func ReadBuffer(length int) error {
	startIndex := startIndex + len(buf)

	if startIndex <= bufferHead {
		return errors.New("Start Index is Error")
	}

	return nil
}

func PackData() (buf []byte) {

	return nil
}

func PackDataUDP() (buf []byte) {

	return nil
}

///////////////////////////////////////////////////////////////////////////////////
const BufferHeadLength int = 4

var BufferHead []byte = []byte{0, 0, 0, 0}

func handBuffer(client *Client, buf []byte) error {
	mainPack := &GameProto.MainPack{}
	protoData := buf

	//protobuf解码
	err := proto.Unmarshal(protoData, mainPack)
	if err != nil {
		// panic(err)
		logger.Emer("marshal error: ", err.Error())
		return err
	}
	logger.Info("Receive from client:", client.Conn.RemoteAddr(), mainPack)

	err2 := handleReq(client, mainPack, false)

	if err2 != nil {
		return err2
	}
	return nil
}

func sendBuffer(client *Client, mainpack *GameProto.MainPack) (*GameProto.MainPack, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	if mainpack == nil {
		return nil, errors.New("mainpack is nil")
	}

	conn := client.Conn

	if conn == nil {
		return mainpack, errors.New("conn is nil")
	}

	data, err := proto.Marshal(mainpack)

	if err != nil {
		logger.Debug("marshal error: ", err.Error())
		return nil, err
	}

	bodylen := len(data)

	BufferHead[0] = byte(bodylen)

	buff := append(BufferHead, data...)

	logger.Debug("send mainpack: ", conn.RemoteAddr(), mainpack)

	logger.Debug("send buff: ", buff)

	_, err2 := conn.Write(buff)

	return mainpack, err2
}

func sendBufferUdp(client *Client, mainpack *GameProto.MainPack) (*GameProto.MainPack, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	if mainpack == nil {
		return nil, errors.New("mainpack is nil")
	}

	conn := client.UDPConn

	if conn == nil {
		return mainpack, errors.New("conn is nil")
	}

	data, err := proto.Marshal(mainpack)

	if err != nil {
		logger.Debug("marshal error: ", err.Error())
	}

	bodylen := len(data)

	BufferHead[0] = byte(bodylen)

	buff := append(BufferHead, data...)

	logger.Debug("send mainpack: ", conn.RemoteAddr(), mainpack)

	logger.Debug("send buff: ", buff)

	conn.WriteToUDP(buff, client.UDPAddr)

	return mainpack, nil
}

func sendBufferTest(conn net.Conn) (*GameProto.MainPack, error) {
	oldData := &GameProto.LoginPack{
		Username: "tx",
		Password: "tx",
	}
	mainpack := &GameProto.MainPack{
		Returncode:  GameProto.ReturnCode_Success,
		Requestcode: GameProto.RequestCode_User,
		Actioncode:  GameProto.ActionCode_Login,
		LoginPack:   oldData,
	}

	data, err := proto.Marshal(mainpack)

	if err != nil {
		logger.Debug("marshal error: ", err.Error())
	}

	bodylen := len(data)

	BufferHead[0] = byte(bodylen)

	buff := append(BufferHead, data...)

	logger.Debug("send mainpack: ", conn.RemoteAddr(), mainpack)

	logger.Debug("send buff: ", buff)

	_, err2 := conn.Write(buff)

	return mainpack, err2
}
