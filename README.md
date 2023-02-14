Payments Service

Product requirements for the payments service

1. A merchant should be able to process a payment through the payment gateway and receive either a successful or unsuccessful response. 

2. A merchant should be able to retrieve the details of a previously made payment. The next section will discuss each of these in more detail.

Requirements:
The payment gateway will need to provide merchants with a way to process a payment. To do this, the merchant should be able to submit a request to the payment gateway. A payment request should include appropriate fields such as the card number, expiry month/date, amount, currency, and cvv.

The server runs on localhost:9000 and serves the following RPC calls:
    CreatePayment(CreatePaymentRequest) returns CreatePaymentResponse
    GetPayment(GetPaymentRequest) returns GetPaymentResponse


Setting up the payments service: 

Note: The service directory also includess the bank simulator(./bank-simulator), which is required for the payments service to be able to fulfil payments request. 

1. Navigate to /bank-simulator. Run cmd: 'make docker-up' while at bank-simulator. 
    This will spin up the bank simulator which is a grpc server that runs on port 9090.
2. Open a new terminal tab, and navigate back to /payments-service. the root directory of this repository. Run cmd: 'make docker-up'
    This will spin up a postgres db with payments table, as well as a grpc server that runs on port 9000
3. Run 'make grpcui.local' in order to spin up the grpc ui tooling provided by (https://github.com/fullstorydev/grpcui)
    GrpcUI will provide a UI in order to interact with the rpc calls and enable end to end testing


Code features:

1. High level technical spec: PostgresDB with grpc server. 
2. Modular pattern providing flexiblity for datastore, api choice(rest vs grpc).
3. Unit tested individual components
4. Encrypted storage of sensitive field using symmetric encryption
5. Factory pattern for connecting to different kinds of acquiring banks encapsulated in the setttlement package
6. input validation for generic payment data parsing
7. Extensible APIs to enable processing of different kind of payment.(Wire Transfers etc)
8. Fully fledged Makefile in order to generate codegen for protos mock,files. running and stopping servers.


Tooling suport:

1. Grpcui for interacting with the server through a easy to use UI
2. pgAdmin4 in order to perform read and query operation w the database. 


Areas of Improvement
