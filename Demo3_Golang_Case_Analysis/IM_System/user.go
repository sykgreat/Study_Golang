package main

import (
	"log"
	"net"
	"strings"
)

type User struct {
	Name string      `json:"name"`
	Addr string      `json:"addr"`
	C    chan string `json:"c"`
	conn net.Conn

	server *Server
}

// NewUser 创建一个用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// OnLine 用户上线业务
func (user *User) OnLine() {
	// 用户上线 将用户加入到online_map中
	user.server.mapLock.Lock()
	user.server.OnlineMap[user.Name] = user
	user.server.mapLock.Unlock()

	// 广播当前用户上线消息
	user.server.BroadCast(user, "已上线")
}

// OffLine 用户下线业务
func (user *User) OffLine() {
	// 用户下线 将用户从online_map中删除
	user.server.mapLock.Lock()
	delete(user.server.OnlineMap, user.Name)
	user.server.mapLock.Unlock()

	// 广播当前用户上线消息
	user.server.BroadCast(user, "下线")
}

// SendMsg 给当前User对应的客户端发送消息
func (user *User) SendMsg(msg string) {
	if _, err := user.conn.Write([]byte(msg)); err != nil {
		log.Fatal(err)
		return
	}
}

// DoMessage 用户处理消息的业务
func (user *User) DoMessage(msg string) {
	if msg == "who" { // 查询当前在线用户都有那些
		user.server.mapLock.Lock()
		for _, user := range user.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ": 在线...\n"
			user.SendMsg(onlineMsg)
		}
		user.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" { // 更改当前账号的用户名
		// 消息格式：rename|xxx
		newName := strings.Split(msg, "|")[1]

		// 判断name是否存在
		_, ok := user.server.OnlineMap[newName]
		if ok {
			user.server.BroadCast(user, "当前用户名已被使用")
		} else {
			user.server.mapLock.Lock()
			delete(user.server.OnlineMap, user.Name)
			user.server.OnlineMap[newName] = user
			user.server.mapLock.Unlock()

			user.Name = newName
			user.SendMsg("您已经更新用户名：" + user.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" { // 私聊
		// 消息格式：to|用户名|消息内容

		// 1.获取对方的用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			user.SendMsg("消息格式不正确 请使用”to|用户名|消息内容“格式。")
			return
		}
		// 2.根据对方用户名 得到对方的user对象
		remoteUser, ok := user.server.OnlineMap[remoteName]
		if !ok {
			user.SendMsg("该用户名不存在\n")
			return
		}
		// 3.获取消息内容 通过对方的user对象将消息发送过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			user.SendMsg("无消息内容 请重新发送\n")
			return
		}

		remoteUser.SendMsg(user.Name + "对您说：" + content)
	} else {
		user.server.BroadCast(user, msg)
	}
}

// ListenMessage 监听当前User channel的方法 一旦有消息 就直接发送给对应的客户端
func (user *User) ListenMessage() {
	for {
		msg := <-user.C

		if _, err := user.conn.Write([]byte(msg + "\n")); err != nil {
			log.Fatal(err)
			return
		}
	}
}
