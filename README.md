# go-health-sqldb
Health check implementation for SQL databases - To be used with [go-health](https://github.com/pcordeiro/go-health)

#### Usage
Get the package
```bash
go get -u github.com/pcordeiro/go-health-sqldb
```

In the code:
```go
health, err := health.NewHealth(
    health.WithComponent(
        health.Component{
            Name:    app.config.Name,
            Version: app.config.Version,
        },
    ),
    health.WithChecks(
        health.Check{
            Name:      "Database",
            Timeout:   2 * time.Second,
            SkipOnErr: false,
            Check: sqldb.NewSqlDbCheck(&sqldb.Config{
                Name:   "MS SQL Server",
                Driver: config.Get().Database.Driver,
                DSN:    config.Get().Database.DSN,
                Select: "SELECT @@VERSION",
            }),
        },
    ),
)

```