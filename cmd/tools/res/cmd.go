package res

import (
	"context"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/entc/load"
	"fmt"
	"github.com/99designs/gqlgen/codegen/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mitchellh/mapstructure"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/urfave/cli/v2"
	"github.com/woocoos/entco/ecx"
	"github.com/woocoos/entco/pkg/authorization"
	"github.com/woocoos/entco/pkg/snowflake"
	"github.com/woocoos/entco/schemax"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/app"
	"github.com/woocoos/knockout/ent/appaction"
	"github.com/woocoos/knockout/ent/appres"
	_ "github.com/woocoos/knockout/ent/runtime"
	"log"
	"path/filepath"
	"strings"
)

type Config struct {
	AppConfig string
	EntConfig string
	GQLConfig string
	Dialect   string
	DSN       string
	AppCode   string
}

const (
	arnSplit = authorization.ArnSplit
)

var (
	fieldFormat = func(field string) string {
		return fmt.Sprintf("%s/*%s", field, arnSplit)
	}
)

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
				},
				&cli.PathFlag{
					Name:    "config",
					Usage:   "the knockout config",
					Value:   "./cmd/adminx/etc/app.yaml",
					Aliases: []string{"f"},
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
					AppCode:   ctx.String("app"),
				})
			},
		},
		{
			Name:  "ent",
			Usage: "gen resource from ent schema. Root schema will generate to app_res",
			Flags: []cli.Flag{
				&cli.PathFlag{
					Name:    "schema",
					Usage:   "the schema path",
					Value:   "./codegen/entgen/schema",
					Aliases: []string{"e"},
				},
				&cli.PathFlag{
					Name:    "config",
					Usage:   "the knockout config",
					Value:   "./cmd/adminx/etc/app.yaml",
					Aliases: []string{"f"},
				},
				&cli.StringFlag{
					Name:    "app",
					Usage:   "application code",
					Aliases: []string{"a"},
				},
			},
			Action: func(ctx *cli.Context) error {
				return GenEntSchemaRes(Config{
					AppConfig: ctx.String("config"),
					EntConfig: ctx.String("gql"),
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

// GenEntSchemaRes generate resource from ent schema
func GenEntSchemaRes(cfg Config) error {
	frame := initApp(cfg)
	entg, err := entc.LoadGraph(cfg.EntConfig, &gen.Config{})
	if err != nil {
		return err
	}
	client, err := ent.Open(frame.AppConfiguration().String("store.portal.driverName"),
		frame.AppConfiguration().String("store.portal.dsn"), ent.Debug())
	if err != nil {
		return err
	}
	appid := client.App.Query().Where(app.Code(cfg.AppCode)).OnlyX(context.Background()).ID

	builder := make([]*ent.AppResCreate, 0)
	for _, schema := range entg.Schemas {
		has, err, arnfn := checkResAnnotation(schema)
		if err != nil {
			return err
		}
		if !has {
			continue
		}
		inst, err := client.AppRes.Query().Where(appres.AppID(appid), appres.TypeName(schema.Name)).First(context.Background())
		if err != nil && !ent.IsNotFound(err) {
			return err
		}

		arn := strings.Builder{}
		if checkTenant(schema) {
			arn.WriteString(arnSplit + schemax.FieldTenantID + arnSplit)
		}
		arn.WriteString(schema.Name + arnSplit)
		arn.WriteString(arnPicker(schema, arnfn))
		apc := client.AppRes.Create().SetName(schema.Name).SetTypeName(schema.Name).
			SetArnPattern(arn.String()).SetCreatedBy(0).SetAppID(appid)
		builder = append(builder, apc)
		if inst != nil {
			apc.SetID(inst.ID)
		}
	}
	err = ecx.WithTx(context.Background(), func(ctx context.Context) (ecx.Transactor, error) {
		return client.Tx(ctx)
	}, func(itx ecx.Transactor) error {
		tx := itx.(*ent.Tx)
		err := tx.AppRes.CreateBulk(builder...).OnConflict().UpdateNewValues().Exec(context.Background())
		return err
	})
	if err == nil {
		log.Print("done")
	}
	return nil
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

// 检查schema是否为根对象
func checkResAnnotation(sch *load.Schema) (has bool, err error, fn arnFieldFunc) {
	var ann schemax.Annotation
	for ak, vals := range sch.Annotations {
		if ak == schemax.EntCoAnnotationName {
			err = mapstructure.Decode(vals, &ann)
			if err != nil {
				return
			}
		}
	}
	if len(ann.Resources) > 0 {
		has = true
		fn = func(name string) bool {
			for _, res := range ann.Resources {
				// exclude tenant_id and org_id
				if res == "tenant_id" || res == ann.TenantField {
					continue
				}
				if res == name {
					return true
				}
			}
			return false
		}
	}
	return
}

// 检查是否具有租户字段
func checkTenant(sch *load.Schema) bool {
	for _, field := range sch.Fields {
		if field.Name == schemax.FieldTenantID {
			return true
		}
	}
	var ann schemax.Annotation
	for ak, vals := range sch.Annotations {
		if ak == schemax.EntCoAnnotationName {
			err := mapstructure.Decode(vals, &ann)
			if err != nil {
				panic(err)
			}
		}
	}
	if ann.TenantField != "" {
		return true
	}
	return false
}

type arnFieldFunc func(name string) bool

func arnPicker(sch *load.Schema, fns ...arnFieldFunc) string {
	arn := strings.Builder{}
	for _, field := range sch.Fields {
		for _, fn := range fns {
			if fn(field.Name) {
				arn.WriteString(fieldFormat(field.Name))
			}
		}
	}
	return arn.String()
}
