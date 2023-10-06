package main

/*
搜尋
needcheck
再次確認
*/
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	//_ "github.com/go-sql-driver/mysql"

	//_ "mysql_master"
	_ "github.com/go-sql-driver/mysql"
	//需要執行指令
	//需要執行這些語法去初始化環境讓他能夠讀到
	//_ "github.com/go-sql-driver/mysql"
	//go env -w GO111MODULE=on
	//go mod init Gone
	//go get github.com/go-sql-driver/mysql
)

//"time"

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析參數，預設是不會解析的
	fmt.Println(r.Form) //這些資訊是輸出到伺服器端的列印資訊
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!20230927") //這個寫入到 w 的是輸出到客戶端的
}

func main() {
	fmt.Println("lets go")
	testsql()
	fmt.Println("lets test")
	http.HandleFunc("/", sayhelloName)       //設定存取的路由
	err := http.ListenAndServe(":9090", nil) //設定監聽的埠
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//testsql()20231005 testok
	fmt.Println("lets test")
}

// ---------------------從這裡開始---------------------------------------
// 將dt.sql合併進入
var sqltype string = "mysql"

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Failed to redirect stderr to file: %v", err)
		fmt.Println(err)
		panic(err)
	}
}

// 如果這個變數只有設定 沒有被使用 golang會跳錯
// butter here 從這裡繼續測試
func testsql() {
	// create_conn
	//root:@/gosql?charset=utf8
	db, err := sql.Open(sqltype, "root:@/gosql?charset=utf8") //needcheck 20231005
	//db, err := sql.Open(sqltype, "astaxie:astaxie@/test?charset=utf8")
	//db, err := sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
	checkErr(err)

	//插入資料
	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")

	checkErr(err)

	res, err := stmt.Exec("astaxie", "研發部門", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新資料 update_data(conn, sql)
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查詢資料 read_data(conn, sql)
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//刪除資料 delete_data(conn, sql)
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

//

// var host = "localhost"
// var user = "root"
// var pwd = ""
// var db = "nodeMCU"

// func create_conn( host string, user string, pwd string, db string){
// 	conn = mysqli_connect(host, user, pwd);

// 	if(!conn){
// 		// throw new Exception("資料庫連線錯誤");
// 	}

// 	mysqli_query(conn, "SET NAmsg utf8");

// 	if(!mysqli_select_db(conn, db)){
// 		// throw new Exception("開啟資料庫失敗");
// 	}
// 	return conn;
// }

// func update_data(conn, sql){
// 	result = mysqli_query(conn, sql);
// 	if(!result){
// 		// throw new Exception("更新資料有誤");
// 	}
// 	return result;
// }

// func create_data(conn, sql){//沒有用到 或是可以跟read_data 一起用
// 	result = mysqli_query(conn, sql);
// 	if(!result){
// 		// throw new Exception("創建資料有誤");
// 	}
// 	return result;
// }

// func read_data(conn, sql){
// 	result = mysqli_query(conn, sql);
// 	if(!result){
// 		// throw new Exception("讀取資料有誤");
// 	}
// 	return result;
// }

// func delete_data(conn, sql){
// 	result = mysqli_query(conn, sql);
// 	if(!result){
// 		// throw new Exception("刪除資料有誤");
// 	}
// 	return result;
// }
//-----------------到這邊----------------------
