package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/allegro/bigcache"
	_ "github.com/go-sql-driver/mysql"
	"github.com/laplace789/mysql_test/config"
)

var cfgDir = flag.String("conf", "", "config dir")

func init() {
	flag.Parse()
}

func main() {
	if *cfgDir == "" {
		panic("no config file")
	}
	cfg := config.Config(*cfgDir)

	// "username:password@tcp(host:port)/數據庫?charset=utf8"

	path := strings.Join([]string{cfg.Mysql.User, ":", cfg.Mysql.Passwd, "@tcp(", cfg.Mysql.Server, ":", cfg.Mysql.Port, ")/", cfg.Mysql.Database, "?charset=utf8"}, "")
	db, _ := sql.Open("mysql", path)

	if err := db.Ping(); err != nil {
		fmt.Println("opon database fail:", err)
		return
	}
	db.SetConnMaxLifetime(100)
	rows, err := db.Query("SELECT tac,manufacturer,model FROM dim_handsets")
	if err != nil {
		log.Fatal(err)
	}
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))
	for rows.Next() {
		var manufacturer, model string
		var tac string
		if err := rows.Scan(&tac, &manufacturer, &model); err != nil {
			log.Fatal(err)
		}
		cache.Set(tac, []byte(fmt.Sprint(manufacturer, " ", model)))
	}
	entry, _ := cache.Get("91138920")
	fmt.Println(string(entry))
}
