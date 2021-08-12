package tserver

// import (
// 	"TGameServer/config"
// 	"errors"
// 	"net"
// 	"os"
// 	"strings"

// 	"github.com/wonderivan/logger"
// )

// type Server struct {
// 	Port string
// }

// var ClientList = []*Client{}
// var RoomList = []*Room{}

// var UDPConn *net.UDPConn
// var UDPAddr *net.UDPAddr

// func StartServer(port string) Server {
// 	server := Server{Port: port}
// 	Start(server)
// 	return server
// }

// func Start(server Server) {
// 	servicePort := server.Port
// 	logger.Info("Starting Tcp Server", servicePort)

// 	CreateRooms()

// 	//Start Udp Server
// 	uDPAddr, err := net.ResolveUDPAddr("udp", config.UDPport)
// 	UDPAddr = uDPAddr
// 	checkErr(err)
// 	uDPConn, err := net.ListenUDP("udp", UDPAddr)
// 	UDPConn = uDPConn
// 	checkErr(err)

// 	go func() {
// 		for {
// 			HandleClientUdp(UDPConn)
// 		}
// 	}()

// 	serverpro := NewConn("0.0.0.0", 54809, 80)
// 	serverpro.Start()
// }

// func AddClient(client *Client) ([]*Client, error) {
// 	ClientList = append(ClientList, client)
// 	logger.Debug("Add client to Server =>", client.Addr, "  ClientCount =>", len(ClientList))

// 	if len(ClientList) > 3000 {
// 		os.Exit(1)
// 	}
// 	return ClientList, nil
// }

// func GetClient(clientName string) (*Client, error) {
// 	for i := 0; i < len(ClientList); i++ {
// 		if ClientList[i] == nil {
// 			continue
// 		}
// 		if ClientList[i].Username == "" {
// 			logger.Emer("Had client Username = nil 👉", clientName)
// 			ClientList[i].Username = clientName
// 			return ClientList[i], nil
// 			// continue
// 		}
// 		if strings.Compare(ClientList[i].Username, clientName) == 0 {
// 			return ClientList[i], nil
// 		}
// 	}

// 	logger.Emer("Had not client 👉", clientName)
// 	return nil, errors.New("Had not client")
// }

// func RemoveClient(client *Client) {
// 	ClientList = RemoveC(ClientList, client)
// 	logger.Debug("Rmv client from Server =>", client.Addr, "  ClientCount =>", len(ClientList))
// }

// func checkErr(err error) {
// 	if err != nil {
// 		logger.Emer(os.Stderr, "Fatal error: %s", err.Error())
// 		// os.Exit(1)
// 	}
// }

// func checkNetErr(err error, conn net.Conn) bool {
// 	if err != nil {
// 		logger.Emer(conn.RemoteAddr())
// 		logger.Emer(os.Stderr, "Fatal error: %s", err.Error())
// 		return true
// 	}
// 	return false
// }
