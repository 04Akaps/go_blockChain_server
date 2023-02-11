package services

// store *db.Store

import (
	"context"

	"go_blockChain_server/models"
	sqlc "go_blockChain_server/mysql/sqlc"
)

type EvmLaunchpadServiceImpl struct {
	query           *sqlc.Queries
	evmLaunchpadCtx context.Context
}

func NewEvmLaunchpadServiceImpl(ctx context.Context, query *sqlc.Queries) EvmLaunchpadService {
	return &EvmLaunchpadServiceImpl{query: query, evmLaunchpadCtx: ctx}
}

func (el *EvmLaunchpadServiceImpl) CreateNewLaunchpad(launchInfo *models.EvmLaunchpad) error {
	return nil
}

func (el *EvmLaunchpadServiceImpl) GetMyAllLaunchpad(eoaAddress *string) *[]models.EvmLaunchpad {
	return nil
}

func (el *EvmLaunchpadServiceImpl) GetMyLaunchpad(contractAddress *string) *models.EvmLaunchpad {
	return nil
}
