package test

import (
	"grpc/tests/suite"
	"testing"

	ssov1 "github.com/DenisEMPS/test/gen/go/sso"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	ctx, st := suite.New(t)

	testTable := []struct {
		name        string
		email       string
		password    string
		expectedErr string
	}{
		{
			name:        "Empty password field",
			email:       gofakeit.Email(),
			password:    "",
			expectedErr: "empty password field",
		},
		{
			name:        "Empty email field",
			email:       "",
			password:    randomFakePassword(),
			expectedErr: "empty email field",
		},
		{
			name:        "",
			email:       "",
			password:    randomFakePassword(),
			expectedErr: "empty email field",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
				Email:    testCase.email,
				Password: testCase.password,
			})

			require.Error(t, err)
			require.Contains(t, err.Error(), testCase.expectedErr)
		})
	}
}
