package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

// https://datatracker.ietf.org/doc/html/rfc4226
// Copy code, idea from https://github.com/rsc/2fa
// counter: C
// key: K
// digits: 6 to 8
func HOTP(key []byte, counter uint64, digits int) (int, error) {
	// Step 1: Generate an HMAC-SHA-1 value, HMAC-SHA-1(K,C)
	h := hmac.New(sha1.New, key)

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
