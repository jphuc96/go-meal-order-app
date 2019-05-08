package service

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"
	"time"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"git.d.foundation/datcom/backend/src/store/menu"
)

func TestService_CreateMenu(t *testing.T) {
	type fields struct {
		Menu menu.ServiceMock
	}
	type args struct {
		tx *sql.Tx
		p  *domain.MenuInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Menu
		wantErr bool
	}{
		// Test cases
		{
			name: "Pass",
			fields: fields{
				menu.ServiceMock{
					IsMenuNameUniqueFunc: func(tx *sql.Tx, menuName string) (bool, error) {
						return false, nil
					},
					CreateFunc: func(tx *sql.Tx, p *domain.MenuInput) (*models.Menu, error) {
						return nil, nil
					},
				},
			},
			args: args{
				&sql.Tx{},
				&domain.MenuInput{OwnerID: 0, MenuName: "0", Deadline: time.Now(), PaymentReminder: time.Now(), Status: 1},
			},
			want:    &models.Menu{OwnerID: 0, MenuName: "0", CreatedAt: time.Now(), Deadline: time.Now(), PaymentReminder: time.Now(), Status: 1},
			wantErr: false,
		},
		{
			name: "Failed at finding menu",
			fields: fields{
				menu.ServiceMock{
					IsMenuNameUniqueFunc: func(tx *sql.Tx, menuName string) (bool, error) {
						return false, errors.New("Check exist error")
					},
				},
			},
			args: args{
				&sql.Tx{},
				&domain.MenuInput{OwnerID: 0, MenuName: "0", Deadline: time.Now(), PaymentReminder: time.Now(), Status: 1},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Duplicate",
			fields: fields{
				menu.ServiceMock{
					IsMenuNameUniqueFunc: func(tx *sql.Tx, menuName string) (bool, error) {
						return true, nil
					},
				},
			},
			args: args{
				&sql.Tx{},
				&domain.MenuInput{OwnerID: 0, MenuName: "0", Deadline: time.Now(), PaymentReminder: time.Now(), Status: 1},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed at creating menu",
			fields: fields{menu.ServiceMock{
				IsMenuNameUniqueFunc: func(tx *sql.Tx, menuName string) (bool, error) {
					return false, nil
				},
				CreateFunc: func(tx *sql.Tx, p *domain.MenuInput) (*models.Menu, error) {
					return nil, errors.New("Cannot insert menu")
				},
			},
			},
			args: args{
				&sql.Tx{},
				&domain.MenuInput{OwnerID: 0, MenuName: "0", Deadline: time.Now(), PaymentReminder: time.Now(), Status: 1},
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
				},
			}
			_, err := s.CreateMenu(tt.args.tx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateMenu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestService_GetLatestMenu(t *testing.T) {
	type fields struct {
		Menu menu.ServiceMock
	}
	type args struct {
		tx *sql.Tx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Menu
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "m != nil",
			fields: fields{
				menu.ServiceMock{
					GetLatestMenuFunc: func(tx *sql.Tx) (*models.Menu, error) {
						return &models.Menu{}, nil
					},
				},
			},
			wantErr: false,
		},
		{
			name: "m == nil",
			fields: fields{
				menu.ServiceMock{
					GetLatestMenuFunc: func(tx *sql.Tx) (*models.Menu, error) {
						return nil, nil
					},
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: Store{
					MenuStore: &tt.fields.Menu,
				},
			}
			got, err := s.GetLatestMenu(tt.args.tx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetLatestMenu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetLatestMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetMenuByID(t *testing.T) {
	type fields struct {
		Menu menu.ServiceMock
	}
	type args struct {
		tx     *sql.Tx
		menuID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Menu
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Pass",
			fields: fields{
				menu.ServiceMock{
					FindByIDFunc: func(tx *sql.Tx, menuID int) (*models.Menu, error) {
						return &models.Menu{}, nil
					},
				},
			},
			want:    &models.Menu{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: Store{
					MenuStore: &tt.fields.Menu,
				},
			}
			got, err := s.GetMenuByID(tt.args.tx, tt.args.menuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetMenuByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetMenuByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateMenu(t *testing.T) {
	type fields struct {
		Menu menu.ServiceMock
	}
	type args struct {
		tx         *sql.Tx
		updateMenu *models.Menu
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Menu
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: Store{
					MenuStore: &tt.fields.Menu,
				},
			}
			got, err := s.UpdateMenu(tt.args.tx, tt.args.updateMenu)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateMenu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.UpdateMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_HandleMenuDeadline(t *testing.T) {
	type fields struct {
		Store Store
	}
	type args struct {
		tx   *sql.Tx
		menu *models.Menu
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: tt.fields.Store,
			}
			if err := s.HandleMenuDeadline(tt.args.tx, tt.args.menu); (err != nil) != tt.wantErr {
				t.Errorf("Service.HandleMenuDeadline() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
