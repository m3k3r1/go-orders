package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/m3k3r1/go-orders/internal/order/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

//sql create table on mysql
//CREATE TABLE `orders` (
//	  `id` varchar(255) NOT NULL,
//	  `price` float NOT NULL,
//	  `tax` float NOT NULL,
//	  `final_price` float NOT NULL,
//	  PRIMARY KEY (`id`)
//	)

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	query, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = query.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}
