package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/simplemaxr/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createdRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	randomAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	if err != nil {
		t.Fatal(err)
	}
	randomAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	if err != nil {
		t.Fatal(err)
	}

	arg := CreateTransferParams{
		FromAccountID: randomAccount1.ID,
		ToAccountID:   randomAccount2.ID,
		Amount:        util.RandomInt(0, randomAccount1.Balance),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	transfer := createdRandomTransfer(t)
	require.NotEmpty(t, transfer)
}

func TestGetTransfer(t *testing.T) {
	transfer := createdRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)
	require.Equal(t, transfer, transfer2)
}

func TestUpdateTransfer(t *testing.T) {
	transfer := createdRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:            transfer.ID,
		FromAccountID: transfer.FromAccountID,
		ToAccountID:   transfer.ToAccountID,
		Amount:        util.RandomInt(0, transfer.Amount),
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, arg.Amount, transfer2.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createdRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createdRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
}
