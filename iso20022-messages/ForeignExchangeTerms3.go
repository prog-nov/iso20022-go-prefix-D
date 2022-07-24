package iso20022

// Information needed to process a currency exchange or conversion.
type ForeignExchangeTerms3 struct {

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *CurrencyCode `xml:"UnitCcy"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *CurrencyCode `xml:"QtdCcy"`

	// The value of one currency expressed in relation to another currency. ExchangeRate expresses the ratio between UnitCurrency and QuotedCurrency (ExchangeRate = UnitCurrency/QuotedCurrency).
	ExchangeRate *BaseOneRate `xml:"XchgRate"`

	// Date and time at which an exchange rate is quoted.
	QuotationDate *ISODateTime `xml:"QtnDt,omitempty"`

	// Party that proposes a  foreign exchange rate.
	QuotingInstitution *PartyIdentification1Choice `xml:"QtgInstn,omitempty"`
}

func (f *ForeignExchangeTerms3) SetUnitCurrency(value string) {
	f.UnitCurrency = (*CurrencyCode)(&value)
}

func (f *ForeignExchangeTerms3) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*CurrencyCode)(&value)
}

func (f *ForeignExchangeTerms3) SetExchangeRate(value string) {
	f.ExchangeRate = (*BaseOneRate)(&value)
}

func (f *ForeignExchangeTerms3) SetQuotationDate(value string) {
	f.QuotationDate = (*ISODateTime)(&value)
}

func (f *ForeignExchangeTerms3) AddQuotingInstitution() *PartyIdentification1Choice {
	f.QuotingInstitution = new(PartyIdentification1Choice)
	return f.QuotingInstitution
}
