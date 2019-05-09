package service

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/store/item"
	"git.d.foundation/datcom/backend/src/store/order"
)

func TestService_CheckItemExist(t *testing.T) {
	type fields struct {
		Item item.ServiceMock
	}
	type args struct {
		tx     *sql.Tx
		itemID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "check exist success",
			fields: fields{
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return false, nil
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: Store{
					ItemStore: &tt.fields.Item,
				},
			}
			got, err := s.CheckItemExist(tt.args.tx, tt.args.itemID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CheckItemExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.CheckItemExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetItemByID(t *testing.T) {
	type fields struct {
		Item item.ServiceMock
	}
	type args struct {
		tx     *sql.Tx
		itemID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Item
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test get item by id",
			fields: fields{
				item.ServiceMock{
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return nil, nil
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: Store{
					ItemStore: &tt.fields.Item,
				},
			}
			got, err := s.GetItemByID(tt.args.tx, tt.args.itemID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetItemByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetItemByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DeleteItem(t *testing.T) {
	type fields struct {
		Item  item.ServiceMock
		Order order.ServiceMock
	}
	type args struct {
		tx     *sql.Tx
		itemID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "s.Store.ItemStore.FindByID fail",
			fields: fields{
				item.ServiceMock{
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return &models.Item{}, errors.New("err")
					},
				},
				order.ServiceMock{},
			},
			wantErr: true,
		},
		{
			name: "s.Store.OrderStore.GetAllOrdersByItemID fail",
			fields: fields{
				item.ServiceMock{
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return &models.Item{}, nil
					},
				},
				order.ServiceMock{
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return []*models.Order{
							&models.Order{},
						}, errors.New("err")
					},
				},
			},
			wantErr: true,
		},
		{
			name: "s.Store.OrderStore.DeleteOrder fail",
			fields: fields{
				item.ServiceMock{
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return &models.Item{}, nil
					},
				},
				order.ServiceMock{
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return []*models.Order{
							&models.Order{},
						}, nil
					},
					DeleteOrderFunc: func(tx *sql.Tx, o *models.Order) error {
						return errors.New("err")
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Pass",
			fields: fields{
				item.ServiceMock{
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return &models.Item{}, nil
					},
					DeleteFunc: func(tx *sql.Tx, i *models.Item) error {
						return nil
					},
				},
				order.ServiceMock{
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return []*models.Order{
							&models.Order{},
						}, nil
					},
					DeleteOrderFunc: func(tx *sql.Tx, o *models.Order) error {
						return nil
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: Store{
					ItemStore:  &tt.fields.Item,
					OrderStore: &tt.fields.Order,
				},
			}
			if err := s.DeleteItem(tt.args.tx, tt.args.itemID); (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
