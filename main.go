package main

import (
	"TGameServer/app/controllers"
	server "TGameServer/app/tserver"
	config "TGameServer/config"
	manager "TGameServer/manager"
)

func main() {
	manager.Testlog()
	manager.SetLogger()
	InitControllers()
	go server.StartServer(config.TCPport)
	manager.Plot()
}

// All Controllers Will Init On Here
func InitControllers() {
	controllers.InitUserController()
	controllers.InitRoomController()
	controllers.InitGameController()
}
