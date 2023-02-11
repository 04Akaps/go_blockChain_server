package services

// store *db.Store

import (
	"context"

	"go_blockChain_server/models"
)

type EvmLaunchpadServiceImpl struct {
	// mysql db *sql.DB
	evmLaunchpadCtx context.Context
}

func NewEvmLaunchpadServiceImpl(ctx context.Context) EvmLaunchpadService {
	return &EvmLaunchpadServiceImpl{evmLaunchpadCtx: ctx}
}

func (el *EvmLaunchpadServiceImpl) CreateNewLunchpad(launchInfo *models.EvmLaunchpad) error {
	return nil
}

func (el *EvmLaunchpadServiceImpl) GetMyAllLaunchpad(eoaAddress *string) *[]models.EvmLaunchpad {
	return nil
}

func (el *EvmLaunchpadServiceImpl) GetMyLaunchpad(contractAddress *string) *models.EvmLaunchpad {
	return nil
}
