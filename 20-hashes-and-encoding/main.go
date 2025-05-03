package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

// Simple helper to compute and print hash
func printHash(name string, hasher io.Writer, data []byte, sumFunc func() []byte) {
	hasher.Write(data)
	fmt.Printf("    %s: %x\n", name, sumFunc())
}

func hashExamples() {
	fmt.Println("1. Hash Functions")
	data := []byte("Hello, World!")

	printHash("SHA256", sha256.New(), data, func() []byte {
		h := sha256.New()
		h.Write(data)
		return h.Sum(nil)
	})

	printHash("SHA512", sha512.New(), data, func() []byte {
		h := sha512.New()
		h.Write(data)
		return h.Sum(nil)
	})

	printHash("SHA1 (deprecated)", sha1.New(), data, func() []byte {
		h := sha1.New()
		h.Write(data)
		return h.Sum(nil)
	})

	printHash("MD5 (deprecated)", md5.New(), data, func() []byte {
		h := md5.New()
		h.Write(data)
		return h.Sum(nil)
	})

	// Streamed hash example
	streamed := sha256.New()
	io.WriteString(streamed, "Hello, ")
	io.WriteString(streamed, "World!")
	fmt.Printf("    Streamed SHA256: %x\n", streamed.Sum(nil))
}

func hmacExamples() {
	fmt.Println("\n2. HMAC Examples")
	key := []byte("secret-key")
	message := []byte("Hello, World!")

	// HMAC with SHA256
	h := hmac.New(sha256.New, key)
	h.Write(message)
	hmac256 := h.Sum(nil)
	fmt.Printf("    HMAC-SHA256: %x\n", hmac256)

	// HMAC with SHA512
	h2 := hmac.New(sha512.New, key)
	h2.Write(message)
	fmt.Printf("    HMAC-SHA512: %x\n", h2.Sum(nil))

	// HMAC verification
	verify := func(msg, key, expected []byte) bool {
		mac := hmac.New(sha256.New, key)
		mac.Write(msg)
		return hmac.Equal(mac.Sum(nil), expected)
	}

	fmt.Printf("    Verification: %v\n", verify(message, key, hmac256))
	fmt.Printf("    Verification (
