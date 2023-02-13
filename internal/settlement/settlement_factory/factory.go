package settlement_factory

import (
	bankSimulator "github.com/sarthakraheja/bank-simulator/protos/v1/github.com/sarthakraheja/bank-simulator/protos"
)

type acquringBankFactory struct {
	acquiringBankClient bankSimulator.AcquiringBankServiceClient
}

type AcquringBankFactory interface {
	BuildJPMorganAcquiringBank() AcquringBank
	BuildChaseAcquringBank() AcquringBank
}

func (a *acquringBankFactory) BuildJPMorganAcquiringBank() AcquringBank {
	return &jpMorgan{
		acquiringBankClient: a.acquiringBankClient,
	}
}

func (a *acquringBankFactory) BuildChaseAcquringBank() AcquringBank {
	return &chase{
		acquiringBankClient: a.acquiringBankClient,
	}
}

func NewAcquiringBankFactory(acquiringBankClient bankSimulator.AcquiringBankServiceClient) AcquringBankFactory {
	return &acquringBankFactory{
		acquiringBankClient: acquiringBankClient,
	}
}
