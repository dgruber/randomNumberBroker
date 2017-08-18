package main

import (
	"context"
	"errors"
	"github.com/pivotal-cf/brokerapi"
)

// see also: https://github.com/pivotal-cf/cf-redis-broker/blob/master/broker/broker.go

type RandomNumberBroker struct {
	// each plan has its own creator:
	// - Create
	// - Destroy
	// - InstanceExists
	InstanceCreators map[string]InstanceCreator
	// each plan has its own binder
	// - Bind
	// - Unbind
	// - InstanceExists
	InstanceBinders map[string]InstanceBinder
}

func (randomNumberBroker *RandomNumberBroker) Services(context context.Context) []brokerapi.Service {
	return createServiceDescription()
}

// Provision ...
func (randomNumberBroker *RandomNumberBroker) Provision(context context.Context, instanceID string, serviceDetails brokerapi.ProvisionDetails, asyncAllowed bool) (spec brokerapi.ProvisionedServiceSpec, err error) {
	spec = brokerapi.ProvisionedServiceSpec{}

	if randomNumberBroker.instanceExists(instanceID) {
		return spec, brokerapi.ErrInstanceAlreadyExists
	}

	if serviceDetails.PlanID == "" {
		return spec, errors.New("Plan ID required")
	}

	if planExists(serviceDetails.PlanID) == false {
		return spec, errors.New("Plan not found")
	}

	instanceCreator, ok := randomNumberBroker.InstanceCreators[serviceDetails.PlanID]
	if !ok {
		return spec, errors.New("Internal error: Instance Creator not registered for plan.")
	}

	if err = instanceCreator.Create(instanceID); err != nil {
		return spec, err
	}

	return spec, nil
}

func (randomNumberBroker *RandomNumberBroker) Deprovision(context context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	spec := brokerapi.DeprovisionServiceSpec{}

	for _, instanceCreator := range randomNumberBroker.InstanceCreators {
		instanceExists, _ := instanceCreator.InstanceExists(instanceID)
		if instanceExists {
			return spec, instanceCreator.Destroy(instanceID)
		}
	}
	return spec, brokerapi.ErrInstanceDoesNotExist
}

func (randomNumberBroker *RandomNumberBroker) Bind(context context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	binding := brokerapi.Binding{}

	for _, repo := range randomNumberBroker.InstanceBinders {
		instanceExists, _ := repo.InstanceExists(instanceID)
		if !instanceExists {
			bindingResult, err := repo.Bind(instanceID, bindingID)
			if err != nil {
				return binding, err
			}
			credentialsMap := map[string]interface{}{
				"RANDOM_NUMBER": bindingResult.RandomNumber,
			}

			binding.Credentials = credentialsMap
			return binding, nil
		}
	}

	return brokerapi.Binding{}, brokerapi.ErrInstanceDoesNotExist
}

func (randomNumberBroker *RandomNumberBroker) Unbind(context context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	for _, repo := range randomNumberBroker.InstanceBinders {
		instanceExists, _ := repo.InstanceExists(instanceID)
		if instanceExists {
			err := repo.Unbind(instanceID, bindingID)
			if err != nil {
				return brokerapi.ErrBindingDoesNotExist
			}
			return nil
		}
	}
	return brokerapi.ErrInstanceDoesNotExist
}

func (randomNumberBroker *RandomNumberBroker) instanceExists(instanceID string) bool {
	for _, instanceCreator := range randomNumberBroker.InstanceCreators {
		instanceExists, _ := instanceCreator.InstanceExists(instanceID)
		if instanceExists {
			return true
		}
	}
	return false
}

// LastOperation ...
// If the broker provisions asynchronously, the Cloud Controller will poll this endpoint
// for the status of the provisioning operation.
func (randomNumberBroker *RandomNumberBroker) LastOperation(context context.Context, instanceID, operationData string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, nil
}

func (randomNumberBroker *RandomNumberBroker) Update(context context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, nil
}
