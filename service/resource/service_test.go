package resource

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/suite"
	"github.com/tsingsun/woocoo"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/test"
	"os"
	"path/filepath"
	"testing"
)

type ServiceSuite struct {
	Service     *Service
	redisServer *miniredis.Miniredis
}

type resourceSuite struct {
	*ServiceSuite
	suite.Suite
}

func TestResourceSuite(t *testing.T) {
	ServiceSuite := &ServiceSuite{}
	ts := &resourceSuite{
		ServiceSuite: ServiceSuite,
	}
	suite.Run(t, ts)
}

func (rs *resourceSuite) SetupSuite() {
	file := filepath.Join(test.BaseDir(), "testdata", "etc", "app.yaml")
	bs, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	cnf := conf.NewFromBytes(bs, conf.WithBaseDir(test.BaseDir())).AsGlobal()
	app := woocoo.New(woocoo.WithAppConfiguration(cnf))
	rs.Service.Client, err = ent.Open(app.AppConfiguration().String("store.test.driverName"), app.AppConfiguration().String("store.test.dsn"), ent.Debug())
	rs.Require().NoError(err)
	rs.redisServer = miniredis.RunT(rs.T())
}

func (rs *resourceSuite) TestGrant() {

}
