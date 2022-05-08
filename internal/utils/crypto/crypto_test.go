package crypto_test

import (
	"election-service/internal/utils/crypto"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestCrypto_IsHashPasswordMatch(t *testing.T) {
	t.Run("Given an raw password", func(t *testing.T) {
		password := "000000"
		t.Run("When hashed is generated from this password", func(t *testing.T) {
			hash := "$2a$10$2t7bXEQo16CyseOkNE2CoeBz191J/fBwZw.msNXQWJmIb7fIn4Ewu"
			result := crypto.IsHashPasswordMatched(password, hash)
			assert.Equal(t, true, result)
		})
	})
}

func TestCrypto_HashPassword(t *testing.T) {
	t.Run("Given an raw password", func(t *testing.T) {
		password := "000000"
		t.Run("When hash without error", func(t *testing.T) {
			var hash string
			err := crypto.HashPassword(password, &hash)
			assert.Equal(t, true, err == nil)
		})
	})
}
