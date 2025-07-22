package job

import (
	"context"
	"github.com/robfig/cron/v3"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout/ent"
	"strings"
)

var logger = log.Component("logCenter job")

type Server struct {
	cnf *conf.Configuration

	isStartCron       bool
	cron              *cron.Cron
	CronSpecJobRelate map[string]map[string]cron.EntryID
}

func NewServer(cnf *conf.Configuration) (*Server, error) {
	s := &Server{
		cnf:               cnf,
		isStartCron:       true,
		CronSpecJobRelate: make(map[string]map[string]cron.EntryID),
	}
	s.buildCron()
	return s, nil
}

// Start implements woocoo.Server but do noting in start, the web server has registered by NewServer.
func (s *Server) Start(ctx context.Context) error {
	if s.isStartCron {
		log.Info("start cron")
		s.cron.Start()
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.cron.Stop()
	return nil
}

func (s *Server) InitJobs(portalDB *ent.Client, kosdk *api.SDK) (err error) {
	if s.cnf.IsSet("password") {
		var pej *PasswordExpiredJob
		pej, err = NewPasswordExpiredJob(s.cnf.Sub("password"))
		if err != nil {
			return
		}
		pej.db = portalDB
		pej.kosdk = kosdk
		// 添加任务
		err = s.AddJob(pej)
		if err != nil {
			return
		}
	}
	return
}

func (s *Server) buildCron() {
	if s.cron != nil {
		s.cron.Stop()
	}
	s.cron = cron.New(cron.WithSeconds())
}

// AddJob 添加调度任务
func (s *Server) AddJob(job Job) error {
	cronSpec := job.CronSpec()
	existMap := make(map[string]int)
	cronSpecJobRelate := s.CronSpecJobRelate[job.JobName()]
	if cronSpecJobRelate == nil {
		cronSpecJobRelate = make(map[string]cron.EntryID)
		s.CronSpecJobRelate[job.JobName()] = cronSpecJobRelate
	}
	// 逗号分割cronSpec，并循环
	for _, spec := range strings.Split(cronSpec, ",") {
		if spec == "" {
			continue
		}
		existMap[spec] = 1
		// 如果spec已经存在于map中，则跳过
		if _, ok := cronSpecJobRelate[spec]; ok {
			continue
		}
		jobId, err := s.cron.AddFunc(spec, job.JobFunc)
		if err != nil {
			return err
		}
		cronSpecJobRelate[spec] = jobId
	}
	// 循环s.CronSpecJobRelate的key，如果不存在于cronSpec则移除
	for spec := range cronSpecJobRelate {
		if _, ok := existMap[spec]; ok {
			continue
		}
		s.cron.Remove(cronSpecJobRelate[spec])
		delete(cronSpecJobRelate, spec)
	}
	return nil
}
