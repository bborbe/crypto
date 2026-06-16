// Copyright (c) 2023 Seibert Group GmbH All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"github.com/bborbe/errors"
	"github.com/golang/glog"
)

//counterfeiter:generate -o mocks/crypter.go --fake-name Crypter . Crypter
type Crypter interface {
	Encrypt(ctx context.Context, value []byte) ([]byte, error)
	Decrypt(ctx context.Context, value []byte) ([]byte, error)
}

func NewCrypter(secretKey SecretKey) Crypter {
	return &crypter{
		secretKey: secretKey,
	}
}

type crypter struct {
	secretKey SecretKey
}

func (c *crypter) Encrypt(ctx context.Context, value []byte) ([]byte, error) {
	glog.V(3).Infof("encrypt with %d bit secretKey started", len(c.secretKey))
	gcm, err := c.createGcm(ctx)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "create gcm failed")
	}

	// We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
	// A nonce should always be randomly generated for every encryption.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, errors.Wrapf(ctx, err, "nonce failed")
	}

	// ciphertext here is actually nonce+ciphertext
	// So that when we decrypt, just knowing the nonce size
	// is enough to separate it from the ciphertext.
	result := gcm.Seal(nonce, nonce, value, nil)

	glog.V(3).Infof("encrypt with %d bit secretKey completed", len(c.secretKey))
	return result, nil
}

func (c *crypter) Decrypt(ctx context.Context, value []byte) ([]byte, error) {
	glog.V(3).Infof("decrypt with %d bit secretKey started", len(c.secretKey))
	gcm, err := c.createGcm(ctx)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "create gcm failed")
	}

	// Since we know the ciphertext is actually nonce+ciphertext
	// And len(nonce) == NonceSize(). We can separate the two.
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := value[:nonceSize], value[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "gcm open failed")
	}

	glog.V(3).Infof("decrypt with %d bit secretKey completed", len(c.secretKey))
	return plaintext, nil
}

func (c *crypter) createGcm(ctx context.Context) (cipher.AEAD, error) {
	aesCipher, err := aes.NewCipher(c.secretKey.Bytes())
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "create cipher failed")
	}
	gcm, err := cipher.NewGCM(aesCipher)
	if err != nil {
		return nil, errors.Wrapf(ctx, err, "create gcm failed")
	}
	return gcm, nil
}
