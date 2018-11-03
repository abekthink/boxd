// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package script

import (
	"testing"

	"github.com/facebookgo/ensure"
)

var (
	tokenName   = "box"
	tokenSupply = uint64(3000000000000)
)

func TestIssueToken(t *testing.T) {
	params := &IssueParams{Name: tokenName, TotalSupply: tokenSupply}
	script := IssueTokenScript(testPubKeyHash, params)

	params2, err := script.GetIssueParams()
	ensure.Nil(t, err)
	ensure.DeepEqual(t, params2, params)

	ensure.DeepEqual(t, script.GetTokenAmount(), tokenSupply)
}

func TestTransferToken(t *testing.T) {
	params := &TransferParams{Amount: tokenSupply}
	script := TransferTokenScript(testPubKeyHash, params)

	params2, err := script.GetTransferParams()
	ensure.Nil(t, err)
	ensure.DeepEqual(t, params2, params)

	ensure.DeepEqual(t, script.GetTokenAmount(), tokenSupply)
}
