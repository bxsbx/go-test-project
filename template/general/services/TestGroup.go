package services

import (
	"context"
)

type TestGroupService struct {
	appCtx context.Context
}

func NewTestGroupService(appCtx context.Context) *TestGroupService {
	return &TestGroupService{appCtx: appCtx}
}

func (s *TestGroupService) TestFunc() (data string, err error) {
	return
}
