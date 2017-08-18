package main_test

import (
	. "github.com/dgruber/randomNumberBroker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"
	"errors"
	"github.com/pivotal-cf/brokerapi"
)

type fakeBinder struct {
	instances map[string]string
}

func (b *fakeBinder) Bind(instanceID, bindingID string) (BindingResult, error) {
	b.instances[instanceID] = bindingID
	return BindingResult{}, nil
}

func (b *fakeBinder) Unbind(instanceID string, bindingID string) error {
	if exists, _ := b.InstanceExists(instanceID); !exists {
		return errors.New("instance does not exist")
	}
	delete(b.instances, instanceID)
	return nil
}

func (b *fakeBinder) InstanceExists(instanceID string) (bool, error) {
	_, exists := b.instances[instanceID]
	return exists, nil
}

type fakeCreator struct {
	instances map[string]string
}

func (c *fakeCreator) Create(instanceID string) error {
	c.instances[instanceID] = instanceID
	return nil
}

func (c *fakeCreator) Destroy(instanceID string) error {
	if exists, _ := c.InstanceExists(instanceID); !exists {
		return errors.New("instance does not exist")
	}
	delete(c.instances, instanceID)
	return nil
}

func (c *fakeCreator) InstanceExists(instanceID string) (bool, error) {
	_, exists := c.instances[instanceID]
	return exists, nil
}

var _ = Describe("Broker", func() {

	var rnb *RandomNumberBroker

	BeforeEach(func() {
		rnb = &RandomNumberBroker{
			InstanceCreators: map[string]InstanceCreator{
				"0768E956-6650-4010-8E5F-2BBED9D03031": &fakeCreator{instances: make(map[string]string)},
			},
			InstanceBinders: map[string]InstanceBinder{
				"0768E956-6650-4010-8E5F-2BBED9D03031": &fakeBinder{instances: make(map[string]string)},
			},
		}
	})

	Describe("Service instance lifecycle", func() {

		Context("when there is no issue expected", func() {

			It("should provision and bind", func() {
				_, err := rnb.Provision(context.TODO(), "instanceID", brokerapi.ProvisionDetails{PlanID: "0768E956-6650-4010-8E5F-2BBED9D03031"}, false)
				Ω(err).Should(BeNil())
				_, err = rnb.Bind(context.TODO(), "instanceID", "bindingID", brokerapi.BindDetails{})
				Ω(err).Should(BeNil())
				err = rnb.Unbind(context.TODO(), "instanceID", "bindingID", brokerapi.UnbindDetails{})
				Ω(err).Should(BeNil())
				_, err = rnb.Deprovision(context.TODO(), "instanceID", brokerapi.DeprovisionDetails{}, false)
				Ω(err).Should(BeNil())
			})
		})

		Context("in case of issues", func() {

			It("should reject double provisioning", func() {
				_, err := rnb.Provision(context.TODO(), "instanceID", brokerapi.ProvisionDetails{PlanID: "0768E956-6650-4010-8E5F-2BBED9D03031"}, false)
				Ω(err).Should(BeNil())
				_, err = rnb.Provision(context.TODO(), "instanceID", brokerapi.ProvisionDetails{PlanID: "0768E956-6650-4010-8E5F-2BBED9D03031"}, false)
				Ω(err).ShouldNot(BeNil())
				_, err = rnb.Deprovision(context.TODO(), "instanceID", brokerapi.DeprovisionDetails{}, false)
				Ω(err).Should(BeNil())
			})

			It("should reject double binding", func() {
				_, err := rnb.Provision(context.TODO(), "instanceID", brokerapi.ProvisionDetails{PlanID: "0768E956-6650-4010-8E5F-2BBED9D03031"}, false)
				Ω(err).Should(BeNil())
				_, err = rnb.Bind(context.TODO(), "instanceID", "bindingID", brokerapi.BindDetails{})
				Ω(err).Should(BeNil())
				_, err = rnb.Bind(context.TODO(), "instanceID", "bindingID", brokerapi.BindDetails{})
				Ω(err).ShouldNot(BeNil())
				err = rnb.Unbind(context.TODO(), "instanceID", "bindingID", brokerapi.UnbindDetails{})
				Ω(err).Should(BeNil())
				_, err = rnb.Deprovision(context.TODO(), "instanceID", brokerapi.DeprovisionDetails{}, false)
				Ω(err).Should(BeNil())
			})

		})

	})

})
