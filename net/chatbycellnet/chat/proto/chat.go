package proto

import (
	"my_go/net/chatbycellnet/cellnet"
	"reflect"
)

// 请求发送聊天内容
type ChatREQ struct {
	Content string // 发送的内容
}

// 有人发送聊天
type ChatACK struct {
	Content string // 发送的内容
	ID      int64  // 发送者的ID
}

func init() {
	//选择使用什么传输方式
	cellnet.RegisterMessageMeta("json", "proto.ChatREQ", reflect.TypeOf((*ChatREQ)(nil)).Elem(), 1)
	cellnet.RegisterMessageMeta("json", "proto.ChatACK", reflect.TypeOf((*ChatACK)(nil)).Elem(), 2)
}
