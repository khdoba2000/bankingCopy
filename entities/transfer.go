package entities

import (
	"errors"

	"github.com/khdoba/banking/constants"
	"github.com/khdoba/banking/pkg/utils"
)

// Transfer
type Transfer struct {
	TransactionImp
	AccountToID string
}

// Validate ...
func (t *Transfer) Validate() error {
	if !utils.IsValidUUID(t.AccountID) {
		return errors.New("invalid AccountID: invalid uuid")
	}

	if !utils.IsValidUUID(t.AccountToID) {
		return errors.New("invalid AccountID: invalid uuid")
	}

	if t.AccountID == t.AccountToID {
		return errors.New("cannot transfer to the same account")
	}
	return nil
}

// GetTypeID ...
func (t *Transfer) GetTypeID() int {
	return constants.TransferTransactionID
}

// IsOut ...
func (t *Transfer) IsOut() bool {
	return true
}
