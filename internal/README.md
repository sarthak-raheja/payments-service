Internal Package descriptions: 

Name Of Package | Dependencies 
Role

Server | Processor, Marshaller, Unmarshaller, Validator   
Role: Server implements handlers and fulfils the request through the handlers.
Depends on Marshaller and Unmarshaller in order to transform internal models to ProtoBuf files
Depends on Validator for field validation
Depends on the Processor package for high level orchestration of the payment. 

Processor | Repository, Router   
Role: Processor is responsible for high level orchestration of the payment
Depends on Settlement Router in order to route the payment request to the appropriate
settlement client
Depends on the repo for all interactions with the database.

Repository | db, Cipher
Role: Wrapper around the database, used to interact with the Underlying postgres db
Depends on the cipher for encryption of senstive fields

Marshaller
Role: Marshalls internal objects to the Protobuf format for service response

Unmarshaller
Role: Unmarshalls Protobuf format to internal objects for processing 

Cipher:
Role: Used to encrypt and decrypt bytes arrays.

Utils:
Role: Generic utility functions, currently used for obfuscating payment method details

Settlement
    Settlement Router: 
        Role: Used in order to route a payment to the appropriate Acquring Bank. 
        Currently uses the CreditCard Type in order to determine the bank to route to. 
    Settlement Factory: 
        Role: Used to construct different kind of acquiring banks. 
        Seperate package in order to abstract the complexity of setting up clients for different kind of acquiring banks

