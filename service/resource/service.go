package resource

import (
	"github.com/woocoos/knockout/ent"
	"net/http"
)

// Service 企业目录服务管理
type Service struct {
	Client     *ent.Client
	HttpClient *http.Client
	OASOptions OASOptions
}

type OASOptions struct {
	Files struct {
		PrefixDir string `yaml:"prefixDir"`
		BaseUrl   string `yaml:"baseUrl"`
	} `yaml:"files"`
	Msgsrv struct {
		BaseUrl string `yaml:"baseUrl"`
	} `yaml:"msgsrv"`
}
