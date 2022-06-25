package main

import "net"

type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn

	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	user := &User{
		Name: conn.RemoteAddr().String(),
		Addr: conn.RemoteAddr().String(),
		C:    make(chan string),
		Conn: conn,

		server: server,
	}
	// 启动监听当前user channel 的message goroutine
	go user.ListenMessage()
	return user
}

func (this *User) Online() {
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.DoMessage("已上线")
}

func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	this.DoMessage("已下线")
}

func (this *User) DoMessage(msg string) {
	this.server.BroadCast(this, msg)
}

// 监听当前channel的goroutine，当有消息时打印出来
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.Conn.Write([]byte(msg + "\n"))
	}

}
