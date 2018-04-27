package main


import (
  "fmt"
  "github.com/belfinor/Helium/db/ldb"
  "github.com/belfinor/Helium/hash/code62"
  "github.com/belfinor/Helium/net/http/errors"
  "io/ioutil"
  "net/http"
  "sync/atomic"
  "time"
)


var cnt int64 = 0


func handleSet( rw http.ResponseWriter, req *http.Request ) {

    if req.Method != "POST" {
      errors.Send( rw, 404 )
      return
    }

    body, err := ioutil.ReadAll( req.Body )
    if err != nil {
      errors.Send( rw, 400 )
      return
    }

    val := ( time.Now().Unix() << 8 ) | ( atomic.AddInt64( &cnt, 1 ) & 0xff )

    code := code62.Calc( val )

    ldb.Set( []byte(code), body )

    fmt.Fprint( rw, code )
}

