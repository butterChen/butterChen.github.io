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

	// "math/rand"
	// "net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

var constr = "root:@/nodemcu?charset=utf8"

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
	fmt.Println("hum =", hum)
	fmt.Println("temp =", temp)
	fmt.Println("send data to insert")
	insertsql(hum, temp)

}

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
	fmt.Println("insert sql")
	db, err := sql.Open(sqltype, constr)
	checkErr(err)

	//插入資料
	stmt, err := db.Prepare("INSERT temp_humi SET temperature=?,humidity=?")
	checkErr(err)

	res, err := stmt.Exec(temp, hum)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	db.Close()

}

// 20231016 我們再這裡 從資料庫取直然後顯示在畫面上
func httpchart(w http.ResponseWriter, _ *http.Request) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line example in Westeros theme",
			Subtitle: "Line chart rendered by the http server this time",
		}))

	id := make([]opts.LineData, 0)
	temp := make([]opts.LineData, 0)
	humi := make([]opts.LineData, 0)
	date := make([]opts.LineData, 0)

	readsql(&id, &temp, &humi, &date)
	line.SetXAxis(date).
		AddSeries("temp", temp).
		AddSeries("humi", humi).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	line.Render(w)
}

func readsql(arid *[]opts.LineData, artemperature *[]opts.LineData, arhumidity *[]opts.LineData, arcreate_time *[]opts.LineData) {
	fmt.Println("lets test")
	db, err := sql.Open(sqltype, constr) //needcheck 20231005
	checkErr(err)

	//查詢資料 read_data(conn, sql)

	rows, err := db.Query("SELECT * FROM temp_humi")

	checkErr(err)

	// arid2 := make([]opts.LineData, 0)
	// artemperature2 := make([]opts.LineData, 0)
	// arhumidity2 := make([]opts.LineData, 0)
	// arcreate_time2 := make([]opts.LineData, 0)
	for rows.Next() {

		var id int

		var temperature string

		var humidity string

		var create_time string

		err = rows.Scan(&id, &temperature, &humidity, &create_time)

		checkErr(err)

		fmt.Println(id)

		*arid = append(*arid, opts.LineData{Value: id})
		fmt.Println(temperature)
		*artemperature = append(*artemperature, opts.LineData{Value: temperature})
		fmt.Println(humidity)
		*arhumidity = append(*arhumidity, opts.LineData{Value: humidity})
		fmt.Println(create_time)
		*arcreate_time = append(*arcreate_time, opts.LineData{Value: create_time})

	}
	db.Close()
	fmt.Println("lets test2 teacher test2")
}
