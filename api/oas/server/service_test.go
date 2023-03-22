package server

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/test"
	"os"
	"path/filepath"
	"testing"
)

type TestSuite struct {
	suite.Suite
	Service     *Service
	redisServer *miniredis.Miniredis
}

// TestSuite runs all the tests in the suite.
func TestTestSuite(t *testing.T) {
	service := Service{}
	file := filepath.Join(test.BaseDir(), "testdata", "etc", "app.yaml")
	bs, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	cnf := conf.NewFromBytes(bs, conf.WithBaseDir(test.BaseDir())).AsGlobal()
	app := woocoo.New(woocoo.WithAppConfiguration(cnf))
	Cnf = app.AppConfiguration()
	service.DB, err = ent.Open("mysql", app.AppConfiguration().String("store.portal.dsn"), ent.Debug())

	mr := miniredis.NewMiniRedis()
	service.Redis = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	suite.Run(t, &TestSuite{
		Service:     &service,
		redisServer: mr,
	})
}

func TestPwd(t *testing.T) {
	req := sha256.New()
	req.Write([]byte("123456"))
	param := hex.EncodeToString(req.Sum(nil))
	t.Log(param)
	sha := sha256.New()
	sha.Write([]byte(param))
	sha.Write([]byte("123456"))
	given := hex.EncodeToString(sha.Sum(nil))
	t.Log(given)
}
