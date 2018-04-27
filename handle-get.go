package main


import (
  "fmt"
  "github.com/belfinor/Helium/db/ldb"
  "github.com/belfinor/Helium/net/http/errors"
  "github.com/belfinor/Helium/net/http/params"
  "net/http"
)


func handleGet( rw http.ResponseWriter, req *http.Request ) {
  if req.Method != "GET" {
    errors.Send( rw, 404 )
    return
  }

  args := params.New(req)
  code := args.GetString("code")

  val := ldb.Get( []byte(code) )
  if val == nil || len(val) == 0 {
    errors.Send( rw, 404 )
    return
  }

  fmt.Fprint( rw, string(val) )
}

