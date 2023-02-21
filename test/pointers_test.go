package test

import (
	logic "go_tdd/logic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWalletSUCCESS(t *testing.T) {
	wallet := logic.Wallet{}
	wallet.Deposit(logic.Bitcoin(10))

	got := wallet.Balance()
	want := logic.Bitcoin(10)

	assert.Equal(t, want, got)
}

func TestWalletFAIL(t *testing.T) {
	wallet := logic.Wallet{}
	wallet.Deposit(logic.Bitcoin(10))

	got := wallet.Balance()
	want := logic.Bitcoin(20)

	assert.NotEqual(t, want, got)
}
