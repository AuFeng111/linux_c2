- ### 自研的linux c2
1、传输加解密√<br> 
2、参数化入参√<br>
3、数据库接入存储 记录c2上线的日志、主机信息等等√<br>
4、log日志x

----------------

- ### c2研发脑图

![c2研发脑图](/png/1.png)

---------------------------------------


### 使用方法
服务端启动
server.exe 8899

![](/png/2.png)

客户端建立连接，首先对ip端口进行base32加密

![](/png/3.png)


执行二进制文件如下操作，就可以进行连接

![](/png/4.png)


然后使用中控端对c2进行控制，如下1、2、3步骤操作

![](/png/5.png)




---------------------------------------


- bug修复日志

23年/5月/20日 client-c2端，存在逻辑bug问题
client端需要进行逻辑判断，然后再给服务端发送对应的协议。。服务端的协议对不上，就有bug了
socket/client-modity/client-c2.go

23年/5月/21日 client-c2端，解决逻辑bug问题 新增检查存活探测

23年5月23日  ps、netstat等无法获取全部内容消息

23年5月28日 解决23日的bug
能完全显示ps和netstat等长命令回显


23年5月31日 接入sqlite3存储c2上线日志
主机名、上线时间、连接的ip
