package config


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-04-26


import (
  "github.com/belfinor/Helium/daemon"
  "github.com/belfinor/Helium/db/ldb"
  "github.com/belfinor/Helium/log"
  "github.com/belfinor/Helium/net/jsonrpc2"
  "encoding/json"
  "fmt"
  "io/ioutil"
)


type Config struct {
  Daemon   daemon.Config `json:"daemon"`
  Log      log.Config `json:"log"`
  Server   jsonrpc2.HttpConfig  `json:"server"`
  Database ldb.Config `json:"database"`
}


var _config *Config


func Load( filename string ) *Config {

  data, err := ioutil.ReadFile( filename )

  if err != nil {
    panic( fmt.Sprintf( "Open file %s error", filename ) )
  }

  var _con Config

  err = json.Unmarshal( data, &_con )

  if err != nil {
    panic( err )
  }

  return &_con
}


func Set( conf *Config ) {
  _config = conf
}


func Init( conf string) *Config {
  if _config == nil {
    _config = Load( conf )
  }

  return _config
}


func Get() *Config {
  return _config
}

