https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go

```go
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	  fmt.Printf("server: %s /\n", r.Method)
	  fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
	  fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))
	  fmt.Printf("server: headers:\n")
	  for headerName, headerValue := range r.Header {
		  fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
	  }

	  reqBody, err := ioutil.ReadAll(r.Body)
	  if err != nil {
			 fmt.Printf("server: could not read request body: %s\n", err)
	  }
	  fmt.Printf("server: request body: %s\n", reqBody)

	  fmt.Fprintf(w, `{"message": "hello!"}`)
  })
```