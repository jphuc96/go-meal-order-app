package service

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"git.d.foundation/datcom/backend/src/store/order"
	"git.d.foundation/datcom/backend/src/store/user"
)

func TestService_AddOrder(t *testing.T) {
	type fields struct {
		Order order.ServiceMock
	}
	type args struct {
		o *domain.OrderInput
		t *sql.Tx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test AddOrder() when order does not exist",
			fields: fields{
				order.ServiceMock{
					ExistFunc: func(tx *sql.Tx, o *domain.OrderInput) (bool, error) {
						return false, nil
					},
					AddFunc: func(tx *sql.Tx, o *domain.OrderInput) (*models.Order, error) {
						return nil, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test AddOrder() when order exists",
			fields: fields{
				order.ServiceMock{
					ExistFunc: func(tx *sql.Tx, o *domain.OrderInput) (bool, error) {
						return true, nil
					},
					GetOrderByOrderInputFunc: func(tx *sql.Tx, o *domain.OrderInput) (*models.Order, error) {
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
					OrderStore: &tt.fields.Order,
				},
			}
			if _, err := s.AddOrder(tt.args.t, tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("Service.AddOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_DeleteOrder(t *testing.T) {
	type fields struct {
		Order order.ServiceMock
	}
	type args struct {
		o *domain.OrderInput
		t *sql.Tx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test DeleteOrder() when order exists",
			fields: fields{
				order.ServiceMock{
					ExistFunc: func(tx *sql.Tx, o *domain.OrderInput) (bool, error) {
						return true, nil
					},
					DeleteFunc: func(tx *sql.Tx, o *domain.OrderInput) error {
						return nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test DeleteOrder() when order does not exist",
			fields: fields{
				order.ServiceMock{
					ExistFunc: func(tx *sql.Tx, o *domain.OrderInput) (bool, error) {
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
					OrderStore: &tt.fields.Order,
				},
			}
			if err := s.DeleteOrder(tt.args.t, tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetOrdersByMenuAndUser(t *testing.T) {
	type fields struct {
		Order order.ServiceMock
	}
	type args struct {
		menuID string
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Item
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test GetOrder() success",
			fields: fields{
				order.ServiceMock{
					GetFunc: func(menuID string, userID string) ([]*domain.Item, error) {
						return nil, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Test GetOrder() fail",
			fields: fields{
				order.ServiceMock{
					GetFunc: func(menuID string, userID string) ([]*domain.Item, error) {
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
					OrderStore: &tt.fields.Order,
				},
			}
			got, err := s.GetOrdersByMenuAndUser(tt.args.menuID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetOrdersByMenuAndUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetOrdersByMenuAndUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetOrderUsersByItem(t *testing.T) {
	type fields struct {
		Order order.ServiceMock
		User  user.ServiceMock
	}
	type args struct {
		tx     *sql.Tx
		itemID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Pass",
			fields: fields{
				order.ServiceMock{
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						m := []*models.Order{
							&models.Order{},
						}
						return m, nil
					},
				},
				user.ServiceMock{
					GetByIDFunc: func(tx *sql.Tx, userID int) (*models.User, error) {
						return &models.User{}, nil
					},
				},
			},
			want: []*models.User{
				&models.User{},
			},
			wantErr: false,
		},
		{
			name: "Fail",
			fields: fields{
				order.ServiceMock{
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						m := []*models.Order{
							&models.Order{},
						}
						return m, nil
					},
				},
				user.ServiceMock{
					GetByIDFunc: func(tx *sql.Tx, userID int) (*models.User, error) {
						return &models.User{}, errors.New("error")
					},
				},
			},

			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: Store{
					OrderStore: &tt.fields.Order,
					UserStore:  &tt.fields.User,
				},
			}
			got, err := s.GetOrderUsersByItem(tt.args.tx, tt.args.itemID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetOrderUsersByItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetOrderUsersByItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
