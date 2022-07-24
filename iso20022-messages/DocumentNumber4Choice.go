package iso20022

// Choice between a short document number, a long document number or a proprietary document number.
type DocumentNumber4Choice struct {

	// Message type number of the document referenced.
	ShortNumber *Exact3NumericText `xml:"ShrtNb"`

	// MX Message identifier of the referenced document.
	LongNumber *ISO20022MessageIdentificationText `xml:"LngNb"`

	// Proprietary document identification.
	ProprietaryNumber *GenericIdentification38 `xml:"PrtryNb"`
}

func (d *DocumentNumber4Choice) SetShortNumber(value string) {
	d.ShortNumber = (*Exact3NumericText)(&value)
}

func (d *DocumentNumber4Choice) SetLongNumber(value string) {
	d.LongNumber = (*ISO20022MessageIdentificationText)(&value)
}

func (d *DocumentNumber4Choice) AddProprietaryNumber() *GenericIdentification38 {
	d.ProprietaryNumber = new(GenericIdentification38)
	return d.ProprietaryNumber
}
