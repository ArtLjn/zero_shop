package logic

import (
	"context"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
	"zero_shop/app/user/internal/svc"
	"zero_shop/app/user/user"
	"zero_shop/pkg"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailLogic) SendEmail(in *user.SendEmailRequest) (*user.SendEmailResponse, error) {
	if err := pkg.CheckFieldsNoEmpty(in); err != nil {
		return &user.SendEmailResponse{
			Code: 400,
			Msg:  err.Error(),
		}, err
	}
	m := gomail.NewMessage()
	cf := l.svcCtx.Config.Email
	rand.Seed(time.Now().UnixNano())

	//生成一个随机6位数
	code := rand.Intn(900000) + 100000
	body := fmt.Sprintf("您好，您的验证码为\n%d\n请妥善保管，本次验证码5分钟内有效", code)
	m.SetHeader("From", cf.Username)
	m.SetHeader("To", in.Email)
	m.SetHeader("Subject", "注册邮件")
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(cf.Host, cf.Port, cf.Username, cf.Password)
	if err := d.DialAndSend(m); err != nil {
		return &user.SendEmailResponse{
			Code: 400,
			Msg:  err.Error(),
		}, err
	}

	// 将验证码存入redis
	_, err := l.svcCtx.Rdb.Set(context.Background(), in.Email, code, 5*time.Minute).Result()
	if err != nil {
		return &user.SendEmailResponse{
			Code: 500,
			Msg:  err.Error(),
		}, err
	}
	return &user.SendEmailResponse{
		Code: 200,
		Msg:  "发送成功",
	}, nil
}
