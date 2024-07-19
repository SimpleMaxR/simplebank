package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/simplemaxr/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createdRandomEntry(t *testing.T) Entry {
	account1 := createRandomAccount(t)
	randomAccount, err := testQueries.GetAccount(context.Background(), account1.ID)
	if err != nil {
		t.Fatal(err)
	}

	arg := CreateEntryParams{
		AccountID: randomAccount.ID,
		Amount:    util.RandomInt(0, randomAccount.Balance),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	return entry
}

func TestCreateEntry(t *testing.T) {
	entry := createdRandomEntry(t)
	require.NotEmpty(t, entry)
}

func TestGetEntry(t *testing.T) {
	entry := createdRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	require.Equal(t, entry, entry2)
}

func TestUpdateEntry(t *testing.T) {
	entry := createdRandomEntry(t)

	arg := UpdateEntryParams{
		ID:        entry.ID,
		AccountID: entry.AccountID,
		Amount:    util.RandomInt(0, entry.Amount),
	}
	entry2, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, entry.AccountID, entry2.AccountID)
	require.Equal(t, arg.Amount, entry2.Amount)
}

func TestDeleteEntry(t *testing.T) {
	entry := createdRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createdRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
