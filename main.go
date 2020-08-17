package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	mssql "github.com/denisenkom/go-mssqldb"
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

	connector, err := mssql.NewConnector(connectionString)
	if err != nil {
		log.Fatal("Failed to create the database connector: ", err.Error())
		return
	}

	// save the server certificates to the current directory.
	connector.NewTLSConn = func(conn net.Conn, config *tls.Config) *tls.Conn {
		config.InsecureSkipVerify = true
		config.VerifyPeerCertificate = func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			for i, crt := range rawCerts {
				path := fmt.Sprintf("%s-%d.der", config.ServerName, i)
				log.Printf("Saving %s certificate chain link #%d to %s...", config.ServerName, i, path)
				ioutil.WriteFile(
					path,
					crt,
					0644)
			}
			return nil
		}
		return tls.Client(conn, config)
	}

	log.Printf("Connecting to %s:%d...", *server, *port)
	db := sql.OpenDB(connector)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Connect failed (BUT check the local directory for the server certificate chain links in .der files): ", err.Error())
	}
}
