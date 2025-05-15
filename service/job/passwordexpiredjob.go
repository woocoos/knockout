package job

import (
	"context"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/api/msg"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/knockout/codegen/entgen/types"
	"github.com/woocoos/knockout/ent"
	"github.com/woocoos/knockout/ent/org"
	"github.com/woocoos/knockout/ent/orguser"
	"github.com/woocoos/knockout/ent/user"
	"github.com/woocoos/knockout/ent/useraddr"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PwdOptions struct {
	CronSpec string `yaml:"cronSpec" json:"cronSpec"`
	// 密码有效天数
	EffectiveDays string `yaml:"cronSpec" json:"effectiveDays"`
	// 常规提醒修改密码，离到期日还比较长
	RemindTimePeriod string `yaml:"remindTimePeriod" json:"remindTimePeriod"`
	// 即将到期提修改密码
	ExpiringTimePeriod string `yaml:"expiringTimePeriod" json:"expiringTimePeriod"`
}

type PasswordExpiredJob struct {
	Options *PwdOptions
	// 一个时间只能有一个job再running
	JobRunning bool

	db    *ent.Client
	kosdk *api.SDK
}

func NewPasswordExpiredJob(cfg *conf.Configuration) (*PasswordExpiredJob, error) {
	options := PwdOptions{
		EffectiveDays:      "365d",
		RemindTimePeriod:   "270d|180d|90d|30d",
		ExpiringTimePeriod: "14d|7d|3d|1d",
	}
	err := cfg.Unmarshal(&options)
	if err != nil {
		return nil, err
	}
	p := &PasswordExpiredJob{
		Options: &options,
	}
	return p, nil
}

func (p *PasswordExpiredJob) JobName() string {
	return "passwordJob"
}

func (p *PasswordExpiredJob) CronSpec() string {
	return p.Options.CronSpec
}

func (p *PasswordExpiredJob) JobFunc() {
	if p.JobRunning {
		logger.Info("job running")
		return
	}
	logger.Info("start job")
	p.JobRunning = true
	defer func() {
		p.JobRunning = false
	}()
	ctx := context.Background()
	// 检测密码
	ups, err := p.db.UserPassword.Query().WithUser(func(query *ent.UserQuery) {
		query.Where(user.StatusEQ(types.UserStatusActive))
	}).All(ctx)
	if err != nil {
		logger.Error("query user password error", zap.Error(err))
		return
	}
	p.checkPwd(ctx, ups)
}

func (p *PasswordExpiredJob) checkPwd(ctx context.Context, ups []*ent.UserPassword) {
	for _, up := range ups {
		// 如果密码处于disable则忽略
		if up.Status == typex.SimpleStatusDisabled {
			continue
		}
		date := up.UpdatedAt
		if date.IsZero() {
			// 没有更新时间则取创建时间
			date = up.CreatedAt
		}
		// 判断密码是否过期
		effectD, err := parseCustomDuration(p.Options.EffectiveDays)
		if err != nil {
			logger.Error("parse duration error", zap.Error(err))
			continue
		}
		if date.Add(effectD).Before(time.Now()) {
			// 密码已过期
			if up.Status == typex.SimpleStatusActive {
				// 密码过期设置密码状态为disabled
				err = p.db.UserPassword.UpdateOneID(up.ID).SetStatus(typex.SimpleStatusDisabled).Exec(ctx)
				if err != nil {
					logger.Error("update user password error", zap.Error(err))
				}
				// 发送邮件通知用户密码已过期，需重置密码能够登录
				tid, err := p.getUserTopOrgId(ctx, up.UserID)
				if err != nil {
					logger.Error("get user top org error", zap.Error(err))
					continue
				}
				usr, addr, err := p.getUserInfo(ctx, up.UserID)
				if err != nil {
					logger.Error("get user info error", zap.Error(err))
					continue
				}
				params := msg.PostableAlerts{
					{
						Annotations: map[string]string{
							"to":          addr.Email,
							"displayName": usr.DisplayName,
							"date":        time.Now().Format("2006-01-02"),
						},
						Alert: &msg.Alert{
							Labels: map[string]string{
								"receiver":  "email",
								"alertname": "UserPasswordExpired",
								"tenant":    strconv.Itoa(tid),
								"timestamp": strconv.Itoa(int(time.Now().Unix())),
							},
						},
					},
				}
				_ = p.postAlerts(ctx, params)
			}
		} else {
			// 常规提醒用户修改密码
			remindTimes := strings.Split(p.Options.RemindTimePeriod, "|")
			for _, t := range remindTimes {
				d, err := parseCustomDuration(t)
				if err != nil {
					logger.Error("parse duration error", zap.Error(err))
					continue
				}
				if isSameDate(date.Add(effectD), time.Now().Add(d)) {
					// 发送邮件通知客户密码已超过多久没改，需修改密码
					tid, err := p.getUserTopOrgId(ctx, up.UserID)
					if err != nil {
						logger.Error("get user top org error", zap.Error(err))
						continue
					}
					usr, addr, err := p.getUserInfo(ctx, up.UserID)
					if err != nil {
						logger.Error("get user info error", zap.Error(err))
						continue
					}
					months := (effectD.Hours() - d.Hours()) / 24 / 30
					params := msg.PostableAlerts{
						{
							Annotations: map[string]string{
								"to":          addr.Email,
								"displayName": usr.DisplayName,
								"months":      strconv.Itoa(int(months)),
							},
							Alert: &msg.Alert{
								Labels: map[string]string{
									"receiver":  "email",
									"alertname": "UserPasswordRemind",
									"tenant":    strconv.Itoa(tid),
									"timestamp": strconv.Itoa(int(time.Now().Unix())),
								},
							},
						},
					}
					_ = p.postAlerts(ctx, params)
					break
				}
			}
			// 提醒用户密码即将过期修改密码
			expiringTimes := strings.Split(p.Options.ExpiringTimePeriod, "|")
			for _, t := range expiringTimes {
				d, err := parseCustomDuration(t)
				if err != nil {
					logger.Error("parse duration error", zap.Error(err))
					continue
				}
				if isSameDate(date.Add(effectD), time.Now().Add(d)) {
					// 发送邮件通知客户密码即将过期，尽快修改密码，否则到期无法登录
					tid, err := p.getUserTopOrgId(ctx, up.UserID)
					if err != nil {
						logger.Error("get user top org error", zap.Error(err))
						continue
					}
					usr, addr, err := p.getUserInfo(ctx, up.UserID)
					if err != nil {
						logger.Error("get user info error", zap.Error(err))
						continue
					}
					days := d.Hours() / 24
					params := msg.PostableAlerts{
						{
							Annotations: map[string]string{
								"to":          addr.Email,
								"displayName": usr.DisplayName,
								"days":        strconv.Itoa(int(days)),
								"date":        time.Now().Add(d).Format("2006-01-02"),
							},
							Alert: &msg.Alert{
								Labels: map[string]string{
									"receiver":  "email",
									"alertname": "UserPasswordExpiring",
									"tenant":    strconv.Itoa(tid),
									"timestamp": strconv.Itoa(int(time.Now().Unix())),
								},
							},
						},
					}
					_ = p.postAlerts(ctx, params)
					break
				}
			}
		}
	}
}

func (p *PasswordExpiredJob) postAlerts(ctx context.Context, params msg.PostableAlerts) error {
	resp, err := p.kosdk.Msg().AlertAPI.PostAlerts(ctx, &msg.PostAlertsRequest{
		PostableAlerts: params,
	})
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	return fmt.Errorf(resp.Status)
}

func isSameDate(t1, t2 time.Time) bool {
	return t1.Truncate(24 * time.Hour).Equal(t2.Truncate(24 * time.Hour))
}

func parseCustomDuration(s string) (time.Duration, error) {
	if strings.HasSuffix(s, "d") {
		days, _ := strconv.Atoi(s[:len(s)-1])
		return time.Duration(days) * 24 * time.Hour, nil
	}
	return time.ParseDuration(s) // 默认支持的单位
}

func (p *PasswordExpiredJob) getUserTopOrgId(ctx context.Context, uid int) (int, error) {
	uorg, err := p.db.OrgUser.Query().Where(orguser.UserIDEQ(uid)).
		QueryOrg().Unique(false).Where(
		org.KindEQ(org.KindRoot),
		org.StatusEQ(typex.SimpleStatusActive),
	).Order(ent.Desc(org.FieldPath)).First(ctx)
	if err != nil {
		return 0, err
	}
	code := strings.Split(uorg.Path, "/")[0]
	oID, err := strconv.ParseInt(code, 36, 64)
	if err != nil {
		return 0, err
	}
	return int(oID), nil
}

func (p *PasswordExpiredJob) getUserInfo(ctx context.Context, uid int) (*ent.User, *ent.UserAddr, error) {
	usr, err := p.db.User.Get(ctx, uid)
	if err != nil {
		return nil, nil, err
	}
	addr, err := usr.QueryAddresses().Where(useraddr.AddrTypeEQ(useraddr.AddrTypeContact)).Only(ctx)
	if err != nil {
		return nil, nil, err
	}
	return usr, addr, nil
}
