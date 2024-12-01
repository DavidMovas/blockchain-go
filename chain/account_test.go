package chain

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	keyStoreDir   = ".keystorenode"
	blockStoreDir = ".keystorenode"
	chainName     = "testblockchain"
	authPass      = "password"
	ownerPass     = "password"
	ownerBalance  = 1000
)

func TestAccountWriteReadSignTxVerifyTx(t *testing.T) {
	defer os.RemoveAll(keyStoreDir)
	// Create a new account
	acc, err := NewAccount()
	if err != nil {
		t.Fatal(err)
	}
	// Persist the new account
	err = acc.Write(keyStoreDir, []byte(ownerPass))
	if err != nil {
		t.Fatal(err)
	}
	// Re-create the persisted account
	path := filepath.Join(keyStoreDir, string(acc.Address()))
	acc, err = ReadAccount(path, []byte(ownerPass))
	if err != nil {
		t.Fatal(err)
	}
	// Create and sign a transaction
	tx := NewTx(acc.Address(), Address("to"), 12, 1)
	stx, err := acc.SignTx(tx)
	if err != nil {
		t.Fatal(err)
	}
	// Verify that the signature of the signed transaction is valid
	valid, err := VerifyTx(stx)
	if err != nil {
		t.Fatal(err)
	}
	if !valid {
		t.Errorf("invalid transaction signature")
	}
}
