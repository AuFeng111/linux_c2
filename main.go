package main

import (
	"C2_Linx/db"
	"C2_Linx/encode"
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

func main()  {
	key := []byte{0x61,0x65,0x6f,0x6e,0x64,0x67,0x5f,0x67,0x6f,0x66}
	a:=encode.Encrypt(key)
	fmt.Println(a)
	db2, err := sql.Open("sqlite3", "D:\\go1.20.2.windows-amd64\\go\\src\\C2_Linx\\db\\userDB.db")
	if err!=nil{
		fmt.Println(err)
	}
	db.Insert(db2,"aufeng","127.0.0.1","12:00:123")
	db.Query(db2)
	//db.Main()
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)
	currentTimeString := currentTime.Format("2023-01-02 15:04:05")
	fmt.Println(currentTimeString)
	fmt.Println(reflect.TypeOf(currentTimeString))

	b :=db.Query2()
	fmt.Println(len(b))
	var name1 string = ""
	for i:=0;i<len(b);i++{
		name1 ="|" +b[i].Hostname+"  " + b[i].ConnIP+"  " +b[i].Time+"|"+"\n"+name1
	}
	fmt.Println(name1)
}
