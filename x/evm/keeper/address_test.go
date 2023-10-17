package keeper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetGetAddressMapping(t *testing.T) {
	k, _, ctx := MockEVMKeeper()
	seiAddr, evmAddr := MockAddressPair()
	foundEVM, ok := k.GetEVMAddress(ctx, seiAddr)
	require.False(t, ok)
	foundSei, ok := k.GetSeiAddress(ctx, evmAddr)
	require.False(t, ok)
	k.SetAddressMapping(ctx, seiAddr, evmAddr)
	foundEVM, ok = k.GetEVMAddress(ctx, seiAddr)
	require.True(t, ok)
	require.Equal(t, evmAddr, foundEVM)
	foundSei, ok = k.GetSeiAddress(ctx, evmAddr)
	require.True(t, ok)
	require.Equal(t, seiAddr, foundSei)
}

func TestDeleteAddressMapping(t *testing.T) {
	k, _, ctx := MockEVMKeeper()
	seiAddr, evmAddr := MockAddressPair()
	k.SetAddressMapping(ctx, seiAddr, evmAddr)
	foundEVM, ok := k.GetEVMAddress(ctx, seiAddr)
	require.True(t, ok)
	require.Equal(t, evmAddr, foundEVM)
	foundSei, ok := k.GetSeiAddress(ctx, evmAddr)
	require.True(t, ok)
	require.Equal(t, seiAddr, foundSei)
	k.DeleteAddressMapping(ctx, seiAddr, evmAddr)
	foundEVM, ok = k.GetEVMAddress(ctx, seiAddr)
	require.False(t, ok)
	foundSei, ok = k.GetSeiAddress(ctx, evmAddr)
	require.False(t, ok)
}