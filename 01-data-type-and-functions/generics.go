package ch01

import (
	"cmp"
	"fmt"
	"reflect"
)

//	func isLarger[T int | int32 | int64](a, b T) bool {
//		return a > b
//	}
func isLarger[T cmp.Ordered](a, b T) bool {
	return a > b
}

type (
	Offers             struct{}
	OffersResponse     struct{}
	Incentives         struct{}
	IncentivesResponse struct{}
)

func dto[In Offers | Incentives, Out OffersResponse | IncentivesResponse](data In) Out {
	tp := reflect.TypeOf(data)
	switch tp.Name() {
	case "Offers":
		fmt.Println("do dto.Offers(data)")
		return any(OffersResponse{}).(Out)
	case "Incentives":
		fmt.Println("do dto.Incentives(data)")
		return any(IncentivesResponse{}).(Out)
	default:
		panic("Invalid input type")
	}
}

func GenericRun() {
	a1 := 10
	b1 := 20
	isLarger(a1, b1)
	a2 := int32(a1)
	b2 := int32(b1)
	isLarger(a2, b2)

	dto[Offers, OffersResponse](Offers{})
	dto[Incentives, IncentivesResponse](Incentives{})
}
