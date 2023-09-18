package logic

import (
	"StandardProject/gozero/book/service/user/cmd/rpc/internal/user"
	"StandardProject/gozero/book/service/user/model"
	"context"

	"StandardProject/gozero/book/service/user/cmd/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReply, error) {
	if in.Id == 1 {
		one, err := l.svcCtx.UserModel.FindOne(2)
		if err != nil {
			return nil, err
		}
		return &user.UserInfoReply{
			Id:     one.Id,
			Name:   one.Name,
			Number: one.Number,
			Gender: one.Gender,
		}, nil
	} else if in.Id == 2 {
		a := model.User{
			Id:     2,
			Name:   "name",
			Number: "number",
			Gender: "nv",
		}
		_, err := l.svcCtx.UserModel.Insert(a)
		if err != nil {
			return nil, err
		}
		return &user.UserInfoReply{
			Id: 2,
		}, nil
	} else if in.Id == 3 {
		a := model.User{
			Id:     2,
			Name:   "bbbbb",
			Number: "1445",
			Gender: "kkk",
		}
		err := l.svcCtx.UserModel.Update(a)
		if err != nil {
			return nil, err
		}
		return &user.UserInfoReply{
			Id:     2,
			Number: a.Number,
		}, nil
	} else if in.Id == 4 {
		err := l.svcCtx.UserModel.Delete(2)
		if err != nil {
			return nil, err
		}
		return &user.UserInfoReply{
			Id: 4,
		}, nil
	}
	return nil, nil
}
