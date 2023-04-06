package res

import (
	"context"
	"github.com/99designs/gqlgen/codegen/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/urfave/cli/v2"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	_ "github.com/woocoos/knockout/ent/runtime"
	"log"
	"path/filepath"
	"strings"
)

type Config struct {
	AppConfig string
	GQLConfig string
	Dialect   string
	DSN       string
	AppCode   string
}

var Cmd = &cli.Command{
	Name:  "res",
	Usage: "tools for knockout resource generator",
	Subcommands: []*cli.Command{
		{
			Name:  "gql-action",
			Usage: "gen actions from graphql",
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:    "gql",
					Usage:   "the gqlgen config path",
					Value:   "./codegen/gqlgen/gqlgen.yaml",
					Aliases: []string{"c"},
					//FilePath: "./codegen/gqlgen/gqlgen.yaml",
				},
				&cli.PathFlag{
					Name:    "config",
					Usage:   "the knockout config",
					Value:   "./cmd/adminx/etc/app.yaml",
					Aliases: []string{"f"},
				},
				&cli.StringFlag{
					Name:    "dialect",
					Value:   "mysql",
					Aliases: []string{"d"},
					Usage:   "database dialect",
				},
				&cli.StringFlag{
					Name:  "dsn",
					Usage: "data source name (connection information)",
					Value: "root:@tcp(localhost:3306)/portal?parseTime=true&loc=Local",
				},
				&cli.StringFlag{
					Name:    "app",
					Usage:   "application code",
					Aliases: []string{"a"},
				},
			},
			Action: func(ctx *cli.Context) error {
				return GenGqlActions(Config{
					AppConfig: ctx.String("config"),
					GQLConfig: ctx.String("gql"),
					Dialect:   ctx.String("dialect"),
					DSN:       ctx.String("dsn"),
					AppCode:   ctx.String("app"),
				})
			},
		},
	},
}

func GenGqlActions(cfg Config) error {
	frame := initApp(cfg)
	gcfg, err := config.LoadConfig(cfg.GQLConfig)
	if err != nil {
		return err
	}
	err = gcfg.LoadSchema()
	if err != nil {
		return err
	}
	client, err := ent.Open(frame.AppConfiguration().String("store.portal.driverName"),
		frame.AppConfiguration().String("store.portal.dsn"), ent.Debug())
	if err != nil {
		return err
	}
	appid := client.App.Query().Where(app.Code(cfg.AppCode)).OnlyX(context.Background()).ID
	builder := make([]*ent.AppActionCreate, 0)
	for _, field := range gcfg.Schema.Query.Fields {
		if strings.HasPrefix(field.Name, "__") {
			continue
		}
		inst, err := client.AppAction.Query().Where(appaction.AppID(appid), appaction.Name(field.Name)).First(context.Background())
		if err != nil && !ent.IsNotFound(err) {
			return err
		}
		apc := client.AppAction.Create().SetName(field.Name).SetKind(appaction.KindGraphql).
			SetMethod(appaction.MethodRead).SetComments(field.Description).SetCreatedBy(0).SetAppID(appid)
		if inst != nil {
			apc.SetID(inst.ID)
		}
		builder = append(builder, apc)
	}
	for _, field := range gcfg.Schema.Mutation.Fields {
		inst, err := client.AppAction.Query().Where(appaction.AppID(appid), appaction.Name(field.Name)).First(context.Background())
		if err != nil && !ent.IsNotFound(err) {
			return err
		}
		apc := client.AppAction.Create().SetName(field.Name).SetKind(appaction.KindGraphql).
			SetMethod(appaction.MethodWrite).SetComments(field.Description).SetCreatedBy(0).SetAppID(appid)
		if inst != nil {
			apc.SetID(inst.ID)
		}
		builder = append(builder, apc)
	}
	err = ecx.WithTx(context.Background(), func(ctx context.Context) (ecx.Transactor, error) {
		return client.Tx(ctx)
	}, func(itx ecx.Transactor) error {
		tx := itx.(*ent.Tx)
		err := tx.AppAction.CreateBulk(builder...).OnConflict().UpdateNewValues().Exec(context.Background())
		return err
	})
	if err == nil {
		log.Print("done")
	}
	return err
}

func initApp(cfg Config) *woocoo.App {
	frame := woocoo.New(woocoo.WithAppConfiguration(
		conf.New(conf.WithBaseDir(filepath.Dir(filepath.Dir(cfg.AppConfig)))).Load()),
	)
	err := snowflake.SetDefaultNode(frame.AppConfiguration().Sub("snowflake"))
	if err != nil {
		log.Fatal(err)
	}
	return frame
}
