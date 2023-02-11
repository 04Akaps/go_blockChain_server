package services

import (
	"context"

	"go_blockChain_server/models"
)

type TestServiceImpl struct {
	testCollection []*models.Test
	testCtx        context.Context
}

func NewTestService(testList []*models.Test, ctx context.Context) TestService {
	return &TestServiceImpl{testCollection: testList, testCtx: ctx}
}

func (t *TestServiceImpl) CreateTest(newTest *models.Test) error {
	t.testCollection = append(t.testCollection, newTest)
	return nil
}

func (t *TestServiceImpl) GetTests() []*models.Test {
	return t.testCollection
}
