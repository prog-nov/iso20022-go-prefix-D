package iso20022

// Parameters applied to the settlement of a security transfer.
type DeliverInformation12 struct {

	// Date and time at which the securities are to be exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	RequestedSettlementDate *ISODate `xml:"ReqdSttlmDt,omitempty"`

	// Date and time at which the securities were exchanged at the International Central Securities Depository (ICSD) or Central Securities Depository (CSD).
	EffectiveSettlementDate *DateAndDateTimeChoice `xml:"FctvSttlmDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Indicates whether the settlement amount includes the stamp duty amount.
	StampDuty *StampDutyType2Code `xml:"StmpDty,omitempty"`

	// Deal amount.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Charge related to the transfer of a financial instrument.
	ChargeDetails []*Charge20 `xml:"ChrgDtls,omitempty"`

	// Commission related to the transfer of a financial instrument.
	CommissionDetails []*Commission17 `xml:"ComssnDtls,omitempty"`

	// Tax related to the transfer of a financial instrument.
	TaxDetails []*Tax21 `xml:"TaxDtls,omitempty"`

	// Specifies foreign exchange details applied to the payment of charges, taxes and commissions as a result of the transfer.
	ForeignExchangeDetails []*ForeignExchangeTerms7 `xml:"FXDtls,omitempty"`

	// Chain of parties involved in the settlement of a transaction.
	SettlementPartiesDetails *DeliveringPartiesAndAccount9 `xml:"SttlmPtiesDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically  delivered.
	PhysicalTransfer *PhysicalTransferType1Code `xml:"PhysTrf,omitempty"`

	// Parameters of a physical delivery.
	PhysicalTransferDetails *DeliveryParameters4 `xml:"PhysTrfDtls,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`
}

func (d *DeliverInformation12) SetRequestedSettlementDate(value string) {
	d.RequestedSettlementDate = (*ISODate)(&value)
}

func (d *DeliverInformation12) AddEffectiveSettlementDate() *DateAndDateTimeChoice {
	d.EffectiveSettlementDate = new(DateAndDateTimeChoice)
	return d.EffectiveSettlementDate
}

func (d *DeliverInformation12) SetSettlementAmount(value, currency string) {
	d.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation12) SetStampDuty(value string) {
	d.StampDuty = (*StampDutyType2Code)(&value)
}

func (d *DeliverInformation12) SetNetAmount(value, currency string) {
	d.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (d *DeliverInformation12) AddChargeDetails() *Charge20 {
	newValue := new(Charge20)
	d.ChargeDetails = append(d.ChargeDetails, newValue)
	return newValue
}

func (d *DeliverInformation12) AddCommissionDetails() *Commission17 {
	newValue := new(Commission17)
	d.CommissionDetails = append(d.CommissionDetails, newValue)
	return newValue
}

func (d *DeliverInformation12) AddTaxDetails() *Tax21 {
	newValue := new(Tax21)
	d.TaxDetails = append(d.TaxDetails, newValue)
	return newValue
}

func (d *DeliverInformation12) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	d.ForeignExchangeDetails = append(d.ForeignExchangeDetails, newValue)
	return newValue
}

func (d *DeliverInformation12) AddSettlementPartiesDetails() *DeliveringPartiesAndAccount9 {
	d.SettlementPartiesDetails = new(DeliveringPartiesAndAccount9)
	return d.SettlementPartiesDetails
}

func (d *DeliverInformation12) SetPhysicalTransfer(value string) {
	d.PhysicalTransfer = (*PhysicalTransferType1Code)(&value)
}

func (d *DeliverInformation12) AddPhysicalTransferDetails() *DeliveryParameters4 {
	d.PhysicalTransferDetails = new(DeliveryParameters4)
	return d.PhysicalTransferDetails
}

func (d *DeliverInformation12) SetClientReference(value string) {
	d.ClientReference = (*Max35Text)(&value)
}
