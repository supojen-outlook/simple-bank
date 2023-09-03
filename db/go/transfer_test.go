package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {

	// arrange
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        100,
	}

	// act
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	// assert
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, int64(100), transfer.Amount)
	require.Equal(t, account1.ID, transfer.FromAccountID)
	require.Equal(t, account2.ID, transfer.ToAccountID)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

}
