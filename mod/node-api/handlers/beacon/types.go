// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package beacon

import (
	types "github.com/berachain/beacon-kit/mod/node-api/types/beacon"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
)

// Backend is the interface for backend of the beacon API.
type Backend[ForkT, ValidatorT any] interface {
	GenesisBackend
	BlockBackend
	ValidatorBackend[ValidatorT]
	StateRootAtSlot(
		slot uint64,
	) (common.Root, error)
	StateForkAtSlot(
		slot uint64,
	) (ForkT, error)
}

type GenesisBackend interface {
	GenesisValidatorsRoot(
		slot uint64,
	) (common.Root, error)
}

type BlockBackend interface {
	BlockRootAtSlot(
		slot uint64,
	) (common.Root, error)
	BlockRewardsAtSlot(
		slot uint64,
	) (*types.BlockRewardsData, error)
}

type ValidatorBackend[ValidatorT any] interface {
	ValidatorByID(
		slot uint64,
		id string,
	) (*types.ValidatorData[ValidatorT], error)
	ValidatorsByIDs(
		slot uint64,
		ids []string,
		statuses []string,
	) ([]*types.ValidatorData[ValidatorT], error)
	ValidatorBalancesByIDs(
		slot uint64,
		ids []string,
	) ([]*types.ValidatorBalanceData, error)
}