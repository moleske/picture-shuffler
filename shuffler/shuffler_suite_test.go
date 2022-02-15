package shuffler_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestShuffler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shuffler Suite")
}
