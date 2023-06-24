package client

import (
	"StandardProject/common/request"
	"context"
)

func AuthToken(accessToken string, thisType, userType, userId string) error {
	path := "/v1/auth/access"
	queryMap := make(map[string]string)
	queryMap["ACCESSTOKEN"] = accessToken
	queryMap["platform"] = thisType
	queryMap["userType"] = userType
	queryMap["userId"] = userId

	err := request.PostQuery(path, queryMap, nil, request.RESPONSE2, request.GetClient(request.LOGIN), context.Background())
	return err
}
