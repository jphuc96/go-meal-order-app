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
	"git.d.foundation/datcom/backend/src/store/pic"
	"git.d.foundation/datcom/backend/src/store/user"
)

func TestService_AddPIC(t *testing.T) {
	type fields struct {
		PIC pic.ServiceMock
	}
	type args struct {
		tx *sql.Tx
		p  *domain.PICInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.PeopleInCharge
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "when AddPIC() success",
			fields: fields{
				pic.ServiceMock{
					ExistFunc: func(tx *sql.Tx, o *domain.PICInput) (bool, error) {
						return false, nil
					},
					AddFunc: func(tx *sql.Tx, o *domain.PICInput) (*models.PeopleInCharge, error) {
						return nil, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "when AddPIC() unsuccess",
			fields: fields{
				pic.ServiceMock{
					ExistFunc: func(tx *sql.Tx, o *domain.PICInput) (bool, error) {
						return true, nil
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
					PICStore: &tt.fields.PIC,
				},
			}
			got, err := s.AddPIC(tt.args.tx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddPIC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddPIC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetPICByMenuID(t *testing.T) {
	type fields struct {
		PIC pic.ServiceMock
	}
	type args struct {
		tx     *sql.Tx
		menuID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.PeopleInCharge
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "when GetPIC() success",
			fields: fields{
				pic.ServiceMock{
					GetByMenuIDFunc: func(tx *sql.Tx, menuID int) ([]*models.PeopleInCharge, error) {
						return nil, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "when GetPIC() failed",
			fields: fields{
				pic.ServiceMock{
					GetByMenuIDFunc: func(tx *sql.Tx, menuID int) ([]*models.PeopleInCharge, error) {
						return nil, errors.New("Failed to get")
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
					PICStore: &tt.fields.PIC,
				},
			}
			got, err := s.GetPICByMenuID(tt.args.tx, tt.args.menuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetPICByMenuID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetPICByMenuID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetUserbyPIC(t *testing.T) {
	type fields struct {
		User user.ServiceMock
		PIC  pic.ServiceMock
	}
	type args struct {
		tx  *sql.Tx
		pic *models.PeopleInCharge
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "when GetPIC() failed",
			fields: fields{
				user.ServiceMock{
					GetByIDFunc: func(tx *sql.Tx, userID int) (*models.User, error) {
						return &models.User{}, nil
					},
				},
				pic.ServiceMock{},
			},
			args: args{
				tx:  &sql.Tx{},
				pic: &models.PeopleInCharge{},
			},
			want:    &models.User{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: Store{
					UserStore: &tt.fields.User,
				},
			}
			got, err := s.GetUserbyPIC(tt.args.tx, tt.args.pic)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetUserbyPIC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetUserbyPIC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetAllOrderUserOfMenu(t *testing.T) {
	type fields struct {
		Item  item.ServiceMock
		Order order.ServiceMock
		User  user.ServiceMock
	}
	type args struct {
		tx     *sql.Tx
		menuID int
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
			name: "GetAllItemsByMenuID fail",
			fields: fields{
				item.ServiceMock{
					GetAllItemsByMenuIDFunc: func(tx *sql.Tx, menuID int) ([]*models.Item, error) {
						return nil, errors.New("err")
					},
				},
				order.ServiceMock{},
				user.ServiceMock{},
			},
			wantErr: true,
		},
		{
			name: "GetAllOrdersByItemID fail",
			fields: fields{
				item.ServiceMock{
					GetAllItemsByMenuIDFunc: func(tx *sql.Tx, menuID int) ([]*models.Item, error) {
						return []*models.Item{
							&models.Item{},
						}, nil
					},
				},
				order.ServiceMock{
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return []*models.Order{
							&models.Order{},
						}, errors.New("err")
					},
				},
				user.ServiceMock{},
			},
			wantErr: true,
		},
		{
			name: "s.Store.UserStore.GetByID fail",
			fields: fields{
				item.ServiceMock{
					GetAllItemsByMenuIDFunc: func(tx *sql.Tx, menuID int) ([]*models.Item, error) {
						return []*models.Item{
							&models.Item{},
						}, nil
					},
				},
				order.ServiceMock{
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return []*models.Order{
							&models.Order{},
						}, nil
					},
				},
				user.ServiceMock{
					GetByIDFunc: func(tx *sql.Tx, userID int) (*models.User, error) {
						return &models.User{}, errors.New("err")
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Pass",
			fields: fields{
				item.ServiceMock{
					GetAllItemsByMenuIDFunc: func(tx *sql.Tx, menuID int) ([]*models.Item, error) {
						return []*models.Item{
							&models.Item{},
						}, nil
					},
				},
				order.ServiceMock{
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return []*models.Order{
							&models.Order{},
						}, nil
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: Store{
					ItemStore:  &tt.fields.Item,
					OrderStore: &tt.fields.Order,
					UserStore:  &tt.fields.User,
				},
			}
			got, err := s.GetAllOrderUserOfMenu(tt.args.tx, tt.args.menuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetAllOrderUserOfMenu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetAllOrderUserOfMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GeneratePIC(t *testing.T) {
	type fields struct {
		Item  item.ServiceMock
		Order order.ServiceMock
		User  user.ServiceMock
	}
	type args struct {
		tx     *sql.Tx
		menuID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.PICUser
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "s.Store.ItemStore.GetAllItemsByMenuID fail",
			fields: fields{
				item.ServiceMock{
					GetAllItemsByMenuIDFunc: func(tx *sql.Tx, menuID int) ([]*models.Item, error) {
						return []*models.Item{
							&models.Item{},
						}, errors.New("err")
					},
				},
				order.ServiceMock{},
				user.ServiceMock{},
			},
			wantErr: true,
		},
		{
			name: "s.Store.OrderStore.GetAllOrdersByItemID fail",
			fields: fields{
				item.ServiceMock{
					GetAllItemsByMenuIDFunc: func(tx *sql.Tx, menuID int) ([]*models.Item, error) {
						return []*models.Item{
							&models.Item{},
						}, nil
					},
				},
				order.ServiceMock{
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return nil, errors.New("err")
					},
				},
				user.ServiceMock{},
			},
			wantErr: true,
		},
		{
			name: "s.GetAllOrderUserOfMenu fail",
			fields: fields{
				item.ServiceMock{
					GetAllItemsByMenuIDFunc: func(tx *sql.Tx, menuID int) ([]*models.Item, error) {
						return []*models.Item{
							&models.Item{},
						}, nil
					},
				},
				order.ServiceMock{
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return []*models.Order{
							&models.Order{},
						}, nil
					},
				},
				user.ServiceMock{
					GetByIDFunc: func(tx *sql.Tx, userID int) (*models.User, error) {
						return &models.User{}, errors.New("err")
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Pass",
			fields: fields{
				item.ServiceMock{
					GetAllItemsByMenuIDFunc: func(tx *sql.Tx, menuID int) ([]*models.Item, error) {
						return []*models.Item{
							&models.Item{},
						}, nil
					},
				},
				order.ServiceMock{
					GetAllOrdersByItemIDFunc: func(tx *sql.Tx, ItemID int) ([]*models.Order, error) {
						return []*models.Order{
							&models.Order{
								UserID: 0,
							},
							&models.Order{
								UserID: 1,
							},
						}, nil
					},
				},
				user.ServiceMock{
					GetByIDFunc: func(tx *sql.Tx, userID int) (*models.User, error) {
						return &models.User{}, nil
					},
				},
			},
			want: []domain.PICUser{
				domain.PICUser{},
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
					UserStore:  &tt.fields.User,
				},
			}
			got, err := s.GeneratePIC(tt.args.tx, tt.args.menuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GeneratePIC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GeneratePIC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DeleteAllPIC(t *testing.T) {
	type fields struct {
		PIC pic.ServiceMock
	}
	type args struct {
		tx     *sql.Tx
		menuID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Pass",
			fields: fields{
				pic.ServiceMock{
					DeleteAllPICFunc: func(tx *sql.Tx, menuID int) error {
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
					PICStore: &tt.fields.PIC,
				},
			}
			if err := s.DeleteAllPIC(tt.args.tx, tt.args.menuID); (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteAllPIC() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
