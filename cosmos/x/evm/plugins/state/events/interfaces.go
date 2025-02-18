// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
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

package events

import (
	coretypes "github.com/berachain/polaris/eth/core/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	// PrecompileLogFactory is used to build an Ethereum log from a Cosmos event.
	PrecompileLogFactory interface {
		// Build builds an Ethereum log from a Cosmos event.
		Build(*sdk.Event) (*coretypes.Log, error)
	}

	// LogsDB defines the required function to add a log to the StateDB. This ensures a
	// precompile runner can only add logs to the StateDB and not modify any other state on the
	// StateDB.
	LogsDB interface {
		// AddLog adds a log to the database.
		AddLog(*coretypes.Log)
	}
)
