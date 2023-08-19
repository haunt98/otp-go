package vault

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type VaultSuite struct {
	suite.Suite

	filename string

	vault *Vault
}

func (s *VaultSuite) SetupTest() {
	s.filename = "test.db"

	var gotErr error
	s.vault, gotErr = NewVault(s.filename, "1WCwWj5b5h7UdZ1D2mqjDMFjy0J5tsUG")
	s.NoError(gotErr)
}

func (s *VaultSuite) TearDownTest() {
	gotErr := os.RemoveAll(s.filename)
	s.NoError(gotErr)
}

func TestVaultSuite(t *testing.T) {
	suite.Run(t, new(VaultSuite))
}

func (s *VaultSuite) Test() {
	data := &EntryData{
		ID:        "id",
		Name:      "name",
		Issuer:    "issuer",
		EntryType: "entry_type",
	}

	gotErr := s.vault.SaveEntry(data)
	s.NoError(gotErr)

	gotData, gotErr := s.vault.GetEntry(data.ID)
	s.NoError(gotErr)
	s.Equal(data, gotData)
}
