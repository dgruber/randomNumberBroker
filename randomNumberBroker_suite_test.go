package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRandomNumberBroker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RandomNumberBroker Suite")
}
