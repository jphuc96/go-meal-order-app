package service

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"git.d.foundation/datcom/backend/models"
	"git.d.foundation/datcom/backend/src/domain"
	"git.d.foundation/datcom/backend/src/store/pic"
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
