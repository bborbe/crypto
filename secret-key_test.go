// Copyright (c) 2023 Seibert Group GmbH All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto_test

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	crypto "github.com/bborbe/crypto"
)

var _ = Describe("SecretKey", func() {
	var ctx context.Context
	var err error
	BeforeEach(func() {
		ctx = context.Background()
	})
	Context("SecretKeyFromFile", func() {
		var secretKey crypto.SecretKey
		var path string
		JustBeforeEach(func() {
			secretKey, err = crypto.SecretKeyFromFile(ctx, path)
		})
		Context("empty path", func() {
			BeforeEach(func() {
				path = ""
			})
			It("returns error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("secret key path is empty"))
			})
			It("returns nil secretKey", func() {
				Expect(secretKey).To(BeNil())
			})
		})
		Context("whitespace only path", func() {
			BeforeEach(func() {
				path = "   "
			})
			It("returns error", func() {
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("secret key path is empty"))
			})
			It("returns nil secretKey", func() {
				Expect(secretKey).To(BeNil())
			})
		})
		Context("not found", func() {
			BeforeEach(func() {
				path = "/tmp/no-existing-path"
			})
			It("returns error", func() {
				Expect(err).NotTo(BeNil())
			})
			It("returns no secretKey", func() {
				Expect(secretKey).To(BeNil())
			})
		})
		Context("found", func() {
			BeforeEach(func() {
				file, err := os.CreateTemp("", "")
				Expect(err).To(BeNil())
				_, err = file.WriteString("hello world")
				Expect(err).To(BeNil())
				path = file.Name()
			})
			AfterEach(func() {
				Expect(os.Remove(path)).To(BeNil())
			})
			It("returns no error", func() {
				Expect(err).To(BeNil())
			})
			It("returns secretKey", func() {
				Expect(secretKey).To(Equal(crypto.SecretKey("hello world")))
			})
		})
	})
})
