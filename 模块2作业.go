package main

import (
    "fmt"
    "net/http"
	"os"
)

func healthz(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200"))
}

func router(w http.ResponseWriter,r *http.Request)  {
   version := os.Getenv("VERSION")
   if len(r.Header) > 0 {
      for k,v := range r.Header {
		//fmt.Printf("%s=%s\n", k, v[0])
		w.Header().Set(k, v[0])
      }
   }
   w.Header().Set("version", version)
   w.Write([]byte("hello"))
   fmt.Println(r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", router)
    http.HandleFunc("/healthz", healthz)
    err := http.ListenAndServe("127.0.0.1:8000", nil)
    if err != nil {
        fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
        return
    }
}
