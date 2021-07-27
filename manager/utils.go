package manager

import (
	"TGameServer/config"
	"time"

	"github.com/wonderivan/logger"
)

//挂起进程，防止销毁
func Plot() {
	time.Sleep(9999999999999999)
}

func SetLogger() {
	logger.SetLogger("./logs/log.json")
	logger.Debug("Server port", config.TCPport)
}

func Testlog() {
	// 通过配置文件配置
	// logger.SetLogger("./logs/log.json")
	//https://github.com/wonderivan/logger
	logger.Trace("this is Trace") // 由于默认输出，只会在控制台输出Debug及其以上日志，所以该条不会输出
	logger.Debug("this is Debug")
	logger.Info("this is Info")
	logger.Warn("this is Warn")
	logger.Error("this is Error")
	logger.Crit("this is Critical")
	logger.Alert("this is Alert")
	logger.Emer("this is Emergency")
}
