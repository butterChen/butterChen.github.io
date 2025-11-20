```c
package main

  

/*

搜尋

needchecks

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

  

    "math/rand"

    // "net/http"

  

    "github.com/go-echarts/go-echarts/v2/charts"

    "github.com/go-echarts/go-echarts/v2/opts"

    "github.com/go-echarts/go-echarts/v2/types"

)

  

//"time"

  

func dhtHandler(w http.ResponseWriter, r *http.Request) {

    fmt.Println("dht")

    if r.URL.Path != "/dht" {

        http.Error(w, "404 not found.", http.StatusNotFound)

        return

    }

  

    if r.Method != "GET" {

        http.Error(w, "Method is not supported.", http.StatusNotFound)

        return

    }

    if err := r.ParseForm(); err != nil {

        fmt.Fprintf(w, "ParseForm() err: %v", err)

        return

    }

  

    // name := r.FormValue("name")

    // address := r.FormValue("address")

    hum := r.URL.Query().Get("Humidity")

    temp := r.URL.Query().Get("Temperature")

    fmt.Println("hum ", hum)

    fmt.Println("temp =", temp)

    insertsql(hum, temp)

  

}

  

func sayhelloName(w http.ResponseWriter, r *http.Request) {

    r.ParseForm()       //解析參數，預設是不會解析的

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

    fmt.Println("lets test")

    http.HandleFunc("/", sayhelloName) //設定存取的路由

    http.HandleFunc("/dht", dhtHandler)

    http.HandleFunc("/httpchart", httpchart)

  

    err := http.ListenAndServe(":9090", nil) //設定監聽的埠

    if err != nil {

        log.Fatal("ListenAndServe: ", err)

    }

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

  

func insertsql(hum string, temp string) {

  

    db, err := sql.Open(sqltype, "root:@/nodemcu?charset=utf8")

    checkErr(err)

  

    //插入資料

    stmt, err := db.Prepare("INSERT temp_humi SET temperature=?,humidity=?")

    checkErr(err)

  

    res, err := stmt.Exec(hum, temp)

    checkErr(err)

  

    id, err := res.LastInsertId()

    checkErr(err)

    fmt.Println(id)

  

    db.Close()

  

}

  

// generate random data for line chart

func generateLineItems() []opts.LineData {

    items := make([]opts.LineData, 0)

    for i := 0; i < 7; i++ {

        items = append(items, opts.LineData{Value: rand.Intn(300)})

    }

    return items

}

  

// 20231016 我們再這裡 從資料庫取直然後顯示在畫面上

func httpchart(w http.ResponseWriter, _ *http.Request) {

    // create a new line instance

    line := charts.NewLine()

    // set some global options like Title/Legend/ToolTip or anything else

    line.SetGlobalOptions(

        charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),

        charts.WithTitleOpts(opts.Title{

            Title:    "Line example in Westeros theme",

            Subtitle: "Line chart rendered by the http server this time",

        }))

  

    // Put data into instance

    // line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).

    //  AddSeries("Category A", generateLineItems()).

    //  AddSeries("Category B", generateLineItems()).

    //  SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

    // line.Render(w)

    readsql(a,b,c,d)

    line.SetXAxis(d).

        AddSeries("Category A",b).

        AddSeries("Category B", c).

        SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

    line.Render(w)

}

  

func readsql(&arid []opts.LineData,&artemperature []opts.LineData,&arhumidity []opts.LineData,&arcreate_time []opts.LineData) {

    fmt.Println("lets test")

    db, err := sql.Open(sqltype, "root:@/nodemcu?charset=utf8") //needcheck 20231005

    checkErr(err)

  

    //查詢資料 read_data(conn, sql)

  

    rows, err := db.Query("SELECT * FROM temp_humi")

  

    checkErr(err)

  

    arid := make([]opts.LineData, 0)

    artemperature := make([]opts.LineData, 0)

    arhumidity := make([]opts.LineData, 0)

    arcreate_time := make([]opts.LineData, 0)

    for rows.Next() {

  

        var id int

  

        var temperature string

  

        var humidity string

  

        var create_time string

  

        err = rows.Scan(&id, &temperature, &humidity, &create_time)

  

        checkErr(err)

  

        fmt.Println(id)

        arid = append(arid, opts.LineData{Value: id})

        fmt.Println(temperature)

        artemperature = append(artemperature, opts.LineData{Value: temperature})

        fmt.Println(humidity)

        arhumidity = append(arhumidity, opts.LineData{Value: humidity})

        fmt.Println(create_time)

        arcreate_time = append(arcreate_time, opts.LineData{Value: create_time})

    }

    db.Close()

    fmt.Println("lets test2")

}
```