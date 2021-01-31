package validators

import (
	"errors"
	"github.com/diogoqds/routes-challenge-api/repositories"
)

type RouteWithSellerIdValidator interface {
	RouteWithSellerId(sellerId int) error
}

type routeSellerValidator struct{}

func (r routeSellerValidator) RouteWithSellerId(sellerId int) error {
	routeWithSellerId, err := repositories.RouteRepo.FinderSellerRoute.FindRouteBySellerId(sellerId)

	if err != nil {
		return err
	}

	if routeWithSellerId.Id != 0 {
		return errors.New("already exists a route with this seller")
	}

	return nil
}

var (
	RouteSellerValidator = routeSellerValidator{}
)
