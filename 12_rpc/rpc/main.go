package main

/*
### 什么是rpc
1. PRC(Remote Procedure Call) 远程过程调用， 简单理解是一个节点请求另一个节点提供的服务
2. 对应rpc的是本地过程调用，函数调用是最基本的本地过程调用
3. 将本地调用变层远程过程调用会面临各种问题
	1. call的id映射
	2. 序列化和返序列化
	3. 网络传输

### RPC技术在架构设计上有四部分组成。
分别是：客户端、客户端存根(clientStub)、服务端、服务端存根(serviceStub)
	客户端(Client)：服务调用发起方，也称为服务消费者。 大大，
	客户端存根(Client Stub)：该程序运行在客户端所在的计算机机器上，主要用来存储要调用的服务器的地址，另外，该程序还负责将客户端请求远端服务器程序的数据信息打包成数据包，通过网络发送给服务端Stub程序；其次，还要接收服务端Stub程序发送的调用结果数据包，并解析返回给客户端。
	服务端(Server)：远端的计算机机器上运行的程序，其中有客户端要调用的方法。
	服务端存根(Server Stub)：接收客户Stub程序通过网络发送的请求消息数据包，并调用服务端中真正的程序功能方法，完成功能调用；其次，将服务端执行调用的结果进行数据处理打包发送给客户端Stub程序。

#### 动态代理
利用动态代理生成对应的客户端，服务端的存根

go语言的rpc序列化与反序列化的协议是（gob）

*/

func main() {

}