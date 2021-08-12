package tserver

import (
	GameProto "TGameServer/Gameproto"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/wonderivan/logger"
)

func ListenCMD() {
	reader := bufio.NewReader(os.Stdin)

	for {
		//从终端读取数据（会以文件的形式）
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readstring err=", err)
		}
		if line == "" {
			continue
		}
		//如多用户输入exit就退出
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("cmd退出")
			return
		}
		//将line发给服务器
		// n, err := conn.Write([]byte(line + "\n"))
		// if err != nil {
		// 	fmt.Println("conn.Write err=", err)
		// }

		if line == "send" {
			logger.Crit("CMD send")
			continue
		}

		BroadcastStr(line)
		logger.Crit("CMD 发送了字节", line)
	}
}

func BroadcastStr(line string) {
	mainPack, _ := BuildProto(GameProto.RequestCode_Room, GameProto.ActionCode_Chat, GameProto.ReturnCode_Success)
	mainPack.User = "GM"
	mainPack.Str = line
	RoomList[0].Broadcast(mainPack)
}
