/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package persistence

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLabels(t *testing.T) {
	tests := []struct {
		label   string
		success bool
	}{
		{label: "", success: false},
		{label: ".", success: false},
		{label: "0", success: true},
		{label: ":", success: false},
		{label: "_", success: false},
		{label: "a", success: true},
		{label: "a#", success: false},
		{label: "a$", success: false},
		{label: "a%", success: false},
		{label: "a-", success: true},
		{label: "a++b", success: true},
		{label: "a+b", success: true},
		{label: "a+bb", success: true},
		{label: "a--b", success: true},
		{label: "a-b", success: true},
		{label: "a-bb", success: true},
		{label: "a::b", success: false},
		{label: "a:b", success: false},
		{label: "a__b", success: true},
		{label: "a_b", success: true},
		{label: "a_bb", success: true},
		{label: "aa", success: true},
	}

	for _, tt := range tests {
		t.Run(tt.label, func(t *testing.T) {
			err := validateLabel(tt.label)
			if tt.success {
				assert.NoError(t, err)
			} else {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.label)
			}
		})
	}
}