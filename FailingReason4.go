package iso20022

// The status of an instruction, advice or request.
type FailingReason4 struct {

	// Specifies the reason why the instruction has a failing status.
	Code *FailingReason4Choice `xml:"Cd"`

	// Provides additional reason information that cannot be provided in a structured field.
	AdditionalReasonInformation *Max210Text `xml:"AddtlRsnInf,omitempty"`
}

func (f *FailingReason4) AddCode() *FailingReason4Choice {
	f.Code = new(FailingReason4Choice)
	return f.Code
}

func (f *FailingReason4) SetAdditionalReasonInformation(value string) {
	f.AdditionalReasonInformation = (*Max210Text)(&value)
}
