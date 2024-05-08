package logic

import (
	"context"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"zero_shop/app/user/cmd/rpc/internal/svc"
	"zero_shop/app/user/cmd/rpc/user"
	"zero_shop/app/user/model"
	"zero_shop/pkg"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register /**
// example:
//
//	{
//	   "phone":"111",
//	   "username":"test01",
//	   "password":"123",
//	   "email":"example@qq.com",
//	   "emailCode":"22222"
//	}
func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// check register email code
	cacheCode, _ := l.svcCtx.Rdb.Get(context.Background(), in.Email).Result()
	if cacheCode != in.EmailCode {
		return &user.RegisterResponse{
			Code: 400,
			Msg:  "register code is wrong",
		}, nil
	}
	var createUser model.User
	l.svcCtx.Tool.CopyProperties(in, &createUser)
	createUser.UserId = uuid.New().String()[:8]
	createUser.RoleId = pkg.CustomerRole
	createUser.Sex = "male"
	if l.svcCtx.Tool.IsEmpty(createUser, model.User{}) {
		return &user.RegisterResponse{
			Code: 400,
			Msg:  "register params is wrong",
		}, nil
	} else if l.svcCtx.Uc.GetUserCond(context.Background(), "email", createUser.Email) != nil {
		return &user.RegisterResponse{
			Code: 400,
			Msg:  "email is already exist",
		}, nil
	} else if l.svcCtx.Uc.GetUserCond(context.Background(), "phone", createUser.Phone) != nil {
		return &user.RegisterResponse{
			Code: 400,
			Msg:  "phone is already exist",
		}, nil
	}
	// encrypt password
	pkg.MD5(&createUser.Password)
	_, err := l.svcCtx.Uc.Insert(context.Background(), &createUser)
	if err != nil {
		return &user.RegisterResponse{
			Code: 500,
			Msg:  "register fail",
		}, nil
	}
	return &user.RegisterResponse{
		Code: 200,
		Msg:  "register success",
	}, nil
}
