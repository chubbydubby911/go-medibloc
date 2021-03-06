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

package net

import (
	"math/rand"
)

// ChainSyncPeersFilter will filter some peers randomly
type ChainSyncPeersFilter struct {
	excludedPIDs map[string]struct{}
}

//SetExcludedPIDs set excludedPIDs
func (filter *ChainSyncPeersFilter) SetExcludedPIDs(excludedPIDs map[string]struct{}) {
	filter.excludedPIDs = excludedPIDs
}

// Filter implements PeerFilterAlgorithm interface
func (filter *ChainSyncPeersFilter) Filter(peers PeersSlice) PeersSlice {
	if filter.excludedPIDs != nil {
		filteredPeersSlice := make(PeersSlice, 0)
		for _, p := range peers {
			if _, ok := filter.excludedPIDs[p.(*Stream).PID()]; ok {
				continue
			}
			filteredPeersSlice = append(filteredPeersSlice, p)
		}
		return filteredPeersSlice
	}
	return peers
}

// RandomPeerFilter will filter a peer randomly
type RandomPeerFilter struct {
}

// Filter implements PeerFilterAlgorithm interface
func (filter *RandomPeerFilter) Filter(peers PeersSlice) PeersSlice {
	if len(peers) == 0 {
		return peers
	}

	selection := rand.Intn(len(peers))
	return peers[selection : selection+1]
}
