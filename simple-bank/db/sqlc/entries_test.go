package db

import (
	"context"
	"github.com/DamThiLanAnh/simple-bank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testStore.CreateEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := creatRandomAccount(t)
	createRandomEntry(t, account)
}

//func TestGetEntry(t *testing.T) {
//	account := creatRandomAccount(t)
//	account2, err := testStore.GetEntries(context.Background(), account1.ID)
//}
//
//func TestListEntry(t *testing.T) {
//
//}
