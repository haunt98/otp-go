package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"hash"
	"time"
)

// https://datatracker.ietf.org/doc/html/rfc4226
func HOTP(key []byte, counter uint64, digits int) (int, error) {
	// RFC shows HOTP use SHA1
	return hotpByCustomHash(sha1.New, key, counter, digits)
}

// HOTP with custom hash algorithm
// counter: C
// key: K
// digits: 6 to 8
// Copy code, idea from https://github.com/rsc/2fa
func hotpByCustomHash(customHash func() hash.Hash, key []byte, counter uint64, digits int) (int, error) {
	// Step 1: Generate an hmac value from input algorithm
	h := hmac.New(customHash, key)

	if err := binary.Write(h, binary.BigEndian, counter); err != nil {
		return 0, fmt.Errorf("binary: failed to write counter: %w", err)
	}

	hmacResult := h.Sum(nil)

	// Step 2: Generate a 4-byte string
	// hmac_result is 20 bytes
	// offset is last, low-order 4 bits of hmac_result[19]
	// 0xF == 1111 -> & 0xF aka & 1111 -> only keep last 4 bits
	// 0 <= offset <= 15
	offset := hmacResult[len(hmacResult)-1] & 0xF
	// Combine hmax_result[offset]...hmac_result[offset+3] to P
	// Only get last 31 bits of P
	// 0x7F == 01111111 -> hide first bit to remove confusion about signed, unsigned
	// 0xFF == 11111111 -> only keep last 8 bits
	binCode := binary.BigEndian.Uint32([]byte{
		hmacResult[offset] & 0x7F,
		hmacResult[offset+1] & 0xFF,
		hmacResult[offset+2] & 0xFF,
		hmacResult[offset+3] & 0xFF,
	})

	// Step 3: Compute an HOTP value
	// d10 is 10^digits
	// 8 is digits limit
	d10 := uint32(1)
	for i := 0; i < digits && i < 8; i++ {
		d10 *= 10
	}

	return int(binCode % d10), nil
}

// https://datatracker.ietf.org/doc/html/rfc6238
// Extension of HOTP
// timeStep: X
// Assume T0 = 0
func TOTP(customHash func() hash.Hash, key []byte, t time.Time, timeStep, digits int) (int, error) {
	// T = (Current Unix time - T0) / X
	counter := uint64(t.Unix()) / uint64(timeStep)

	return hotpByCustomHash(customHash, key, counter, digits)
}
