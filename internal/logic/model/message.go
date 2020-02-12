package model

type Message struct {
	// Id       int32 // id
	SenderId int32 // 发送方

	ReceiverType int8  // 接收方类型，用户或群组
	ReceiveId    int32 // 接受方，用户或者群的id

	Type    int8   // 消息类型
	Content string // 消息内容
}
