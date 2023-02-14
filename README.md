Payments Gateway

Product requirements for the payments service

1. A merchant should be able to process a payment through the payment gateway and receive either a successful or unsuccessful response. 

2. A merchant should be able to retrieve the details of a previously made payment.

Requirements:
The payment gateway will need to provide merchants with a way to process a payment. To do this, the merchant should be able to submit a request to the payment gateway. A payment request should include appropriate fields such as the card number, expiry month/date, amount, currency, and cvv.

The server runs on localhost:9000 and serves the following RPC calls:
    CreatePayment(CreatePaymentRequest) returns CreatePaymentResponse
    GetPayment(GetPaymentRequest) returns GetPaymentResponse


Setting up the payments gateway: 

Note: The service directory also includess the bank simulator(./bank-simulator), which is required for the payments service to be able to fulfil payments request. 

1. [Pre-requisite] Navigate to /bank-simulator. Run cmd: 'make docker-up' while at bank-simulator. 
    This will spin up the bank simulator which is a grpc server that runs on port 9090.
2. Open a new terminal tab, and navigate back to /payments-service. the root directory of this repository. Run cmd: 'make docker-up'
    This will spin up a postgres db with payments table, as well as a grpc server that runs on port 9000
3. Run 'make grpcui.local' in order to spin up the grpc ui tooling provided by (https://github.com/fullstorydev/grpcui)
    GrpcUI will provide a UI in order to interact with the rpc calls and enable end to end testing


Code features:

1. High level technical spec: PostgresDB with grpc server. 
2. Modular pattern providing flexiblity for datastore, api choice(rest vs grpc).
3. Unit tested individual components
4. [Bonus]Encrypted storage of sensitive field using symmetric encryption
5. Factory pattern for connecting to different kinds of acquiring banks encapsulated in the setttlement package
6. input validation for generic payment data parsing
7. Extensible APIs to enable processing of different kind of payment.(Wire Transfers etc)
8. [Bonus] Fully fledged Makefile in order to generate codegen for protos mock,files. running and stopping servers.
9. [Bonus] Grouped imports in order to easily undertand the different internal and external libraries the service depends on. 
10. [Bonus]The CreatePayment rpc call is idempotent to prevent duplicate requests and charges to the end client.


Tooling suport:

1. Grpcui for interacting with the server through a easy to use UI (https://github.com/fullstorydev/grpcui)
2. pgAdmin4 in order to perform read and query operation w the database. (https://www.postgresql.org/ftp/pgadmin/pgadmin4/)


Areas of Improvement
1. Error handling needs to be improved as it can be more user friendly as currently it leaks out internal logic of the service. 
2. The data model can be improved for memory efficient storage (storing the encrypted payment method is currently expensive)
3. Improved monitoring, currently we are not logging anywhere throughout the lifecycle of the request.
4. Integration testing
5. Would want to spin up the bank simulator as part of the boot up process for payment-simulator.  Ran into an issue when attempting to set that up. 

Assumptions:
1. We can succesfully process a payment synchronously
2. The acquiring bank interface provides the Authorise and Capture payment method
3. We can have multiple acquiring bank.
4. Routing based on creditCard type but can be extended to arbritary rules as needed.


Known Issues:
1. Currently obfuscation is being done at the handler level, would want to move that over to the model file. 
2. Some modules(processor,server) are not currentlyunit tested.


Design Doc for more details: https://docs.google.com/document/d/1n-L3vI89nAXaH5CSr0snrFsXOTQsBnrOlXg4c5gAVRM/edit