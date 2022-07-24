package iso20022

// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
type FinancialInstrument30 struct {

	// Unique and unambiguous identifier of a security, assigned under a formal or proprietary identification scheme.
	Identification *SecurityIdentification3Choice `xml:"Id"`

	// Name of the financial instrument in free format text.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Specifies the quantity of assets to be transferred in units or in a percentage rate.
	Quantity *Quantity12Choice `xml:"Qty,omitempty"`

	// Average cost per share of a security, including all charges and commissions.
	AverageAcquisitionPrice *ActiveOrHistoricCurrencyAndAmount `xml:"AvrgAcqstnPric,omitempty"`

	// Net asset on balance sheet - total portfolio value minus or plus the unrealised gain or loss.
	TotalBookValue *ActiveOrHistoricCurrencyAndAmount `xml:"TtlBookVal,omitempty"`

	// Account held in the name of a party that is not the name of the beneficial owner of the shares.
	TransfereeAccount *Account6 `xml:"TrfeeAcct,omitempty"`

	// Party that delivers securities to the receiving agent at the place of settlement, for example, a central securities depository.
	DeliveringAgentDetails *PartyIdentificationAndAccount93 `xml:"DlvrgAgtDtls,omitempty"`
}

func (f *FinancialInstrument30) AddIdentification() *SecurityIdentification3Choice {
	f.Identification = new(SecurityIdentification3Choice)
	return f.Identification
}

func (f *FinancialInstrument30) SetName(value string) {
	f.Name = (*Max350Text)(&value)
}

func (f *FinancialInstrument30) AddQuantity() *Quantity12Choice {
	f.Quantity = new(Quantity12Choice)
	return f.Quantity
}

func (f *FinancialInstrument30) SetAverageAcquisitionPrice(value, currency string) {
	f.AverageAcquisitionPrice = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument30) SetTotalBookValue(value, currency string) {
	f.TotalBookValue = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrument30) AddTransfereeAccount() *Account6 {
	f.TransfereeAccount = new(Account6)
	return f.TransfereeAccount
}

func (f *FinancialInstrument30) AddDeliveringAgentDetails() *PartyIdentificationAndAccount93 {
	f.DeliveringAgentDetails = new(PartyIdentificationAndAccount93)
	return f.DeliveringAgentDetails
}
