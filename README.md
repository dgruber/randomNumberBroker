# Random Number Service Broker for Cloud Foundry

This is an example of a minimal service broker written in Go. It is based on 
the code and structure of the official examples (kudos to the Pivotal redis
service broker team) and implements a minimal random number service.

## Installation

Create the service broker as Cloud Foundry admin. Replace username and password.

Push the service broker app. Adapt the _manifest.yml_.

    cf push 

Register the service broker in space.

    cf create-service-broker randombroker daniel 63f53854eacffbb5aa36ae91f1f827c9 https://random-number-service-broker.cfapps.io --space-scoped

List service brokers:

    cf service-brokers

See it in the marketplace:

    cf m

Test it.

    curl -s http://random-number-service-broker.cfapps.io/v2/catalog -u daniel:63f53854eacffbb5aa36ae91f1f827c9

## Usage

Create an instance of the Random Number Service broker:

    cf cs RandomNumberBroker default rand

Push an app:

    cd $GOPATH/github.com/dgruber/cf-inspect
    cf push --no-start

Bind the app:

    cf bs inspect rand
    cf start inspect

If using this [cf-inspect](https://github.com/dgruber/cf-inspect) thing go to the url and see the random number appears as enviornment variable:

Key | Value
--- | --- 
Name | rand
Label | RandomNumberBroker
Plan | default
Tags | [random example]
Credential | 6099489395697162318
