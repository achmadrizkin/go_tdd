package test

import (
	"testing"

	logic "go_tdd/logic"

	"github.com/stretchr/testify/assert"
)

func TestSayHelloFAIL(t *testing.T) {
	assert.NotEqual(t, "William", logic.SayHello("William"))
}

func TestSayHelloSUCCESS(t *testing.T) {
	assert.Equal(t, "", logic.SayHello("William"))
}

func TestSayHelloWelcomeSUCCESS(t *testing.T) {
	assert.Equal(t, "Hello William. Welcome!", logic.SayHelloWelcome("William"))
}

func TestSayHelloWelcomeFAIL(t *testing.T) {
	assert.NotEqual(t, "Hello William. Welcome!", logic.SayHelloWelcome("ACENG"))
}
