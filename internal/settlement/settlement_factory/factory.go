package settlement_factory

type acquringBankFactory struct{}

type AcquringBankFactory interface {
	BuildJPMorganAcquiringBank() AcquringBank
	BuildChaseAcquringBank() AcquringBank
}

func (*acquringBankFactory) BuildJPMorganAcquiringBank() AcquringBank {
	return &jpMorgan{}
}

func (*acquringBankFactory) BuildChaseAcquringBank() AcquringBank {
	return &chase{}
}

func NewAcquiringBankFactory() AcquringBankFactory {
	return &acquringBankFactory{}
}
