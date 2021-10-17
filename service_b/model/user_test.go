package model

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name    string
		want    *User
		wantErr bool
	}{
		{
			name: "正常系",
			want: &User{
				ID: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
