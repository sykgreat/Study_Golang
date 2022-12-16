package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"sync"
	"time"
)

type Server struct {
	Ip   string `json:"ip,omitempty"`
	Port int    `json:"port,omitempty"`

	// 在线用户的列表
	OnlineMap map[string]*User `json:"onLineMap,omitempty"`
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string `json:"message,omitempty"`
}

// NewServer 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

// ListenMessager 监听Message广播消息channel的goroutine 一旦有消息 就发送给全部的在线user
func (server *Server) ListenMessager() {
	for {
		msg := <-server.Message

		server.mapLock.Lock()
		for _, cli := range server.OnlineMap {
			cli.C <- msg
		}
		server.mapLock.Unlock()
	}
}

// BroadCast 广播消息方法
func (server *Server) BroadCast(user *User, msg string) {
	setMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	server.Message <- setMsg
}

func (server *Server) Handler(conn net.Conn) {
	// 。。。当前链路的业务

	user := NewUser(conn, server)

	user.OnLine()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.OffLine()
				return
			}
			if err != nil && err != io.EOF {
				log.Fatal(err)
				return
			}

			// 提取用户消息（去除'\n'）
			msg := string(buf[:n-1])

			// 用户针对msg进行消息处理
			user.DoMessage(msg)

			// 用户的任意消息 代表当前用户是活跃的
			isLive <- true
		}
	}()

	for {
		// 当前handler阻塞
		select {
		case <-isLive:
			// 当前用户是活跃的 重置定时器
			// 不做如何事情 只为了激活select 重置定时器
		case <-time.After(5 * time.Minute):
			// 已经超时 将当前user强制关闭
			user.SendMsg("你被踢了")

			// 销毁用户资源
			close(user.C)

			// 关闭连接
			if err := conn.Close(); err != nil {
				log.Fatal(err)
				return
			}

			// 退出当前的Handler
			runtime.Goexit()
		}
	}
}

// Start 启动服务器的接口
func (server *Server) Start() {
	// socket listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		log.Fatal(err)
		return
	}

	// close listener socket
	defer func(listen net.Listener) {
		if err := listen.Close(); err != nil {
			log.Fatal(err)
			return
		}
	}(listener)

	// 启动监听message的goroutine
	go server.ListenMessager()

	for {
		//  accept
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		// do handler
		go server.Handler(conn)
	}
}
