package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func Main() {
	//db, err := sql.Open("sqlite3", "D:\\go1.20.2.windows-amd64\\go\\src\\C2_Linx\\db\\userDB.db")
	db, err := sql.Open("sqlite3", "D:\\go1.20.2.windows-amd64\\go\\src\\C2_Linx\\db\\userDB.db")

	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)
	Insert(db,"aufeng", "10.0.0.1", "2023-5-31")
	Query(db)
	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id值:  ",id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("affect值:   ",affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid的值:  ",uid)
		fmt.Println("username的值:   ",username)
		fmt.Println("department:   ",department)
		fmt.Println("created:    ",created)
	}

	//删除数据
	//stmt, err = db.Prepare("delete from userinfo where uid=?")
	//checkErr(err)
	//
	//res, err = stmt.Exec(id)
	//checkErr(err)
	//
	//affect, err = res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect)

	db.Close()

}

func To_Insert(hostname string, connip string, time string)  {
	//db, err := sql.Open("sqlite3", "/root/桌面/test/db/userDB.db")
	db, err := sql.Open("sqlite3", "D:\\go1.20.2.windows-amd64\\go\\src\\C2_Linx\\db\\userDB.db")
	checkErr(err)
	Insert(db,hostname, connip, time)
}

func Insert(db *sql.DB,hostname string, connip string, time string){
	stmt, err := db.Prepare("INSERT INTO C2_INFO(hostname, connip, time) values(?,?,?)")
	checkErr(err)

	_, err = stmt.Exec(hostname, connip, time)
	checkErr(err)
}

func Query(db *sql.DB)  {
	rows, err := db.Query("SELECT * FROM C2_INFO")
	checkErr(err)

	for rows.Next() {
		var uid int
		var hostname string
		var connip string
		var time string
		err = rows.Scan(&uid, &hostname, &connip, &time)
		checkErr(err)
		fmt.Println("uid的值:  ",uid)
		fmt.Println("hostname:   ",hostname)
		fmt.Println("connip:   ",connip)
		fmt.Println("time:    ",time)
	}
}
type C2Info struct {
	UID      int
	Hostname string
	ConnIP   string
	Time     string
}

func Query2(db *sql.DB) []C2Info {
	//db, err := sql.Open("sqlite3", "/root/桌面/test/db/userDB.db")
	//db, err := sql.Open("sqlite3", "D:\\go1.20.2.windows-amd64\\go\\src\\C2_Linx\\db\\userDB.db")
	rows, err := db.Query("SELECT * FROM C2_INFO")
	checkErr(err)

	var c2InfoArr []C2Info
	for rows.Next() {
		var uid int
		var hostname string
		var connip string
		var time string
		err = rows.Scan(&uid, &hostname, &connip, &time)
		checkErr(err)
		fmt.Println("uid的值:  ", uid)
		fmt.Println("hostname:   ", hostname)
		fmt.Println("connip:   ", connip)
		fmt.Println("time:    ", time)

		c2Info := C2Info{UID: uid, Hostname: hostname, ConnIP: connip, Time: time}
		c2InfoArr = append(c2InfoArr, c2Info)
	}

	return c2InfoArr
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}