package resource

import "github.com/woocoos/knockout/ent"

// Service 企业目录服务管理
type Service struct {
	Client        *ent.Client
	FilePrefixDir string
}
