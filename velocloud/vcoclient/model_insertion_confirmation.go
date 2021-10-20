package vcoclient

type InsertionConfirmation struct {
	// The id of the newly-created object.
	Id int32 `json:"id,omitempty"`
	// The number of rows modified
	Rows int32 `json:"rows"`
	// An error message explaining why the method failed
	Error_ string `json:"error,omitempty"`
}
