package usecase

import "github.com/m3k3r1/go-orders/internal/order/entity"

type GetTotalOutputDTO struct {
	Total int `json:"total"`
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: orderRepository}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	total, err := c.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}
	return &GetTotalOutputDTO{Total: total}, nil
}
