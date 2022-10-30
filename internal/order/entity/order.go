package entity

import "errors"

type Order struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotal() (int, error)
}

func NewOrder(id string, price, tax float64) (*Order, error) {
	order := &Order{ID: id, Price: price, Tax: tax}
	if err := order.Validate(); err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("order ID should not be empty")
	}
	if o.Price == 0 {
		return errors.New("order price should not be empty")
	}
	if o.Tax == 0 {
		return errors.New("order tax should not be empty")
	}
	return nil
}
