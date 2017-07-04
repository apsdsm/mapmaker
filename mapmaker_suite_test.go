package mapmaker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMapmaker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mapmaker Suite")
}
