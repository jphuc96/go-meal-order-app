package service

import (
	"net/http"
	"reflect"
	"testing"

	"git.d.foundation/datcom/backend/src/domain"
	"git.d.foundation/datcom/backend/src/store/user"
)

func TestService_FortressVerify(t *testing.T) {
	type fields struct {
		Store Store
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.FTResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: tt.fields.Store,
			}
			got, err := s.FortressVerify(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.FortressVerify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.FortressVerify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_AuthCheck(t *testing.T) {
	type fields struct {
		Header *http.Header
		User   user.ServiceMock
	}
	type args struct {
		r *http.Request
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
				Store: Store{
					UserStore: &tt.fields.User,
				},
			}
			if err := s.AuthCheck(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Service.AuthCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetGoogleUserInfo(t *testing.T) {
	type fields struct {
		Store Store
	}
	type args struct {
		idToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.GoogleUser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Store: tt.fields.Store,
			}
			got, err := s.GetGoogleUserInfo(tt.args.idToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetGoogleUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetGoogleUserInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
