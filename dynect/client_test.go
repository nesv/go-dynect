package dynect

import (
	"os"
	"testing"
)

var (
	DynCustomerName string
	DynUsername string
	DynPassword string
)

func init() {
	DynCustomerName = os.Getenv("DYNECT_CUSTOMER_NAME")
	DynUsername = os.Getenv("DYNECT_USER_NAME")
	DynPassword = os.Getenv("DYNECT_PASSWORD")
}

func TestSetup(t *testing.T) {
	if len(DynCustomerName) == 0 {
		t.Fatal("DYNECT_CUSTOMER_NAME not set")
	}

	if len(DynUsername) == 0 {
		t.Fatal("DYNECT_USER_NAME not set")
	}

	if len(DynPassword) == 0 {
		t.Fatal("DYNECT_PASSWORD not set")
	}
}

func TestLoginLogout(t *testing.T) {
	client := NewClient(DynCustomerName)
	client.Verbose(true)
	err := client.Login(DynUsername, DynPassword)
	if err != nil {
		t.Error(err)
	}

	err = client.Logout()
	if err != nil {
		t.Error(err)
	}
}
