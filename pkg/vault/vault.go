package vault

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

// 2MB
const (
	badgerIndexCacheSize = 2 << 20
	badgerKeyMinBytes    = 32

	badgerKeyEntryPrefix = "entry:"
)

var (
	ErrMasterKeyTooShort = fmt.Errorf("master key must be at least %d bytes", badgerKeyMinBytes)
	ErrKeyNotFound       = errors.New("key not found")
	ErrInvalidValue      = errors.New("invalid value")
)

type Vault struct {
	badgerDB *badger.DB
}

func NewVault(path, masterKey string) (*Vault, error) {
	badgerKey := []byte(masterKey)
	if len(badgerKey) < badgerKeyMinBytes {
		return nil, ErrMasterKeyTooShort
	}

	badgerOpt := badger.DefaultOptions(path).
		WithEncryptionKey(badgerKey[:32]).
		WithIndexCacheSize(badgerIndexCacheSize)

	badgerDB, err := badger.Open(badgerOpt)
	if err != nil {
		return nil, fmt.Errorf("badger: failed to open %s: %w", path, err)
	}

	return &Vault{
		badgerDB: badgerDB,
	}, nil
}

func (v *Vault) SaveEntry(data *EntryData) error {
	// Ignore empty data
	if data == nil {
		return nil
	}

	if err := v.badgerDB.Update(func(txn *badger.Txn) error {
		dataBytes, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("json: failed to marshal: %w", err)
		}

		if err := txn.Set(v.getKeyEntry(data.ID), dataBytes); err != nil {
			return fmt.Errorf("badger: failed to set: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("badger: failed to update: %w", err)
	}

	return nil
}

func (v *Vault) GetEntry(id string) (*EntryData, error) {
	data := &EntryData{}

	if err := v.badgerDB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(v.getKeyEntry(id))
		if err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				return errors.Join(ErrKeyNotFound, err)
			}

			return fmt.Errorf("badger: failed to get: %w", err)
		}

		if err := item.Value(func(val []byte) error {
			if err := json.Unmarshal(val, data); err != nil {
				return fmt.Errorf("json: failed to unmarshal: %w", err)
			}

			// Feel like a hack
			// I use type to detect otp data
			// Side effect is json.Unmarshal 2 times
			switch data.EntryType {
			case EntryTypeTOTP:
				data.OTPData = &TOTPData{}
			default:
				return ErrInvalidValue
			}

			if err := json.Unmarshal(val, data); err != nil {
				return fmt.Errorf("json: failed to unmarshal: %w", err)
			}

			return nil
		}); err != nil {
			return fmt.Errorf("badger: failed to get value: %w", err)
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("badger: failed to view: %w", err)
	}

	return data, nil
}

func (v *Vault) DeleteEntry(id string) error {
	if err := v.badgerDB.Update(func(txn *badger.Txn) error {
		if err := txn.Delete(v.getKeyEntry(id)); err != nil {
			if errors.Is(err, badger.ErrKeyNotFound) {
				return errors.Join(ErrKeyNotFound, err)
			}

			return fmt.Errorf("badger: failed to delete: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("badger: failed to update: %w", err)
	}

	return nil
}

func (v *Vault) getKeyEntry(id string) []byte {
	return []byte(badgerKeyEntryPrefix + id)
}
