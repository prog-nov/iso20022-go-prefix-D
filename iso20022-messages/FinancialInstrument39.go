package iso20022

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument39 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies whether the financial instrument is transferred as an asset or as cash.
	TransferType *TransferType1Code `xml:"TrfTp"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveOrHistoricCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Identifies the currency to be used to transfer the holdings. Some transfer agents register holdings grouped by currency in addition to using the ISIN for multi-currency fund shares.
	TransferCurrency *CurrencyCode `xml:"TrfCcy,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account16 `xml:"TrfeeAcct,omitempty"`

	// Sub-accounts that are grouped in a master or omnibus account.
	SubAccountDetails *SubAccount1 `xml:"SubAcctDtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount93 `xml:"RcvgAgtDtls,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (f *FinancialInstrument39) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument39) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument39) SetTransferType(value string) {
	f.TransferType = (*TransferType1Code)(&value)
}

func (f *FinancialInstrument39) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument39) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument39) SetTransferCurrency(value string) {
	f.TransferCurrency = (*CurrencyCode)(&value)
}

func (f *FinancialInstrument39) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument39) AddTransfereeAccount() *Account16 {
	f.TransfereeAccount = new(Account16)
	return f.TransfereeAccount
}

func (f *FinancialInstrument39) AddSubAccountDetails() *SubAccount1 {
	f.SubAccountDetails = new(SubAccount1)
	return f.SubAccountDetails
}

func (f *FinancialInstrument39) AddReceivingAgentDetails() *PartyIdentificationAndAccount93 {
	f.ReceivingAgentDetails = new(PartyIdentificationAndAccount93)
	return f.ReceivingAgentDetails
}

func (f *FinancialInstrument39) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return f.DeliveringAgentDetails
}
