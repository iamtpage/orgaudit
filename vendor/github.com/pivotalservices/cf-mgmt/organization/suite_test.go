package organization_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var test *testing.T

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	test = t
	RunSpecs(t, "Test Suite")

}
