package vcoclient

type RowsModifiedConfirmation struct {
	// An error message explaining why the method failed
	Error_ string `json:"error,omitempty"`
	// The number of rows modified
	Rows int32 `json:"rows"`
}
