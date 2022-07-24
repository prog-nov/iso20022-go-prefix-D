package iso20022

// Provides further details specific to the individual direct debit transaction(s) included in the message.
type DirectDebitTransactionInformation13 struct {

	// Set of elements used to reference a payment instruction.
	PaymentIdentification *PaymentIdentification1 `xml:"PmtId"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation24 `xml:"PmtTpInf,omitempty"`

	// Amount of money to be moved between the debtor and creditor, before deduction of charges, expressed in the currency as ordered by the initiating party.
	InstructedAmount *ActiveOrHistoricCurrencyAndAmount `xml:"InstdAmt"`

	// Specifies which party/parties will bear the charges associated with the processing of the payment transaction.
	ChargeBearer *ChargeBearerType1Code `xml:"ChrgBr,omitempty"`

	// Provides information specific to the direct debit mandate.
	DirectDebitTransaction *DirectDebitTransaction7 `xml:"DrctDbtTx,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt"`

	// Unambiguous identification of the account of the debtor agent at its servicing agent in the payment chain.
	DebtorAgentAccount *CashAccount24 `xml:"DbtrAgtAcct,omitempty"`

	// Party that owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr"`

	// Unambiguous identification of the account of the debtor to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Further information, related to the processing of the payment instruction, that may need to be acted upon by the creditor agent, depending on agreement between creditor and the creditor agent.
	InstructionForCreditorAgent *Max140Text `xml:"InstrForCdtrAgt,omitempty"`

	// Underlying reason for the payment transaction.
	// Usage: Purpose is used by the end-customers, that is initiating party, (ultimate) debtor, (ultimate) creditor to provide information concerning the nature of the payment. Purpose is a content element, which is not used for processing by any of the agents involved in the payment chain.
	Purpose *Purpose2Choice `xml:"Purp,omitempty"`

	// Information needed due to regulatory and statutory requirements.
	RegulatoryReporting []*RegulatoryReporting3 `xml:"RgltryRptg,omitempty"`

	// Provides details on the tax.
	Tax *TaxInformation3 `xml:"Tax,omitempty"`

	// Provides information related to the handling of the remittance information by any of the agents in the transaction processing chain.
	RelatedRemittanceInformation []*RemittanceLocation2 `xml:"RltdRmtInf,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, such as commercial invoices in an accounts' receivable system.
	RemittanceInformation *RemittanceInformation7 `xml:"RmtInf,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (d *DirectDebitTransactionInformation13) AddPaymentIdentification() *PaymentIdentification1 {
	d.PaymentIdentification = new(PaymentIdentification1)
	return d.PaymentIdentification
}

func (d *DirectDebitTransactionInformation13) AddPaymentTypeInformation() *PaymentTypeInformation24 {
	d.PaymentTypeInformation = new(PaymentTypeInformation24)
	return d.PaymentTypeInformation
}

func (d *DirectDebitTransactionInformation13) SetInstructedAmount(value, currency string) {
	d.InstructedAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (d *DirectDebitTransactionInformation13) SetChargeBearer(value string) {
	d.ChargeBearer = (*ChargeBearerType1Code)(&value)
}

func (d *DirectDebitTransactionInformation13) AddDirectDebitTransaction() *DirectDebitTransaction7 {
	d.DirectDebitTransaction = new(DirectDebitTransaction7)
	return d.DirectDebitTransaction
}

func (d *DirectDebitTransactionInformation13) AddUltimateCreditor() *PartyIdentification43 {
	d.UltimateCreditor = new(PartyIdentification43)
	return d.UltimateCreditor
}

func (d *DirectDebitTransactionInformation13) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	d.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return d.DebtorAgent
}

func (d *DirectDebitTransactionInformation13) AddDebtorAgentAccount() *CashAccount24 {
	d.DebtorAgentAccount = new(CashAccount24)
	return d.DebtorAgentAccount
}

func (d *DirectDebitTransactionInformation13) AddDebtor() *PartyIdentification43 {
	d.Debtor = new(PartyIdentification43)
	return d.Debtor
}

func (d *DirectDebitTransactionInformation13) AddDebtorAccount() *CashAccount24 {
	d.DebtorAccount = new(CashAccount24)
	return d.DebtorAccount
}

func (d *DirectDebitTransactionInformation13) AddUltimateDebtor() *PartyIdentification43 {
	d.UltimateDebtor = new(PartyIdentification43)
	return d.UltimateDebtor
}

func (d *DirectDebitTransactionInformation13) SetInstructionForCreditorAgent(value string) {
	d.InstructionForCreditorAgent = (*Max140Text)(&value)
}

func (d *DirectDebitTransactionInformation13) AddPurpose() *Purpose2Choice {
	d.Purpose = new(Purpose2Choice)
	return d.Purpose
}

func (d *DirectDebitTransactionInformation13) AddRegulatoryReporting() *RegulatoryReporting3 {
	newValue := new(RegulatoryReporting3)
	d.RegulatoryReporting = append(d.RegulatoryReporting, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation13) AddTax() *TaxInformation3 {
	d.Tax = new(TaxInformation3)
	return d.Tax
}

func (d *DirectDebitTransactionInformation13) AddRelatedRemittanceInformation() *RemittanceLocation2 {
	newValue := new(RemittanceLocation2)
	d.RelatedRemittanceInformation = append(d.RelatedRemittanceInformation, newValue)
	return newValue
}

func (d *DirectDebitTransactionInformation13) AddRemittanceInformation() *RemittanceInformation7 {
	d.RemittanceInformation = new(RemittanceInformation7)
	return d.RemittanceInformation
}

func (d *DirectDebitTransactionInformation13) AddSupplementaryData() *SupplementaryData1 {
	newValue := new(SupplementaryData1)
	d.SupplementaryData = append(d.SupplementaryData, newValue)
	return newValue
}
