package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkUsername(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "given a correct username no error is thrown",
			args: args{
				username: "username",
			},
			wantErr: assert.NoError,
		},
		{
			name: "given an incorrect username because of its lenght an error is thrown",
			args: args{
				username: "usernameusernameusernameusernameusernameusernameusernameusernameusername",
			},
			wantErr: assert.Error,
		},
		{
			name: "given an incorrect username because of its characters an error is thrown",
			args: args{
				username: "u s e r n a m e",
			},
			wantErr: assert.Error,
		},
		{
			name: "given an incorrect username because of its characters an error is thrown",
			args: args{
				username: "--username",
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := checkUsername(tt.args.username)
			tt.wantErr(t, err)
		})
	}

}
