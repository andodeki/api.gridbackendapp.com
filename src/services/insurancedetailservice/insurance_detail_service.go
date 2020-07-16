package insurancedetail

import "github.com/andodeki/code/HA/api.gridbackendapp.com/src/repository/db"

type InsuranceDetailsServiceInterface interface {
}
type insuranceDetailService struct {
	insurencedetaildbRepo db.InsuranceDetailsDBRepository
}

func NewInsuranceDetailsServiceInterface(insurencedetaildbRepo db.InsuranceDetailsDBRepository) InsuranceDetailsServiceInterface {
	return &insuranceDetailService{
		insurencedetaildbRepo: insurencedetaildbRepo,
	}
}
