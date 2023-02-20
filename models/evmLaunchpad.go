package models

type EvmLaunchpadHeader struct{}

type EvmLaunchpad struct {
	EoaAddress      string `json:"eoa_address" binding:"required,startswith=0x"`
	ContractAddress string `json:"contract_address" binding:"required,startswith=0x"`
	NetworkChainId  int64  `json:"network_chain_id" binding:"required"`
	Price           int64  `json:"price" binding:"required,min=-0"`
	MetaDataUri     string `json:"meta_data_uri" binding:"required"`
	RawTransaction  string `json:"raw_transaction" binding:"required"`
}
