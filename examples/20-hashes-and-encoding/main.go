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

// hashExamples demonstrates various cryptographic hash functions
func hashExamples() {
	fmt.Println("    Hash Functions:")
	data := []byte("Hello, World!")

	// SHA256
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	fmt.Printf("        SHA256: %x\n", sha256Hash.Sum(nil))

	// SHA512
	sha512Hash := sha512.New()
	sha512Hash.Write(data)
	fmt.Printf("        SHA512: %x\n", sha512Hash.Sum(nil))

	// SHA1 (Note: SHA1 is considered cryptographically broken)
	sha1Hash := sha1.New()
	sha1Hash.Write(data)
	fmt.Printf("        SHA1: %x\n", sha1Hash.Sum(nil))

	// MD5 (Note: MD5 is considered cryptographically broken)
	md5Hash := md5.New()
	md5Hash.Write(data)
	fmt.Printf("        MD5: %x\n", md5Hash.Sum(nil))

	// Hash streaming example
	streamHash := sha256.New()
	io.WriteString(streamHash, "Hello, ")
	io.WriteString(streamHash, "World!")
	fmt.Printf("        Streamed SHA256: %x\n", streamHash.Sum(nil))
}

// hmacExamples demonstrates HMAC (Hash-based Message Authentication Code)
func hmacExamples() {
	fmt.Println("    HMAC Examples:")

	key := []byte("secret-key")
	message := []byte("Hello, World!")

	// HMAC-SHA256
	hmacSha256 := hmac.New(sha256.New, key)
	hmacSha256.Write(message)
	fmt.Printf("        HMAC-SHA256: %x\n", hmacSha256.Sum(nil))

	// HMAC-SHA512
	hmacSha512 := hmac.New(sha512.New, key)
	hmacSha512.Write(message)
	fmt.Printf("        HMAC-SHA512: %x\n", hmacSha512.Sum(nil))

	// Verify HMAC
	expectedHMAC := hmacSha256.Sum(nil)
	verifyHMAC := func(message, key []byte, expectedMAC []byte) bool {
		mac := hmac.New(sha256.New, key)
		mac.Write(message)
		messageMAC := mac.Sum(nil)
		return hmac.Equal(messageMAC, expectedMAC)
	}

	fmt.Printf("        HMAC verification: %v\n", verifyHMAC(message, key, expectedHMAC))
	fmt.Printf("        HMAC verification (tampered): %v\n", verifyHMAC([]byte("Tampered message"), key, expectedHMAC))
}

// hashingPatterns demonstrates common hashing patterns
func hashingPatterns() {
	fmt.Println("    Hashing Patterns:")

	// Hash multiple items
	hash := sha256.New()
	items := []string{"item1", "item2", "item3"}
	for _, item := range items {
		io.WriteString(hash, item)
	}
	fmt.Printf("        Combined hash: %x\n", hash.Sum(nil))

	// Hash with salt
	password := []byte("my-password")
	salt := []byte("random-salt")

	hash.Reset()
	hash.Write(salt)
	hash.Write(password)
	fmt.Printf("        Salted hash: %x\n", hash.Sum(nil))

	// File hashing simulation
	simulateFileHash := func(chunks ...[]byte) []byte {
		h := sha256.New()
		for _, chunk := range chunks {
			h.Write(chunk)
		}
		return h.Sum(nil)
	}

	fileChunks := [][]byte{
		[]byte("chunk1"),
		[]byte("chunk2"),
		[]byte("chunk3"),
	}
	fmt.Printf("        File hash: %x\n", simulateFileHash(fileChunks...))
}

// base64Examples demonstrates Base64 encoding and decoding
func base64Examples() {
	fmt.Println("    Base64 Encoding:")

	data := []byte("Hello, World!")

	// Standard Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Printf("        Standard Base64: %s\n", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Printf("        Error decoding Base64: %v\n", err)
	} else {
		fmt.Printf("        Decoded Base64: %s\n", decoded)
	}

	// URL-safe Base64
	urlEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Printf("        URL-safe Base64: %s\n", urlEncoded)

	// Raw Base64 (no padding)
	rawEncoded := base64.RawStdEncoding.EncodeToString(data)
	fmt.Printf("        Raw Base64 (no padding): %s\n", rawEncoded)

	// Base64 with custom padding
	customEncoder := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/").
		WithPadding('*')
	customEncoded := customEncoder.EncodeToString(data)
	fmt.Printf("        Custom Base64: %s\n", customEncoded)
}

// base32Examples demonstrates Base32 encoding and decoding
func base32Examples() {
	fmt.Println("    Base32 Encoding:")

	data := []byte("Hello, World!")

	// Standard Base32
	encoded := base32.StdEncoding.EncodeToString(data)
	fmt.Printf("        Standard Base32: %s\n", encoded)

	decoded, err := base32.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Printf("        Error decoding Base32: %v\n", err)
	} else {
		fmt.Printf("        Decoded Base32: %s\n", decoded)
	}

	// Hex Base32
	hexEncoded := base32.HexEncoding.EncodeToString(data)
	fmt.Printf("        Hex Base32: %s\n", hexEncoded)

	// Base32 with custom padding
	customEncoder := base32.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567").
		WithPadding('*')
	customEncoded := customEncoder.EncodeToString(data)
	fmt.Printf("        Custom Base32: %s\n", customEncoded)
}

// hexExamples demonstrates hexadecimal encoding and decoding
func hexExamples() {
	fmt.Println("    Hex Encoding:")

	data := []byte("Hello, World!")

	// Standard hex encoding
	hexEncoded := hex.EncodeToString(data)
	fmt.Printf("        Hex encoded: %s\n", hexEncoded)

	hexDecoded, err := hex.DecodeString(hexEncoded)
	if err != nil {
		fmt.Printf("        Error decoding hex: %v\n", err)
	} else {
		fmt.Printf("        Hex decoded: %s\n", hexDecoded)
	}

	// Hex dump
	fmt.Println("        Hex dump:")
	fmt.Printf("%s", hex.Dump(data))
}

func main() {
	fmt.Println("=== Go Cryptographic Hashes and Encoding Examples ===")

	fmt.Println("\n1. Hash Functions:")
	hashExamples()

	fmt.Println("\n2. HMAC Examples:")
	hmacExamples()

	fmt.Println("\n3. Hashing Patterns:")
	hashingPatterns()

	fmt.Println("\n4. Base64 Encoding:")
	base64Examples()

	fmt.Println("\n5. Base32 Encoding:")
	base32Examples()

	fmt.Println("\n6. Hex Encoding:")
	hexExamples()
}
