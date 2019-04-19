package service

import (
	"errors"
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
		p *domain.MenuInput
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
					FindByNameFunc: func(menuName string) (*models.Menu, error) {
						return nil, nil
					},
					CreateFunc: func(p *domain.MenuInput) (*models.Menu, error) {
						return nil, nil
					},
				},
			},
			args: args{
				&domain.MenuInput{OwnerID: 0, MenuName: "0", Deadline: time.Now(), PaymentReminder: time.Now(), Status: 1},
			},
			want:    &models.Menu{OwnerID: 0, MenuName: "0", CreatedAt: time.Now(), Deadline: time.Now(), PaymentReminder: time.Now(), Status: 1},
			wantErr: false,
		},
		{
			name: "Failed at finding menu",
			fields: fields{
				menu.ServiceMock{
					FindByNameFunc: func(menuName string) (*models.Menu, error) {
						return nil, errors.New("Menu exists")
					},
				},
			},
			args: args{
				&domain.MenuInput{OwnerID: 0, MenuName: "0", Deadline: time.Now(), PaymentReminder: time.Now(), Status: 1},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Failed at creating menu",
			fields: fields{menu.ServiceMock{
				FindByNameFunc: func(menuName string) (*models.Menu, error) {
					return &models.Menu{OwnerID: 0, MenuName: "0", CreatedAt: time.Now(), Deadline: time.Now(), PaymentReminder: time.Now(), Status: 1}, nil
				},
				CreateFunc: func(p *domain.MenuInput) (*models.Menu, error) {
					return nil, errors.New("Cannot insert menu")
				},
			},
			},
			args: args{
				&domain.MenuInput{OwnerID: 0, MenuName: "0", Deadline: time.Now(), PaymentReminder: time.Now(), Status: 1},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Menu: &tt.fields.Menu,
			}
			_, err := s.CreateMenu(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateMenu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
