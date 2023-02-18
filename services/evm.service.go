package services

import (
	"go_blockChain_server/models"
	sqlc "go_blockChain_server/mysql/sqlc"
)

type EvmLaunchpadService interface {
	CreateNewLaunchpad(*models.EvmLaunchpad) error         // 새롭게 만들어진 Launchpad 저장용
	GetMyAllLaunchpad(string) ([]sqlc.EvmLaunchpad, error) // 나의 모든 launchpad 보기
	DeleteAllLaunchpadByAdmin()                            // admin을 위한 함수
	GetOneLaunchpad(string) (*models.EvmLaunchpad, error)  // launchpad 하나만 보기
}
