package service

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"git.d.foundation/datcom/backend/src/store/item"
	"git.d.foundation/datcom/backend/src/store/order"
)

func TestService_DeleteItem(t *testing.T) {
	type fields struct {
		Item  item.ServiceMock
		Order order.ServiceMock
	}
	type args struct {
		tx *sql.Tx
		i  *domain.Item
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
			name: "PassNoOrder",
			fields: fields{
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return true, nil
					},
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return &models.Item{
							ID:       7,
							ItemName: "Mon 7",
							MenuID:   22,
						}, nil
					},
					DeleteFunc: func(tx *sql.Tx, i *models.Item) error {
						return nil
					},
				},
				order.ServiceMock{
					CheckOrderExistByItemIDFunc: func(tx *sql.Tx, ItemID int) (bool, error) {
						return false, nil
					},
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return nil, nil
					},
					DeleteOrderFunc: func(tx *sql.Tx, o *models.Order) error {
						return nil
					},
				},
			},
			args: args{
				i: &domain.Item{
					ItemName: "Mon 7",
					MenuID:   22,
				},
			},
			want: &models.Item{
				ID:       7,
				ItemName: "Mon 7",
				MenuID:   22,
			},
			wantErr: false,
		},
		{
			name: "PassWithOrders",
			fields: fields{
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return true, nil
					},
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return &models.Item{
							ID:       7,
							ItemName: "Mon 7",
							MenuID:   22,
						}, nil
					},
					DeleteFunc: func(tx *sql.Tx, i *models.Item) error {
						return nil
					},
				},
				order.ServiceMock{
					CheckOrderExistByItemIDFunc: func(tx *sql.Tx, ItemID int) (bool, error) {
						return true, nil
					},
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return []*models.Order{
							&models.Order{
								ItemID: 7,
								UserID: 1,
							},
						}, nil
					},
					DeleteOrderFunc: func(tx *sql.Tx, o *models.Order) error {
						return nil
					},
				},
			},
			args: args{
				i: &domain.Item{
					ItemName: "Mon 7",
					MenuID:   22,
				},
			},
			want: &models.Item{
				ID:       7,
				ItemName: "Mon 7",
				MenuID:   22,
			},
			wantErr: false,
		},
		{
			name: "NoItem",
			fields: fields{
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return false, nil
					},
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return nil, errors.New("Item does not exist")
					},
					DeleteFunc: func(tx *sql.Tx, i *models.Item) error {
						return nil
					},
				},
				order.ServiceMock{
					CheckOrderExistByItemIDFunc: func(tx *sql.Tx, ItemID int) (bool, error) {
						return false, nil
					},
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return nil, nil
					},
					DeleteOrderFunc: func(tx *sql.Tx, o *models.Order) error {
						return nil
					},
				},
			},
			args: args{
				i: &domain.Item{
					ItemName: "Mon 7",
					MenuID:   22,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "CheckItemError",
			fields: fields{
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return false, errors.New("Check Item Error")
					},
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return nil, nil
					},
					DeleteFunc: func(tx *sql.Tx, i *models.Item) error {
						return nil
					},
				},
				order.ServiceMock{
					CheckOrderExistByItemIDFunc: func(tx *sql.Tx, ItemID int) (bool, error) {
						return false, nil
					},
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return nil, nil
					},
					DeleteOrderFunc: func(tx *sql.Tx, o *models.Order) error {
						return nil
					},
				},
			},
			args: args{
				i: &domain.Item{
					ItemName: "Mon 7",
					MenuID:   22,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Find Item Error",
			fields: fields{
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return true, nil
					},
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return nil, errors.New("Find Item Error")
					},
					DeleteFunc: func(tx *sql.Tx, i *models.Item) error {
						return nil
					},
				},
				order.ServiceMock{
					CheckOrderExistByItemIDFunc: func(tx *sql.Tx, ItemID int) (bool, error) {
						return false, nil
					},
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return nil, nil
					},
					DeleteOrderFunc: func(tx *sql.Tx, o *models.Order) error {
						return nil
					},
				},
			},
			args: args{
				i: &domain.Item{
					ItemName: "Mon 7",
					MenuID:   22,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Delete Item Error",
			fields: fields{
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return true, nil
					},
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return nil, nil
					},
					DeleteFunc: func(tx *sql.Tx, i *models.Item) error {
						return errors.New("Delete Item Error")
					},
				},
				order.ServiceMock{
					CheckOrderExistByItemIDFunc: func(tx *sql.Tx, ItemID int) (bool, error) {
						return false, nil
					},
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return nil, nil
					},
					DeleteOrderFunc: func(tx *sql.Tx, o *models.Order) error {
						return nil
					},
				},
			},
			args: args{
				i: &domain.Item{
					ItemName: "Mon 7",
					MenuID:   22,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Check Order Error",
			fields: fields{
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return true, nil
					},
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return nil, nil
					},
					DeleteFunc: func(tx *sql.Tx, i *models.Item) error {
						return nil
					},
				},
				order.ServiceMock{
					CheckOrderExistByItemIDFunc: func(tx *sql.Tx, ItemID int) (bool, error) {
						return false, errors.New("Check Order Error")
					},
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return nil, nil
					},
					DeleteOrderFunc: func(tx *sql.Tx, o *models.Order) error {
						return nil
					},
				},
			},
			args: args{
				i: &domain.Item{
					ItemName: "Mon 7",
					MenuID:   22,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Get Orders Error",
			fields: fields{
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return true, nil
					},
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return nil, nil
					},
					DeleteFunc: func(tx *sql.Tx, i *models.Item) error {
						return nil
					},
				},
				order.ServiceMock{
					CheckOrderExistByItemIDFunc: func(tx *sql.Tx, ItemID int) (bool, error) {
						return true, nil
					},
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return nil, errors.New("Get Orders Error")
					},
					DeleteOrderFunc: func(tx *sql.Tx, o *models.Order) error {
						return nil
					},
				},
			},
			args: args{
				i: &domain.Item{
					ItemName: "Mon 7",
					MenuID:   22,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Delete Order Error",
			fields: fields{
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return true, nil
					},
					FindByIDFunc: func(tx *sql.Tx, itemID int) (*models.Item, error) {
						return nil, nil
					},
					DeleteFunc: func(tx *sql.Tx, i *models.Item) error {
						return nil
					},
				},
				order.ServiceMock{
					CheckOrderExistByItemIDFunc: func(tx *sql.Tx, ItemID int) (bool, error) {
						return true, nil
					},
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return []*models.Order{
							&models.Order{
								ItemID: 7,
								UserID: 1,
							},
						}, nil
					},
					DeleteOrderFunc: func(tx *sql.Tx, o *models.Order) error {
						return errors.New("Delete Order Error")
					},
				},
			},
			args: args{
				i: &domain.Item{
					ItemName: "Mon 7",
					MenuID:   22,
				},
			},
			want:    nil,
			wantErr: true,
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
			got, err := s.DeleteItem(tt.args.tx, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.DeleteItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
