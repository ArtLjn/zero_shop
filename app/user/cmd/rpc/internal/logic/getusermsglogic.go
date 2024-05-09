package logic

import (
	"context"
	"zero_shop/app/user/cmd/rpc/internal/svc"
	"zero_shop/app/user/cmd/rpc/user"
	"zero_shop/app/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMsgLogic {
	return &GetUserMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserMsgLogic) GetUserMsg(in *user.GetUserMsgRequest) (*user.GetUserMsgResponse, error) {
	u := l.svcCtx.Uc.GetUserCond(l.ctx, "user_id", in.UserId)
	if u == nil {
		return nil, model.ErrNotFound
	}
	var resp user.GetUserMsgResponse
	l.svcCtx.Tool.CopyProperties(u, &resp)
	return &resp, nil
}
