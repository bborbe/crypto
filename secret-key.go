// Copyright (c) 2023 Seibert Group GmbH All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto

import (
	"context"
	"os"
	"strings"

	"github.com/bborbe/errors"
)

// https://dev.to/breda/secret-key-encryption-with-go-using-aes-316d

func SecretKeyFromFile(ctx context.Context, path string) (SecretKey, error) {
	if strings.TrimSpace(path) == "" {
		return nil, errors.New(ctx, "secret key path is empty")
	}
	bytes, err := os.ReadFile(path) // #nosec G304 -- path from configuration
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "read file %s failed", path)
	}
	return bytes, nil

}

// SecretKey should be the AES secretKey, either 16 or 32 bytes
// to select AES-128 or AES-256.
type SecretKey []byte

func (s SecretKey) Bytes() []byte {
	return s
}

func (s SecretKey) Ptr() *SecretKey {
	return &s
}
