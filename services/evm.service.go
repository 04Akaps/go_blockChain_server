package services

import "go_blockChain_server/models"

type EvmLaunchpadService interface {
	CreateNewLaunchpad(*models.EvmLaunchpad) error    // 새롭게 만들어진 Launchpad 저장용
	GetMyAllLaunchpad(*string) *[]models.EvmLaunchpad // 나의 모든 launchpad 보기
	GetMyLaunchpad(*string) *models.EvmLaunchpad      // 나의 launchpad 보기
}
