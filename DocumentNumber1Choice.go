package iso20022

// Choice between a short document number, a long document number or a proprietary document number.
type DocumentNumber1Choice struct {

	// Message type number of the document referenced.
	ShortNumber *Exact3NumericText `xml:"ShrtNb"`

	// MX Message identifier of the referenced document.
	LongNumber *ISO20022MessageIdentificationText `xml:"LngNb"`

	// Proprietary document identification.
	ProprietaryNumber *GenericIdentification19 `xml:"PrtryNb"`
}

func (d *DocumentNumber1Choice) SetShortNumber(value string) {
	d.ShortNumber = (*Exact3NumericText)(&value)
}

func (d *DocumentNumber1Choice) SetLongNumber(value string) {
	d.LongNumber = (*ISO20022MessageIdentificationText)(&value)
}

func (d *DocumentNumber1Choice) AddProprietaryNumber() *GenericIdentification19 {
	d.ProprietaryNumber = new(GenericIdentification19)
	return d.ProprietaryNumber
}
