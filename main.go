package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/rgl/dump-sql-server-certificate-chain-go-mssqldb"
)

var (
	server = flag.String("server", "", "e.g. sql.example.com")
	port   = flag.Int("port", 1433, "1433")
)

func main() {
	log.SetOutput(os.Stdout) // for not disturbing PowerShell...

	flag.Parse()

	if *server == "" {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	connectionString := fmt.Sprintf(
		"Server=%s; Port=%d; Encrypt=true; App Name=dump-sql-server-certificate-chain;",
		*server,
		*port)

	log.Printf("Connection to %s:%d...", *server, *port)
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Open connection failed: ", err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Connect failed (BUT check the local directory for the server certificate chain links in .der files): ", err.Error())
	}
}
