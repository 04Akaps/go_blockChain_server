package models

type EvmLaunchpad struct {
	EoaAddress      string `json:"eoa_address"`
	ContractAddress string `json:"contract_address"`
	NetworkChainId  int64  `json:"network_chain_id"`
	Price           int64  `json:"price"`
	MetaDataUri     string `json:"meta_data_uri" binding:"required"`
}
