package client

import (
	"StandardProject/common/http"
	"context"
)

func AuthToken(accessToken string, thisType, userType, userId string) error {
	path := "/v1/auth/access"
	queryMap := make(map[string]string)
	queryMap["ACCESSTOKEN"] = accessToken
	queryMap["platform"] = thisType
	queryMap["userType"] = userType
	queryMap["userId"] = userId

	err := http.PostQuery(path, queryMap, nil, http.RESPONSE2, http.GetClient(http.LOGIN), context.Background())
	return err
}
