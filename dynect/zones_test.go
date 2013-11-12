package dynect

import (
	"os"
	"testing"
)

var (
	DynCustomerName string
)

func init() {
	DynCustomerName = os.Getenv("DYNECT_CUSTOMER_NAME")
}

func TestSetup(t *testing.T) {
	if len(DynCustomerName) == 0 {
		t.Fatal("DYNECT_CUSTOMER_NAME not set")
	}
}
