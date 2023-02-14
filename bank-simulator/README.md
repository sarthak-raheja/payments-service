Acquiring Bank Simulator

Description: grpc server that has 2 defines the following rpc endpoints. 

1. AuthorisePayment (AuthoriseTransactionRequest -> AuthoriseTransactionResponse)
    Used to authorise credit card payments by validating the request is appropriate for fullfillment

2. CapturePayment (CapturePaymentRequest -> CapturePaymentResponse)
    Used to fullfil the payment request from the payment gateway.

Server runs on port 9090. 
The Payment gateway relies on the Bank simulator in order to fulfil create payment request. 


Since this is a simulator, we are also simulating failure cases. 
    For AuthorisePayment RPC call, we would fail 20% of the requests to that endpoint using the math/rand library.
    For CapturePayment RPC call, we would fail 10% of the requests to that endpoint.

To adjust the probability of failure, we can update the conditional at bank-simulator/server/server.go (L22 & L31)
    For a failure rate of 50%: rand<5
    For a failure rate of 100%: rand<10
    
The acquiring bank simulator is used for authorising and fulfilments of credit card payments only. 