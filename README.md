# About

This is a tool for dumping a given SQL Server certificate chain into local files.

The encrypted SQL Server Tabular Data Stream (TDS) protocol works above, with the exception of the first pre-login message, a standard TLS layer. But because of that initial message we cannot use regular tools (e.g. `openssl s_client`) to troubleshoot certificate issues, hence this tool exists.

This tool uses a [modified version](https://github.com/rgl/dump-sql-server-certificate-chain-go-mssqldb) of the [denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb) driver.

# Build

Setup the Go workspace:

	mkdir -p dump-sql-server-certificate-chain/src/github.com/rgl/dump-sql-server-certificate-chain
	cd dump-sql-server-certificate-chain
	git clone --recursive https://github.com/rgl/dump-sql-server-certificate-chain src/github.com/rgl/dump-sql-server-certificate-chain
	export GOPATH=$PWD
	export PATH=$PWD/bin:$PATH
	hash -r # reset bash path

Build:

	cd src/github.com/rgl/dump-sql-server-certificate-chain
	go get
	go build

Execute:

    ./dump-sql-server-certificate-chain -server sql.example.com

List the dumped chain certificates:

    ls -l *.der
