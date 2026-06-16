// Copyright (c) 2023 Seibert Group GmbH All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto_test

import (
	"context"
	"encoding/hex"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	crypto "github.com/bborbe/crypto"
)

var _ = Describe("Crypter", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	It("invalid length", func() {
		var err error
		// openssl rand -hex 16
		key, err := hex.DecodeString("b1c1fe36eeb532b37e58")
		Expect(err).To(BeNil())
		crypter := crypto.NewCrypter(key)
		_, err = crypter.Encrypt(ctx, []byte("hello world"))
		Expect(err).NotTo(BeNil())
	})
	It("16 bit", func() {
		var err error
		// openssl rand -hex 16
		key, err := hex.DecodeString("26a739e39162cbe6c55942d4f12893f4")
		Expect(err).To(BeNil())
		Expect(key).To(HaveLen(16))
		crypter := crypto.NewCrypter(key)
		encrypted, err := crypter.Encrypt(ctx, []byte("hello world"))
		Expect(err).To(BeNil())
		plain, err := crypter.Decrypt(ctx, encrypted)
		Expect(err).To(BeNil())
		Expect(plain).To(Equal([]byte("hello world")))
	})
	It("24 bit", func() {
		var err error
		// openssl rand -hex 16
		key, err := hex.DecodeString("0c8ed413276982d5515edb3e2eea720c06647a67acc3a6a3")
		Expect(err).To(BeNil())
		Expect(key).To(HaveLen(24))
		crypter := crypto.NewCrypter(key)
		encrypted, err := crypter.Encrypt(ctx, []byte("hello world"))
		Expect(err).To(BeNil())
		plain, err := crypter.Decrypt(ctx, encrypted)
		Expect(err).To(BeNil())
		Expect(plain).To(Equal([]byte("hello world")))
	})
	It("32 bit", func() {
		var err error
		// openssl rand -hex 32
		key, err := hex.DecodeString(
			"7630d82b45d30339fd687fd9e607700eb53fd4fe33af85f2349531f7f78cb67d",
		)
		Expect(err).To(BeNil())
		Expect(key).To(HaveLen(32))
		crypter := crypto.NewCrypter(key)
		encrypted, err := crypter.Encrypt(ctx, []byte("hello world"))
		Expect(err).To(BeNil())
		plain, err := crypter.Decrypt(ctx, encrypted)
		Expect(err).To(BeNil())
		Expect(plain).To(Equal([]byte("hello world")))
	})
})
