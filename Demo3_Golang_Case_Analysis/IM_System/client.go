package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type Client struct {
	ServerIp   string `json:"ServerIp,omitempty"`
	ServerPort int    `json:"serverPort,omitempty"`
	Name       string `json:"name,omitempty"`
	conn       net.Conn
	flagStr    int // 当前client的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flagStr:    999,
	}

	// 连接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	client.conn = conn

	//返回对象
	return client
}

// DealResponse 处理server回应的消息 直接显示到标准输出即可
func (client *Client) DealResponse() {
	// 一旦client.conn有数据 就直接copy到stdout标准输出上 永久阻塞监听
	if _, err := io.Copy(os.Stdout, client.conn); err != nil {
		log.Fatal(err)
		return
	}
}

// 菜单
func (client *Client) menu() bool {
	var flagStr int

	log.Println("1.公聊模式")
	log.Println("2.私聊模式")
	log.Println("3.更新用户名")
	log.Println("0.退出")

	if _, err := fmt.Scanln(&flagStr); err != nil {
		log.Fatal(err)
		return false
	}

	if flagStr >= 0 && flagStr <= 3 {
		client.flagStr = flagStr
		return true
	} else {
		log.Println(">>>>>>>>>>>>>>>请输入合法范围内的数组...<<<<<<<<<<<<<<<")
		return false
	}
}

// SelectUsers 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	if _, err := client.conn.Write([]byte(sendMsg)); err != nil {
		log.Fatal(err)
		return
	}
}

// PrivateChat 私聊模式
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()
	log.Println(">>>>>>>>>>>>>>>请输入聊天对象[用户名], exit退出:...<<<<<<<<<<<<<<<")
	if _, err := fmt.Scanln(&remoteName); err != nil {
		log.Fatal(err)
		return
	}

	for remoteName != "exit" {
		fmt.Println(">>>>>>>>>>>>>>>请输入消息内容, exit退出:...<<<<<<<<<<<<<<<")
		if _, err := fmt.Scanln(&chatMsg); err != nil {
			log.Fatal(err)
			return
		}

		for chatMsg != "exit" {
			//消息不为空则发送
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				if _, err := client.conn.Write([]byte(sendMsg)); err != nil {
					log.Fatal(err)
					break
				}
			}

			chatMsg = ""
			log.Println(">>>>>>>>>>>>>>>请输入消息内容, exit退出:...<<<<<<<<<<<<<<<")
			if _, err := fmt.Scanln(&chatMsg); err != nil {
				log.Fatal(err)
				return
			}
		}

		client.SelectUsers()
		log.Println(">>>>>>>>>>>>>>>请输入聊天对象[用户名], exit退出:...<<<<<<<<<<<<<<<")
		if _, err := fmt.Scanln(&remoteName); err != nil {
			log.Fatal(err)
			return
		}
	}
}

// PublicChat 公聊模式
func (client *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string

	log.Println(">>>>>>>>>>>>>>>请输入聊天内容 exit退出...<<<<<<<<<<<<<<<")
	if _, err := fmt.Scanln(&chatMsg); err != nil {
		log.Fatal(err)
		return
	}

	for chatMsg != "exit" {
		//  发给服务器

		// 消息不为空 则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			if _, err := client.conn.Write([]byte(sendMsg)); err != nil {
				log.Fatal(err)
				break
			}
		}

		chatMsg = ""
		log.Println(">>>>>>>>>>>>>>>请输入聊天内容 exit退出...<<<<<<<<<<<<<<<")
		if _, err := fmt.Scanln(&chatMsg); err != nil {
			log.Fatal(err)
			return
		}
	}
}

// UpdateName 更新用户名
func (client *Client) UpdateName() bool {
	log.Println(">>>>>>>>>>>>>>>请输入用户名...<<<<<<<<<<<<<<<")
	if _, err := fmt.Scanln(&client.Name); err != nil {
		log.Fatal(err)
		return false
	}

	sendMsg := "rename|" + client.Name + "\n"
	if _, err := client.conn.Write([]byte(sendMsg)); err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func (client *Client) Run() {
	for client.flagStr != 0 {
		for client.menu() != true {

		}
		// 根据不同模式处理不同的业务
		switch client.flagStr {
		case 1:
			// 公聊模式
			client.PublicChat()
			break
		case 2:
			// 私聊模式
			client.PrivateChat()
			break
		case 3:
			// 更新用户名
			client.UpdateName()
			break
		}
	}

}

var serverIp string
var serverPort int

// 。/client -ip 127.0.0.1 -port 8080
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址（默认是127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口（默认是8080）")
}

func main() {
	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		log.Println(">>>>>>>>>>>>>>>连接服务器失败...<<<<<<<<<<<<<<<")
		return
	}

	// 单独开启一个goroutine去处理server的回执消息
	go client.DealResponse()

	log.Println(">>>>>>>>>>>>>>>连接服务器成功...<<<<<<<<<<<<<<<")

	// 启动客户端业务
	client.Run()
}
