package tests

import (
	"fmt"
	"github.com/haoyayun/libsignal-protocol-go/util/keyhelper"
	"testing"
)

func TestRegistrationID(t *testing.T) {
	regID := keyhelper.GenerateRegistrationID()
	fmt.Println(regID)
}
