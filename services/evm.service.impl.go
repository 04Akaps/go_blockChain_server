package services

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
	arg := sqlc.CreateEvmLaunchpadParams{
		EoaAddress:      launchInfo.EoaAddress,
		ContractAddress: launchInfo.ContractAddress,
		NetworkChainID:  int32(launchInfo.NetworkChainId),
		Price:           int32(launchInfo.Price),
		MetaDataUri:     launchInfo.MetaDataUri,
	}

	_, err := el.query.CreateEvmLaunchpad(el.evmLaunchpadCtx, arg)
	if err != nil {
		return err
	}

	return nil
}

func (el *EvmLaunchpadServiceImpl) GetMyAllLaunchpad(eoaAddress string) ([]sqlc.EvmLaunchpad, error) {
	result, err := el.query.GetMyAllLaunchpad(el.evmLaunchpadCtx, eoaAddress)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (el *EvmLaunchpadServiceImpl) GetOneLaunchpad(caAddress string) (*models.EvmLaunchpad, error) {
	result, err := el.query.GetLaunchpad(el.evmLaunchpadCtx, caAddress)
	if err != nil {
		return nil, err
	}
	returnValue := models.EvmLaunchpad{
		EoaAddress:      result.EoaAddress,
		ContractAddress: result.ContractAddress,
		NetworkChainId:  int64(result.NetworkChainID),
		Price:           int64(result.Price),
		MetaDataUri:     result.MetaDataUri,
	}
	return &returnValue, nil
}

func (el *EvmLaunchpadServiceImpl) DeleteAllLaunchpadByAdmin() {
}
