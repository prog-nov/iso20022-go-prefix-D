package iso20022

// Specifies whether the status is provided with a reason or not.
type DeniedStatus1Choice struct {

	// Indicates that there is no reason available or to report.
	NoSpecifiedReason *NoReasonCode `xml:"NoSpcfdRsn"`

	// Specifies the reason of the DeniedStatus.
	Reason []*DeniedReason1 `xml:"Rsn,omitempty"`
}

func (d *DeniedStatus1Choice) SetNoSpecifiedReason(value string) {
	d.NoSpecifiedReason = (*NoReasonCode)(&value)
}

func (d *DeniedStatus1Choice) AddReason() *DeniedReason1 {
	newValue := new(DeniedReason1)
	d.Reason = append(d.Reason, newValue)
	return newValue
}
