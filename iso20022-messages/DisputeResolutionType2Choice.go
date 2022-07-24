package iso20022

// Choice between a code or a proprietary code as to the nature of the dispute about the collateral amount.
type DisputeResolutionType2Choice struct {

	// Code to specify the type of dispute that is to be resolved regarding the disputed collateral amount.
	Code *DisputeResolutionType2Code `xml:"Cd"`

	// Proprietary identification of the type of dispute that is to be resolved regarding the disputed collateral amount.
	ProprietaryIdentification *GenericIdentification30 `xml:"PrtryId"`
}

func (d *DisputeResolutionType2Choice) SetCode(value string) {
	d.Code = (*DisputeResolutionType2Code)(&value)
}

func (d *DisputeResolutionType2Choice) AddProprietaryIdentification() *GenericIdentification30 {
	d.ProprietaryIdentification = new(GenericIdentification30)
	return d.ProprietaryIdentification
}
