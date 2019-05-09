package service

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"git.d.foundation/datcom/backend/src/store/item"
	"git.d.foundation/datcom/backend/src/store/menu"
)

func TestService_AddItems(t *testing.T) {
	type fields struct {
		Menu menu.ServiceMock
		Item item.ServiceMock
	}
	type args struct {
		tx     *sql.Tx
		items  *domain.ItemInput
		menuID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Item
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Pass",
			fields: fields{
				menu.ServiceMock{
					CheckMenuExistFunc: func(tx *sql.Tx, menuID int) (bool, error) {
						return true, nil
					},
				},
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return false, nil
					},
					AddFunc: func(tx *sql.Tx, i *domain.Item) (*models.Item, error) {
						return &models.Item{
							ItemName: "Mon 7",
							MenuID:   22,
						}, nil
					},
				},
			},
			args: args{
				items: &domain.ItemInput{
					Items: []domain.Item{
						{
							ItemName: "Mon 7",
						},
					},
				},
				menuID: 22,
			},
			want: []*models.Item{
				&models.Item{
					ItemName: "Mon 7",
					MenuID:   22,
				},
			},
			wantErr: false,
		},
		{
			name: "CheckMenuExistFailed",
			fields: fields{
				menu.ServiceMock{
					CheckMenuExistFunc: func(tx *sql.Tx, menuID int) (bool, error) {
						return false, errors.New("Find Error")
					},
				},
				item.ServiceMock{},
			},
			args: args{
				menuID: 22,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "MenuNotExist",
			fields: fields{
				menu.ServiceMock{
					CheckMenuExistFunc: func(tx *sql.Tx, menuID int) (bool, error) {
						return false, nil
					},
				},
				item.ServiceMock{},
			},
			args: args{
				menuID: 22,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "CheckItemExistFailed",
			fields: fields{
				menu.ServiceMock{
					CheckMenuExistFunc: func(tx *sql.Tx, menuID int) (bool, error) {
						return true, nil
					},
				},
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return false, errors.New("Check Error")
					},
				},
			},
			args: args{
				items: &domain.ItemInput{
					Items: []domain.Item{
						{
							ItemName: "Mon 7",
						},
					},
				},
				menuID: 22,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ItemExists",
			fields: fields{
				menu.ServiceMock{
					CheckMenuExistFunc: func(tx *sql.Tx, menuID int) (bool, error) {
						return true, nil
					},
				},
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return true, nil
					},
				},
			},
			args: args{
				items: &domain.ItemInput{
					Items: []domain.Item{
						{
							ItemName: "Mon 7",
						},
					},
				},
				menuID: 22,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "AddItemFailed",
			fields: fields{
				menu.ServiceMock{
					CheckMenuExistFunc: func(tx *sql.Tx, menuID int) (bool, error) {
						return true, nil
					},
				},
				item.ServiceMock{
					CheckItemExistFunc: func(tx *sql.Tx, itemID int) (bool, error) {
						return false, nil
					},
					AddFunc: func(tx *sql.Tx, i *domain.Item) (*models.Item, error) {
						return nil, errors.New("Add Item Failed")
					},
				},
			},
			args: args{
				items: &domain.ItemInput{
					Items: []domain.Item{
						{
							ItemName: "Mon 7",
						},
					},
				},
				menuID: 22,
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store{
					MenuStore: &tt.fields.Menu,
					ItemStore: &tt.fields.Item,
				},
			}
			got, err := s.AddItems(tt.args.tx, tt.args.items, tt.args.menuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetAllItemsByMenuID(t *testing.T) {
	type fields struct {
		Item item.ServiceMock
	}
	type args struct {
		tx     *sql.Tx
		menuID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*models.Item
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Passed",
			fields: fields{
				item.ServiceMock{
					GetAllItemsByMenuIDFunc: func(tx *sql.Tx, menuID int) ([]*models.Item, error) {
						return []*models.Item{}, nil
					},
				},
			},
			want:    []*models.Item{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store{
					ItemStore: &tt.fields.Item,
				},
			}
			got, err := s.GetAllItemsByMenuID(tt.args.tx, tt.args.menuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetAllItemsByMenuID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetAllItemsByMenuID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_AddItemToMenu(t *testing.T) {
	type fields struct {
		Item item.ServiceMock
		Menu menu.ServiceMock
	}
	type args struct {
		tx       *sql.Tx
		itemName string
		menuID   int
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
			name: "MenuNotExist",
			fields: fields{
				item.ServiceMock{},
				menu.ServiceMock{
					CheckMenuExistFunc: func(tx *sql.Tx, menuID int) (bool, error) {
						return false, nil
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Error",
			fields: fields{
				item.ServiceMock{},
				menu.ServiceMock{
					CheckMenuExistFunc: func(tx *sql.Tx, menuID int) (bool, error) {
						return false, errors.New("error")
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Passed",
			fields: fields{
				item.ServiceMock{
					AddFunc: func(tx *sql.Tx, i *domain.Item) (*models.Item, error) {
						return &models.Item{}, nil
					},
				},
				menu.ServiceMock{
					CheckMenuExistFunc: func(tx *sql.Tx, menuID int) (bool, error) {
						return true, nil
					},
				},
			},
			want:    &models.Item{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store{
					ItemStore: &tt.fields.Item,
					MenuStore: &tt.fields.Menu,
				},
			}
			got, err := s.AddItemToMenu(tt.args.tx, tt.args.itemName, tt.args.menuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.AddItemToMenu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.AddItemToMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}
