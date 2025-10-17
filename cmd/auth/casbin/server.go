package casbin

import (
	"context"

	"entgo.io/ent/dialect"
	casbinv2 "github.com/casbin/casbin/v2"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/rpc/grpcx"
	entadapter "github.com/woocoos/casbin-ent-adapter"
	casbinent "github.com/woocoos/casbin-ent-adapter/ent"
	"github.com/woocoos/knockout-go/pkg/authz/casbin"
	"github.com/woocoos/knockout-go/pkg/authz/casbin/proto"
)

type ServerOption func(*Server)

func WithAuthDB(drv dialect.Driver) ServerOption {
	return func(srv *Server) {
		srv.drv = drv
	}
}

// Server 是基于casbin的鉴权Grpc服务
type Server struct {
	proto.UnimplementedCasbinServer
	enforcer casbinv2.IEnforcer
	client   *casbinent.Client

	drv        dialect.Driver
	grpcServer *grpcx.Server
}

func (s *Server) Start(ctx context.Context) error {
	return s.grpcServer.Start(ctx)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.grpcServer.Stop(ctx)
}

func NewServer(cnf *conf.AppConfiguration, opts ...ServerOption) (*Server, error) {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}
	s.client = casbinent.NewClient(casbinent.Driver(s.drv))
	adapter, err := entadapter.NewAdapterWithClient(s.client)
	if err != nil {
		return nil, err
	}
	au, err := casbin.NewAuthorizer(cnf.Sub("authz"), casbin.WithAdapter(adapter))
	if err != nil {
		return nil, err
	}
	s.enforcer = au.BaseEnforcer()

	s.grpcServer = grpcx.New(grpcx.WithConfiguration(cnf.Sub("casbinServer.grpc")))
	proto.RegisterCasbinServer(s.grpcServer.Engine(), s)

	return s, nil
}

func (s *Server) Enforce(ctx context.Context, req *proto.EnforceRequest) (*proto.BoolReply, error) {
	params := make([]any, 0, len(req.Params))
	for _, v := range req.Params {
		params = append(params, v)
	}
	res, err := s.enforcer.Enforce(params...)
	if err != nil {
		return nil, err
	}
	return &proto.BoolReply{
		Res: res,
	}, nil
}

func (s *Server) GetImplicitPermissionsForUser(ctx context.Context, req *proto.PermissionRequest) (*proto.Array2DReply, error) {
	res, err := s.enforcer.GetImplicitPermissionsForUser(req.User, req.Domain...)
	if err != nil {
		return nil, err
	}
	return s.wrapPlainPolicy(res), err
}

func (s *Server) wrapPlainPolicy(policy [][]string) *proto.Array2DReply {
	if len(policy) == 0 {
		return &proto.Array2DReply{}
	}

	policyReply := &proto.Array2DReply{}
	policyReply.D2 = make([]*proto.Array2DReplyD, len(policy))
	for e := range policy {
		policyReply.D2[e] = &proto.Array2DReplyD{D1: policy[e]}
	}

	return policyReply
}
