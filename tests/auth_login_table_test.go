package test

import (
	"grpc/tests/suite"
	"testing"

	ssov1 "github.com/DenisEMPS/test/gen/go/sso"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	ctx, st := suite.New(t)

	testTable := []struct {
		name        string
		email       string
		password    string
		appID       int32
		expectedErr string
	}{

		{
			name:        "Empty password",
			email:       gofakeit.Email(),
			password:    "",
			appID:       appID,
			expectedErr: "empty password field",
		},
		{
			name:        "Empty email",
			email:       "",
			password:    randomFakePassword(),
			appID:       appID,
			expectedErr: "empty email field",
		},
		{
			name:        "Empty both",
			email:       "",
			password:    "",
			appID:       appID,
			expectedErr: "empty email field",
		},
		{
			name:        "Login with Non-Matching Password",
			email:       gofakeit.Email(),
			password:    randomFakePassword(),
			appID:       appID,
			expectedErr: "invalid arguments",
		},
		{
			name:        "Empty appID field",
			email:       gofakeit.Email(),
			password:    randomFakePassword(),
			appID:       emptyAppID,
			expectedErr: "invalid arguments",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := st.AuthClient.Register(ctx, &ssov1.RegisterRequest{
				Email:    gofakeit.Email(),
				Password: randomFakePassword(),
			})
			require.NoError(t, err)

			_, err = st.AuthClient.Login(ctx, &ssov1.LoginRequest{
				Email:    testCase.email,
				Password: testCase.password,
				AppId:    testCase.appID,
			})

			require.Error(t, err)
			require.Contains(t, err.Error(), testCase.expectedErr)
		})
	}

}
