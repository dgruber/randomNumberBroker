package main_test

import (
	. "github.com/dgruber/randomNumberBroker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creator", func() {

	Describe("should manage service instance creation", func() {

		It("should be possible to create and delete a service instance", func() {
			c := NewCreator()
			Ω(c).ShouldNot(BeNil())

			err := c.Create("instanceID")
			Ω(err).Should(BeNil())

			exists, err := c.InstanceExists("instanceID")
			Ω(err).Should(BeNil())
			Ω(exists).Should(BeTrue())

			exists, err = c.InstanceExists("doesnotexist")
			Ω(err).Should(BeNil())
			Ω(exists).Should(BeFalse())

			err = c.Destroy("instanceID")
			Ω(err).Should(BeNil())

			exists, err = c.InstanceExists("instanceID")
			Ω(err).Should(BeNil())
			Ω(exists).Should(BeFalse())
		})

	})

})
