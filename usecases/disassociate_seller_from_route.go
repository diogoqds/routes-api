package usecases

import (
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type DisassociateSeller interface {
	Disassociate(id int) (bool, error)
}

type DisassociateSellerUseCase struct {
}

func (d DisassociateSellerUseCase) Disassociate(id int) (bool, error) {
	return repositories.RouteRepo.RouteSellerDeleter.Disassociate(id)
}

var (
	DisassociateSellerService DisassociateSeller = DisassociateSellerUseCase{}
)
