package client

import (
	"StandardProject/common/http"
	"context"
)

type LoginClient struct {
	AppCtx context.Context
	Client *http.Client
}

func NewLoginClient(appCtx context.Context) *LoginClient {
	return &LoginClient{AppCtx: appCtx, Client: http.DefaultClient()}
}

func (l *LoginClient) AuthToken(accessToken string, thisType, userType, userId string) error {
	path := "/v1/auth/access"
	queryMap := make(map[string]string)
	queryMap["ACCESSTOKEN"] = accessToken
	queryMap["platform"] = thisType
	queryMap["userType"] = userType
	queryMap["userId"] = userId
	err := l.Client.Post(path, queryMap, nil, nil, l.AppCtx)
	return err
}
