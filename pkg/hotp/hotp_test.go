package hotp

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	rfcKeySHA1   = []byte("12345678901234567890")
	rfcKeySHA256 = []byte("12345678901234567890123456789012")
	rfcKeySHA512 = []byte("1234567890123456789012345678901234567890123456789012345678901234")
)

func TestHOTP(t *testing.T) {
	tests := []struct {
		name        string
		key         []byte
		counters    []uint64
		digits      int
		wantResults []int
	}{
		{
			name:     "rfc",
			key:      rfcKeySHA1,
			counters: []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			digits:   6,
			wantResults: []int{
				755224,
				287082,
				359152,
				969429,
				338314,
				254676,
				287922,
				162583,
				399871,
				520489,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for i := range tc.counters {
				gotResult, gotErr := HOTP(tc.key, tc.counters[i], tc.digits)
				assert.NoError(t, gotErr)
				assert.Equal(t, tc.wantResults[i], gotResult)
			}
		})
	}
}

func TestTOTP(t *testing.T) {
	tests := []struct {
		name        string
		hashAlgo    func() hash.Hash
		key         []byte
		timesInUnix []uint64
		digits      int
		timeStep    int
		wantResults []int
	}{
		{
			name:     "rfc sha1",
			hashAlgo: sha1.New,
			key:      rfcKeySHA1,
			timesInUnix: []uint64{
				59,
				1111111109,
				1111111111,
				1234567890,
				2000000000,
				20000000000,
			},
			digits:   8,
			timeStep: 30,
			wantResults: []int{
				94287082,
				7081804,
				14050471,
				89005924,
				69279037,
				65353130,
			},
		},
		{
			name:     "rfc sha256",
			hashAlgo: sha256.New,
			key:      rfcKeySHA256,
			timesInUnix: []uint64{
				59,
				1111111109,
				1111111111,
				1234567890,
				2000000000,
				20000000000,
			},
			digits:   8,
			timeStep: 30,
			wantResults: []int{
				46119246,
				68084774,
				67062674,
				91819424,
				90698825,
				77737706,
			},
		},
		{
			name:     "rfc sha512",
			hashAlgo: sha512.New,
			key:      rfcKeySHA512,
			timesInUnix: []uint64{
				59,
				1111111109,
				1111111111,
				1234567890,
				2000000000,
				20000000000,
			},
			digits:   8,
			timeStep: 30,
			wantResults: []int{
				90693936,
				25091201,
				99943326,
				93441116,
				38618901,
				47863826,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for i := range tc.timesInUnix {
				gotResult, gotErr := TOTP(tc.hashAlgo, tc.key, time.Unix(int64(tc.timesInUnix[i]), 0), tc.timeStep, tc.digits)
				assert.NoError(t, gotErr)
				assert.Equal(t, tc.wantResults[i], gotResult)
			}
		})
	}
}
