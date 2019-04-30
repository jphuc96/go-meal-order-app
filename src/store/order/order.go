package order

import (
	"context"
	"database/sql"
	"errors"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
)

type orderService struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &orderService{
		db: db,
	}
}

func (os *orderService) Add(tx *sql.Tx, o *domain.OrderInput) (*models.Order, error) {
	order := &models.Order{
		UserID: o.UserID,
		ItemID: o.ItemID,
	}

	err := order.Insert(context.Background(), tx, boil.Infer())
	if err != nil {
		return nil, err
	}

	return order, err
}

func (os *orderService) Delete(tx *sql.Tx, o *domain.OrderInput) error {
	order, err := models.Orders(qm.Where("user_id=? and item_id=?", o.UserID, o.ItemID)).One(context.Background(), tx)
	if err != nil {
		return err
	}

	_, err = order.Delete(context.Background(), tx)

	return err
}

func (os *orderService) Exist(tx *sql.Tx, o *domain.OrderInput) (bool, error) {
	b, err := models.Orders(qm.Where("user_id=? and item_id=?", o.UserID, o.ItemID)).Exists(context.Background(), tx)
	if err != nil {
		return false, err
	}
	return b, nil
}

func (os *orderService) Get(menuID string, userID string) ([]*domain.Item, error) {
	b, err := models.Menus(qm.Where("id=?", menuID)).Exists(context.Background(), os.db)
	if err != nil {
		return nil, err
	}
	if b == false {
		return nil, errors.New(domain.MenuNotExist)
	}

	b, err = models.Users(qm.Where("id=?", userID)).Exists(context.Background(), os.db)
	if err != nil {
		return nil, err
	}
	if b == false {
		return nil, errors.New(domain.UserNotExist)
	}

	orders, err := models.Orders(qm.Where("user_id=?", userID)).All(context.Background(), os.db)
	if err != nil {
		return nil, err
	}

	items := make([]*models.Item, len(orders))
	returnItems := make([]*domain.Item, len(orders))

	for i, order := range orders {
		items[i], err = models.Items(qm.Where("id=? AND menu_id=?", order.ItemID, menuID)).One(context.Background(), os.db)
		if err != nil {
			return nil, err
		}

		returnItems[i] = &domain.Item{
			ID:       items[i].ID,
			ItemName: items[i].ItemName,
			MenuID:   items[i].MenuID,
		}
	}

	return returnItems, err
}

// DeleteOrder ..
func (s *orderService) DeleteOrder(tx *sql.Tx, o *models.Order) error {
	_, err := o.Delete(context.Background(), tx)
	return err
}

// CheckOrderExistByItemID ..
func (s *orderService) CheckOrderExistByItemID(tx *sql.Tx, ItemID int) (bool, error) {
	return models.Orders(
		qm.Where("item_id=?", ItemID),
	).Exists(context.Background(), tx)
}

// GetAllOrdersByItemID ..
func (s *orderService) GetAllOrdersByItemID(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
	return models.Orders(
		qm.Where("item_id=?", ItemID),
	).All(context.Background(), tx)
}

func (s *orderService) GetOrderByOrderInput(tx *sql.Tx, o *domain.OrderInput) (*models.Order, error) {
	return models.Orders(qm.Where("user_id=? AND item_id=?", o.UserID, o.ItemID)).One(context.Background(), tx)
}
