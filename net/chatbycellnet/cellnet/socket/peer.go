package socket

import "my_go/net/chatbycellnet/cellnet"

//在连接器 接受器、会话中可以通过socketPeer 结构派发消息

// 通讯端共享的数据
type socketPeer struct {
	// 事件处理函数回调
	eventFunc cellnet.EventFunc

	// 事件队列（可选）
	queue cellnet.EventQueue

	// 单独保存的保存cellnet.Peer接口
	peerInterface cellnet.Peer

	name string

	address string
}

// 获取通讯端的名称
func (s *socketPeer) Name() string {
	return s.name
}

// 设置通讯端的名称
func (s *socketPeer) SetName(v string) {
	s.name = v
}

// 获取队列
func (s *socketPeer) Queue() cellnet.EventQueue {
	return s.queue
}

// 设置通讯端事件回调
func (s *socketPeer) SetEvent(f cellnet.EventFunc) {
	s.eventFunc = f
}

// socket包内部派发事件
func (s *socketPeer) fireEvent(ev interface{}) interface{} {

	if s.eventFunc == nil {
		return nil
	}

	return s.eventFunc(ev)
}
