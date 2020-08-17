# About

This is a tool for dumping a given SQL Server certificate chain into local files.

The encrypted SQL Server Tabular Data Stream (TDS) protocol works above, with the exception of the first pre-login message, a standard TLS layer. But because of that initial message we cannot use regular tools (e.g. `openssl s_client`) to troubleshoot certificate issues, hence this tool exists.

This tool uses a [modified version](https://github.com/rgl/dump-sql-server-certificate-chain-go-mssqldb) of the [denisenkom/go-mssqldb](https://github.com/denisenkom/go-mssqldb) driver.

# Build

Install Go 1.15.

Build:

```bash
go build
```

Execute:

```bash
./dump-sql-server-certificate-chain -server sql.example.com
```

List the dumped chain certificates:

```bash
ls -l *.der
```

See one of them:

```bash
openssl \
    x509 \
    -noout \
    -text \
    -inform der \
    -in sql.example.com-0.der
```

