package ziface

import (
	"net"
)

//定义连接接口
type IConnection interface {
	//启动链接
	Start()
	//停止链接
	Stop()
	//获取当前连接
	GetTcpConnection() *net.TCPConn
	//得到连接id
	GetConnId() uint32
	//客户端连接地址
	GetRemoteAddr() net.Addr
	//发送消息
	SendMsg(msgId uint32, data []byte) error
	//写数据
	StartWriter()
	//获取服务
	GetTcpServer() IServer
}
