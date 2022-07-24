package iso20022

// Choice between a code or a proprietary code for a date code.
type DateCode10Choice struct {

	// Standard code to specify the type of date.
	Code *DateType8Code `xml:"Cd"`

	// Proprietary identification of the type of date.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (d *DateCode10Choice) SetCode(value string) {
	d.Code = (*DateType8Code)(&value)
}

func (d *DateCode10Choice) AddProprietary() *GenericIdentification20 {
	d.Proprietary = new(GenericIdentification20)
	return d.Proprietary
}
