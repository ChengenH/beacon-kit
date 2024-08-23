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

package log

import (
	sdklog "cosmossdk.io/log"
	"github.com/berachain/beacon-kit/mod/log"
)

type SDKLogger[LoggerT log.AdvancedLogger[LoggerT]] struct {
	Logger LoggerT
}

func WrapSDKLogger[LoggerT log.AdvancedLogger[LoggerT]](
	logger LoggerT,
) *SDKLogger[LoggerT] {
	return &SDKLogger[LoggerT]{
		Logger: logger,
	}
}

func (l *SDKLogger[LoggerT]) Info(msg string, keyVals ...any) {
	l.Logger.Info(msg, keyVals...)
}

func (l *SDKLogger[LoggerT]) Warn(msg string, keyVals ...any) {
	l.Logger.Warn(msg, keyVals...)
}

func (l *SDKLogger[LoggerT]) Error(msg string, keyVals ...any) {
	l.Logger.Error(msg, keyVals...)
}

func (l *SDKLogger[LoggerT]) Debug(msg string, keyVals ...any) {
	l.Logger.Debug(msg, keyVals...)
}

func (l *SDKLogger[LoggerT]) With(keyVals ...any) sdklog.Logger {
	return &SDKLogger[LoggerT]{Logger: l.Logger.With(keyVals...)}
}

func (l *SDKLogger[LoggerT]) Impl() any {
	return l.Logger
}