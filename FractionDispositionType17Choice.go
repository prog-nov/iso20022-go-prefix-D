package iso20022

// Choice between a standard code or a proprietary code to specify the type of fraction disposition.
type FractionDispositionType17Choice struct {

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	Code *FractionDispositionType10Code `xml:"Cd"`

	// Proprietary identification of the type of fraction disposition.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (f *FractionDispositionType17Choice) SetCode(value string) {
	f.Code = (*FractionDispositionType10Code)(&value)
}

func (f *FractionDispositionType17Choice) AddProprietary() *GenericIdentification20 {
	f.Proprietary = new(GenericIdentification20)
	return f.Proprietary
}
