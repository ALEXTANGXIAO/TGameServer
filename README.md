# TGameServer
[![](https://img.shields.io/badge/made%20by-Alex%20Tang-blue.svg?style=flat-square)](https://protocol.ai)
[![](https://img.shields.io/badge/project-go-yellow.svg?style=flat-square)](https://libp2p.io/)
[![](https://img.shields.io/badge/freenode-%23libp2p-yellow.svg?style=flat-square)](https://webchat.freenode.net/?channels=%23libp2p)
[![codecov](https://codecov.io/gh/libp2p/go-reuseport/branch/master/graph/badge.svg)](https://codecov.io/gh/libp2p/go-reuseport)
[![Travis CI](https://travis-ci.org/libp2p/go-reuseport.svg?branch=master)](https://travis-ci.org/libp2p/go-reuseport)
[![Discourse posts](https://img.shields.io/discourse/https/discuss.libp2p.io/posts.svg)](https://discuss.libp2p.io)

### TGame服务端
#
### 使用Go + TCP + UDP +My

### doc: https://github.com/ALEXTANGXIAO/TGameServer
### 客户端地址: https://github.com/ALEXTANGXIAO/TGameUnity

Tested on `darwin`, `linux`, and `windows`.

---
```
├── Gameproto       // 协议
├── app             // 核心逻辑处理层
├── config          // 配置
├── logs            // 日志配置
├── manager         // 管理器
└── main.go         // 入口.go
```
```
TGame服务端 version1.0.0

参考资料

Go 语言使用 net 包实现 Socket 网络编程 https://zhuanlan.zhihu.com/p/143346084

Go语言TCP/UDP编程 https://blog.csdn.net/guidao13/article/details/84037778

Go 语言自学教程入门到精通实战进阶提升（学习路线+思维导图+视频教程+面试题+学习工具+大厂实战手册）https://www.magedu.com/84998.html
```
```
1.go env -w GO111MODULE=auto
2.go env -w GOPROXY=https://goproxy.cn/  或者
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/

go get -u github.com/golang/protobuf/protoc-gen-go

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go get github.com/golang/protobuf/proto
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest


protobuf解决办法
https://blog.csdn.net/qq_36532540/article/details/107304949

使用protobuf gogo
https://github.com/gogo/protobuf.git
```