package tserver

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/wonderivan/logger"
)

type Conn struct {
	IP             string
	Port           uint32
	TCPConn        *net.TCPConn
	TaskChan       chan *Task
	ExitChan       chan bool
	Closed         bool
	WorkerPool     []chan *Task
	WorkerPoolSize uint32
	PreWorkerQueue uint32
	ConnMap        map[uint32]*net.TCPConn
}

type Task struct {
	Uniid   uint32
	MsgChan []byte
}

func NewTask(uniid uint32, msg []byte) *Task {
	task := Task{
		Uniid:   uniid,
		MsgChan: msg,
	}

	return &task
}

func NewConn(IP string, Port uint32, WorkerPoolSize uint32) *Conn {
	conn := &Conn{
		IP:             IP,
		Port:           Port,
		TaskChan:       make(chan *Task),
		ExitChan:       make(chan bool),
		WorkerPool:     make([]chan *Task, WorkerPoolSize),
		WorkerPoolSize: WorkerPoolSize,
		PreWorkerQueue: 1024,
		ConnMap:        make(map[uint32]*net.TCPConn),
	}
	return conn
}

func (c *Conn) Start() {
	log.Printf("%s:%d start...\n", c.IP, c.Port)
	go func() {
		c.StartWorkerPool()
		addr, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf("%s:%d", c.IP, c.Port))
		if err != nil {
			log.Println("resolve tcp addr err ", err)
			return
		}
		listener, err := net.ListenTCP("tcp4", addr)
		if err != nil {
			log.Println("listen tcp err ", err)
			return
		}
		var connid uint32
		connid = 0
		for {
			conn, err := listener.AcceptTCP()

			connid++
			uniid := connid

			client := InstanceClient(conn, uniid)
			ClientList, _ = AddClient(client)

			logger.Debug("new conn:", conn.RemoteAddr())
			if err != nil {
				log.Println("accept tcp err ", err)
				continue
			}

			go c.StartRead(client)
			go c.StartWrite(client)
			c.ConnMap[uniid] = conn

			log.Println("uniid :", uniid, c.ConnMap[uniid])
		}
	}()
	select {}
}
func (c *Conn) StartRead(client *Client) {
	log.Println("read groutine is waiting")
	defer c.Stop(client.Uniid)
	defer log.Println("read groutine exit")
	buf := make([]byte, 1024)
	for {
		cnt, err := client.Conn.Read(buf)
		if err != nil {
			log.Println("startread read bytes error ", err)
			continue
		}

		logger.Info(client.Conn.RemoteAddr(), "start read from client ", buf[:buf[0]])
		if BufferHeadLength > cnt {
			continue
		}
		buf = buf[BufferHeadLength:cnt]

		if c.WorkerPoolSize > 0 {
			c.SendMsgToWorker(buf, client.Uniid)
		} else {
			go c.HandleMsg(buf, client.Uniid)
		}
	}
}
func (c *Conn) StartWrite(client *Client) {
	log.Println("write groutine is waiting")
	defer log.Println("write groutine exit")
	for {
		select {
		case data := <-c.TaskChan:
			if client == nil {
				log.Println("client is nil")
			} else {

				log.Println(c.ConnMap[data.Uniid].RemoteAddr())
				for i := 0; i < len(ClientList); i++ {
					if ClientList[i].Uniid == data.Uniid {
						handBuffer(ClientList[i], data.MsgChan)
						break
					}
				}
			}
		case <-c.ExitChan:
			return
		}
	}
}
func (c *Conn) HandleMsg(data []byte, uniid uint32) {
	task := NewTask(uniid, data)
	c.TaskChan <- task
}
func (c *Conn) SendMsgToWorker(data []byte, uniid uint32) {
	rand.Seed(time.Now().UnixNano())
	workerId := rand.Intn(int(c.WorkerPoolSize))
	task := NewTask(uniid, data)
	c.WorkerPool[workerId] <- task
}
func (c *Conn) StartWorkerPool() {
	for i := 0; i < int(c.WorkerPoolSize); i++ {
		c.WorkerPool[i] = make(chan *Task, c.PreWorkerQueue)
		go c.StartOneWorker(i, c.WorkerPool[i])
	}
}
func (c *Conn) StartOneWorker(workerId int, queue chan *Task) {
	log.Println("start one worker groutine is waiting:", workerId)
	for {
		select {
		case data := <-queue:
			c.HandleMsg(data.MsgChan, data.Uniid)
			log.Println("one worker groutine is finshed:", workerId)
		}
	}
}

func (c *Conn) Stop(uniid uint32) {
	c.ExitChan <- true
	c.ConnMap[uniid].Close()
	delete(c.ConnMap, uniid)
}
