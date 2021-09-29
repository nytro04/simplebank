package db

import (
	"context"
	"testing"
	"time"

	"github.com/nytro04/simplebank/util"
	"github.com/stretchr/testify/require"
)


func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams {
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.Code)
	require.NotZero(t, entry.CreatedAt)

	return entry

}
func TestGetEntry(t *testing.T)  {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)

	entry2, err := testQueries.GetEntry(context.Background(), int32(entry1.Code))
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.Code, entry2.Code)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestCreateEntry(t *testing.T)  {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestListEntries(t *testing.T)  {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams {
		AccountID: account.ID,
		Limit: 5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}

}