Bank Simulator

Description: grpc server that has 2 defines the following rpc endpoints. 
1. AuthorisePayment (AuthoriseTransactionRequest -> AuthoriseTransactionResponse)
    Used to authorise credit card payments by validating the request is appropriate for fullfillment

2. CapturePayment (CapturePaymentRequest -> CapturePaymentResponse)
    Used to fullfil the payment request from the payment gateway

Server runs on port 9090. 
The Payment gateway relies on the Bank simulator in order to fulfil create payment request. 

The make docker-up command in the payments-service repository spins up the server at port 9090. 

Since this is a simulator, the 20% of requests to the endpoint would return a failure response

