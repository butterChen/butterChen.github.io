```

break    default      func    interface    select
case     defer        go      map          struct
chan     else         goto    package      switch
const    fallthrough  if      range        type
continue for          import  return       var
```

``` go
channel, context, sync.WaitGroup, Select, sync.Mutex
```

#go
```
go func() {
		fmt.Println("goroutine Done!")
		PrintMemUsage()
		fmt.Println("goroutine Done!")
	}()
 
```
#defer
```
defer wg.Done() //defer表示最後執行，因此該行為最後執行wg.Done()將計數器-1
		defer log.Println("goroutine drop out")
```