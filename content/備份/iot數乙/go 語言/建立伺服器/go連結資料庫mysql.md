[[go語言]]
```
需要執行這些語法去初始化環境讓他能夠讀到
_ "github.com/go-sql-driver/mysql"
go env -w GO111MODULE=on
go mod init Gone 
go get github.com/go-sql-driver/mysql
```

```go
package main

  

import (

    "database/sql"

    "fmt"

  

    _ "mysql_master"

    // _ "github.com/go-sql-driver/mysql"

    //"time"

)

  

func main() {

    //db, err := sql.Open("mysql", "user:password@/dbname")

    db, err := sql.Open("mysql", "root:@/gosql?charset=utf8")

    checkErr(err)

  

    //插入資料

    stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")

    checkErr(err)

  

    res, err := stmt.Exec("astaxie", "研發部門", "2012-12-09")

    checkErr(err)

  

    id, err := res.LastInsertId()

    checkErr(err)

  

    fmt.Println(id)

    //更新資料

    stmt, err = db.Prepare("update userinfo set username=? where uid=?")

    checkErr(err)

  

    res, err = stmt.Exec("astaxieupdate", id)

    checkErr(err)

  

    affect, err := res.RowsAffected()

    checkErr(err)

  

    fmt.Println(affect)

  

    //查詢資料

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

  

    //刪除資料

    stmt, err = db.Prepare("delete from userinfo where uid=?")

    checkErr(err)

  

    res, err = stmt.Exec(id)

    checkErr(err)

  

    affect, err = res.RowsAffected()

    checkErr(err)

  

    fmt.Println(affect)

  

    db.Close()

}

  

func checkErr(err error) {

    if err != nil {

        panic(err)

    } else {

        println("success!")

    }

}
```