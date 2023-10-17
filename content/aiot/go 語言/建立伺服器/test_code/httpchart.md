```go
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

    a := make([]opts.LineData, 0)

    b := make([]opts.LineData, 0)

    c := make([]opts.LineData, 0)

    d := make([]opts.LineData, 0)

    //arcreate_time := make([]opts.LineData, 0)

    readsql(&a, &b, &c, &d)

    line.SetXAxis(d).

        AddSeries("Category A", b).

        AddSeries("Category B", c).

        SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

    line.Render(w)

}

```