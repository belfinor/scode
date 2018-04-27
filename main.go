package main


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-04-26


import (
    "flag"
    "github.com/belfinor/Helium/daemon"
    "github.com/belfinor/Helium/db/ldb"
    "github.com/belfinor/Helium/log"
    "github.com/belfinor/scode/config"
    "net/http"
    "strconv"
)


func main() {

    conf   := ""
    is_daemon := false

    flag.StringVar( &conf, "c", "/etc/scode.json", "config file name" )
    flag.BoolVar( &is_daemon, "d", false, "run as daemon" )

    flag.Parse()

    cfg := config.Init( conf )

    if is_daemon {
      daemon.Run( &cfg.Daemon )
    }

    log.Init( &cfg.Log )

    if is_daemon {
        log.Info( "start application as daemon" )
    } else {
        log.Info( "start application" )
    }

    ldb.Init( &cfg.Database )

    http.HandleFunc( "/get", handleGet )

    log.Info( "start http server addr=" + cfg.Server.Host + ":" + strconv.Itoa(cfg.Server.Port) )
    http.ListenAndServe( cfg.Server.Host + ":" + strconv.Itoa(cfg.Server.Port), nil)
}

