// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package runtime

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
	"github.com/stretchr/testify/require"
)

func TestSASCredentialPolicy(t *testing.T) {
	const key = "foo"
	cred := exported.NewSASCredential(key)

	const headerName = "fake-auth"
	policy := NewSASCredentialPolicy(cred, headerName, nil)
	require.NotNil(t, policy)

	pl := exported.NewPipeline(shared.TransportFunc(func(req *http.Request) (*http.Response, error) {
		require.EqualValues(t, key, req.Header.Get(headerName))
		return &http.Response{}, nil
	}), policy)

	req, err := NewRequest(context.Background(), http.MethodGet, "http://contoso.com")
	require.NoError(t, err)

	_, err = pl.Do(req)
	require.NoError(t, err)
}
