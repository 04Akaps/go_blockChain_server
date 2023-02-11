package services

import "go_blockChain_server/models"

type TestService interface {
	CreateTest(*models.Test) error
	GetTests() []*models.Test
}
