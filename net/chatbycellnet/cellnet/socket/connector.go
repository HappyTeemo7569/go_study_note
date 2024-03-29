package socket

import (
	"my_go/net/chatbycellnet/cellnet"
	"my_go/net/chatbycellnet/cellnet/internal"
	"net"
)

//连接器 主要用来建立TCP连接
type socketConnector struct {
	socketPeer
	internal.SessionManager
	ses cellnet.Session
}

func (c *socketConnector) Start(address string) cellnet.Peer {
	c.address = address
	go c.connect(address)
	return c
}

func (c *socketConnector) Session() cellnet.Session {
	return c.ses
}

func (c *socketConnector) Stop() {

}

// 连接器，传入连接地址和发送封包次数
func (c *socketConnector) connect(address string) {

	// 尝试用Socket连接地址
	conn, err := net.Dial("tcp", address)
	ses := newSession(conn, &c.socketPeer)
	c.ses = ses

	// 发生错误时退出
	if err != nil {
		c.fireEvent(ConnectErrorEvent{ses, err})
		return
	}

	log.Infof("#connected(%s) %s", c.Name(), c.address)
	ses.start()

}

func NewConnector(f cellnet.EventFunc, q cellnet.EventQueue) cellnet.Peer {

	p := &socketConnector{
		SessionManager: internal.NewSessionManager(),
	}
	p.queue = q
	p.eventFunc = f
	p.peerInterface = p

	return p
}
