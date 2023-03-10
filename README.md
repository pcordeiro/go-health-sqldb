# go-health-sqldb
Health check implementation for SQL databases - To be used with [go-health](https://github.com/pcordeiro/go-health)

#### Usage
Get the package
```bash
go get -u github.com/pcordeiro/go-health-sqldb
```

In the code:
```go
import(
    _ "github.com/denisenkom/go-mssqldb" // the sql database driver for the database health check

   	"github.com/pcordeiro/go-health"
	health_sqldb "github.com/pcordeiro/go-health-sqldb"
)

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

// set the router (which ever one you like. In this example I'm using fiber)
router.Get("/", func(ctx *fiber.Ctx) error {
    // execute the checks
    result := health.Check(ctx.Context())

    if result.Status != "OK" {
        ctx.Status(fiber.StatusServiceUnavailable)
    } else {
        ctx.Status(fiber.StatusOK)
    }

    return ctx.JSON(result)
})
```