package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/supojen-outlook/simple-bank/util"
)

func TestCreateAccount(t *testing.T) {

	// arrange
	arg := CreateAccountParams{
		Owner:    util.RandomName(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	// act
	account, err := testQueries.CreateAccount(context.Background(), arg)

	// assert
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	// teardown
	deleteRandomAccount(t, account.ID)
}

func TestGetAccount(t *testing.T) {

	// arrange
	account1 := createRandomAccount(t)

	// act
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	// assert
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.ID, account2.ID)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

	// teardown
	deleteRandomAccount(t, account1.ID)
}

func TestUpdateAccount(t *testing.T) {

	// arrange
	account := createRandomAccount(t)
	arg := UpdtaeAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	// act
	updatedAccount, err := testQueries.UpdtaeAccount(context.Background(), arg)

	// assert
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, account.Currency, updatedAccount.Currency)
	require.Equal(t, account.ID, updatedAccount.ID)
	require.WithinDuration(t, account.CreatedAt, updatedAccount.CreatedAt, time.Second)

	// teardown
	deleteRandomAccount(t, account.ID)
}

func TestDeleteAccount(t *testing.T) {

	// arrange
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	// act
	accountInDB, err := testQueries.GetAccount(context.Background(), account.ID)

	// assert
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountInDB)
}

func TestListAccount(t *testing.T) {

	// arrange
	for i := 0; i < 5; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountsParams{
		ID:    0,
		Limit: 5,
	}

	// act
	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	// assert
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

	// teardown
	for _, account := range accounts {
		deleteRandomAccount(t, account.ID)
	}
}

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomName(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	return account
}

func deleteRandomAccount(t *testing.T, id int64) {
	err := testQueries.DeleteAccount(context.TODO(), id)
	require.NoError(t, err)
}
