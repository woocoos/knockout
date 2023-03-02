package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/stretchr/testify/assert"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/enttest"
	"testing"
)

func TestCreateDb(t *testing.T) {
	enttest.Open(t, "mysql", "root:@tcp(localhost:3306)/portal?parseTime=true&loc=Local",
		enttest.WithMigrateOptions(schema.WithDropIndex(true), schema.WithDropColumn(true), schema.WithForeignKeys(false)))
}

func TestGlobalID(t *testing.T) {
	user := ent.User{}
	user.ID = 100
	s, err := user.GlobalID(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "dXNlcjoxMDA=", s)

	c := ent.Cursor{ID: 198555049289472}
	bs := bytes.NewBuffer([]byte{})
	c.MarshalGQL(bs)
	var gid = struct {
		ID int `json:"id"`
	}{
		ID: 198555049289472,
	}
	gidbs := bytes.NewBuffer([]byte{})
	quote := []byte{'"'}
	gidbs.Write(quote)
	wc := base64.NewEncoder(base64.RawStdEncoding, gidbs)
	err = msgpack.NewEncoder(wc).Encode(gid)
	gidbs.Write(quote)
	assert.NoError(t, err)
	t.Log(gidbs.String())
	t.Log(bs.String())
}
