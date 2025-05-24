package db

import (
	"context"
	"testing"

	"github.com/rodrigostrj/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.NotZero(t, account.ID)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

}

func TestGetAccount(t *testing.T) {

}

func TestUpdateAccount(t *testing.T) {

}

func TestCreateManyAccounts(t *testing.T) {

}

func TestGetManyAccounts(t *testing.T) {

}

func TestDeleteAccount(t *testing.T) {

}
