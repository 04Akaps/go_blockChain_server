package middleware

import (
	"errors"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

var re = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

func CheckCreatelaunchpadAddress(ctx *gin.Context, client *ethclient.Client, eoaAddress string, caAddress string) (bool, error) {
	eoaAddressIsBool := re.MatchString(eoaAddress)
	caAddressIsBool := re.MatchString(caAddress)

	if !eoaAddressIsBool {
		return false, errors.New("message : EoaAddress is Not Boolean")
	}

	if !caAddressIsBool {
		return false, errors.New("message : CaAddress is Not Boolean")
	}

	address := common.HexToAddress(caAddress)
	bytecode, err := client.CodeAt(ctx, address, nil) // nill is latest block
	if err != nil {
		return false, errors.New("message : codeAt is Failed")
	}

	isContract := len(bytecode) > 0

	if !isContract {
		return false, errors.New("message : caAddress is Not Contract")
	}

	return true, nil
}
