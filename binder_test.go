package main_test

import (
	. "github.com/dgruber/randomNumberBroker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Binder", func() {

	Describe("should work as expected and create random numbers", func() {

		It("should be possible to create a binder", func() {
			b := NewBinder()
			Ω(b).ShouldNot(BeNil())

			br, err := b.Bind("instanceID", "bindingID")
			Ω(err).Should(BeNil())
			Ω(br.RandomNumber).ShouldNot(Equal(""))

			exists, err := b.InstanceExists("instanceID")
			Ω(err).Should(BeNil())
			Ω(exists).Should(BeTrue())

			exists, err = b.InstanceExists("doesnotexist")
			Ω(err).Should(BeNil())
			Ω(exists).Should(BeFalse())

			err = b.Unbind("instanceID", "bindingID")
			Ω(err).Should(BeNil())

			exists, err = b.InstanceExists("instanceID")
			Ω(err).Should(BeNil())
			Ω(exists).Should(BeFalse())
		})

	})

})
