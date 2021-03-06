// Copyright (C) 2018  MediBloc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>

package core_test

import (
	"testing"

	"github.com/medibloc/go-medibloc/consensus/dpos"
	"github.com/medibloc/go-medibloc/core"
	"github.com/medibloc/go-medibloc/core/pb"
	"github.com/medibloc/go-medibloc/storage"
	"github.com/medibloc/go-medibloc/util/testutil"
	"github.com/mitchellh/copystructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGenesisBlock(t *testing.T) {
	genesisBlock, dynasties, dist := testutil.NewTestGenesisBlock(t, 21)

	assert.True(t, core.CheckGenesisBlock(genesisBlock))
	txs := genesisBlock.Transactions()
	initialMessage := "Genesis block of MediBloc"
	defaultPayload := &core.DefaultPayload{
		Message: initialMessage,
	}
	payloadBuf, err := defaultPayload.ToBytes()
	assert.NoError(t, err)
	assert.Equalf(t, txs[0].Payload(), payloadBuf, "Initial tx payload should equal '%s'", initialMessage)

	t.Log(len(txs[0].Hash()))
	t.Log(len(txs[1].Hash()))

	for i := 0; i < len(dist); i++ {
		assert.True(t, dist[i].Addr.Equals(txs[1+2*i].To()))
		assert.Equal(t, "1000000000000000000", txs[1+2*i].Value().String())
	}

	dposState := genesisBlock.State().DposState()
	for _, dynasty := range dynasties {
		addr := dynasty.Addr

		isCandidate, err := dposState.IsCandidate(addr)
		require.NoError(t, err)
		assert.True(t, isCandidate)
		inDynasty, err := dposState.InDynasty(addr)
		require.NoError(t, err)
		assert.True(t, inDynasty)
	}

	accState := genesisBlock.State().AccState()
	for _, holder := range dist {
		addr := holder.Addr
		acc, err := accState.GetAccount(addr)
		assert.NoError(t, err)

		assert.Equal(t, "1000000000000000000", acc.Balance.String())
		assert.Equal(t, 1, len(acc.TxsToSlice()))
	}
}

func TestCheckGenesisBlock(t *testing.T) {
	conf, _, _ := testutil.NewTestGenesisConf(t, testutil.DynastySize)
	stor, err := storage.NewMemoryStorage()
	require.NoError(t, err)
	consensus := dpos.New(testutil.DynastySize)
	genesis, err := core.NewGenesisBlock(conf, consensus, stor)
	require.NoError(t, err)

	ok := core.CheckGenesisConf(genesis, conf)
	require.True(t, ok)

	modified := copystructure.Must(copystructure.Copy(conf)).(*corepb.Genesis)
	modified.Meta.ChainId = 9898
	require.False(t, core.CheckGenesisConf(genesis, modified))

	modified = copystructure.Must(copystructure.Copy(conf)).(*corepb.Genesis)
	modified.Meta.DynastySize = 22
	require.False(t, core.CheckGenesisConf(genesis, modified))

	modified = copystructure.Must(copystructure.Copy(conf)).(*corepb.Genesis)
	modified.Consensus.Dpos.Dynasty = modified.Consensus.Dpos.Dynasty[1:]
	require.False(t, core.CheckGenesisConf(genesis, modified))

	modified = copystructure.Must(copystructure.Copy(conf)).(*corepb.Genesis)
	modified.Consensus.Dpos.Dynasty[0] = "Wrong Address"
	require.False(t, core.CheckGenesisConf(genesis, modified))

	modified = copystructure.Must(copystructure.Copy(conf)).(*corepb.Genesis)
	modified.TokenDistribution = modified.TokenDistribution[1:]
	require.False(t, core.CheckGenesisConf(genesis, modified))

	modified = copystructure.Must(copystructure.Copy(conf)).(*corepb.Genesis)
	modified.TokenDistribution[2].Address = "Wrong Address"
	require.False(t, core.CheckGenesisConf(genesis, modified))

	modified = copystructure.Must(copystructure.Copy(conf)).(*corepb.Genesis)
	modified.TokenDistribution[3].Balance = "989898"
	require.False(t, core.CheckGenesisConf(genesis, modified))

	modified = copystructure.Must(copystructure.Copy(conf)).(*corepb.Genesis)
	modified.TokenDistribution[4].Balance = "Wrong Value"
	require.False(t, core.CheckGenesisConf(genesis, modified))
}
