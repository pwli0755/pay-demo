package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"pay/internal/conf"
	"pay/internal/data/ent"

	"database/sql"
	"entgo.io/ent/dialect"
	"github.com/go-kratos/kratos/v2/log"
	//"github.com/go-redis/redis/extra/redisotel"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/google/wire"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	// init postgres driver
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewProductRepo)

// Data .
type Data struct {
	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		return nil, nil, err
	}
	sqlDrv2 := entsql.OpenDB(dialect.Postgres, drv)
	sqlDrv := dialect.DebugWithContext(sqlDrv2, func(ctx context.Context, i ...interface{}) {
		log.WithContext(ctx).Info(i...)
		tracer := otel.Tracer("ent.")
		kind := trace.SpanKindServer
		_, span := tracer.Start(ctx,
			"Query",
			trace.WithAttributes(
				attribute.String("sql", fmt.Sprint(i...)),
			),
			trace.WithSpanKind(kind),
		)
		span.End()
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: conf.Redis.Addr,
		//Password: conf.Redis.Password,
		//DB:       int(conf.Redis.Db),
		//DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	//rdb.AddHook(redisotel.TracingHook{})
	d := &Data{
		db:  client,
		rdb: rdb,
	}
	return d, func() {
		log.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
