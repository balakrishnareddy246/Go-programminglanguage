package main

import (
      "io/ioutil"
      "log"
      "net/http"
      
)


type Hello struct {
    l *log.Logger
}
fucn NewHello(l *log.logger) *Hello {
    return &Hello{l}
}

func (h, *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request)
    h,l,Println("Hello world")
return
})

fmt.Fprintf(rw, "Hello %s", d)
d, err := ioutil.ReadAll(r,Body)
}