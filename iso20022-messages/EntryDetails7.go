package iso20022

// Identifies the underlying transaction(s) and/or batched entries.
type EntryDetails7 struct {

	// Provides details on batched transactions.
	Batch *BatchInformation2 `xml:"Btch,omitempty"`

	// Provides information on the underlying transaction(s).
	TransactionDetails []*EntryTransaction8 `xml:"TxDtls,omitempty"`
}

func (e *EntryDetails7) AddBatch() *BatchInformation2 {
	e.Batch = new(BatchInformation2)
	return e.Batch
}

func (e *EntryDetails7) AddTransactionDetails() *EntryTransaction8 {
	newValue := new(EntryTransaction8)
	e.TransactionDetails = append(e.TransactionDetails, newValue)
	return newValue
}
