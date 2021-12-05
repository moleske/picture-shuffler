package shuffler_test

import (
	"errors"
	"github.com/moleske/picture-shuffler/shuffler"
	"github.com/moleske/picture-shuffler/shuffler/stubs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filewalker", func() {
	var window stubs.StubWindow
	var directory stubs.StubDirEntry

	BeforeEach(func() {
		window = stubs.StubWindow{}
		directory = stubs.StubDirEntry{IsDirectory: false}
	})

	It("returns error if passed an error", func() {
		myError := errors.New("my error")
		Expect(shuffler.FileWalkerFunction(&window, "somepath", directory, myError)).To(Equal(myError))
	})

	It("returns nil when a directory is given", func() {
		directory = stubs.StubDirEntry{IsDirectory: true}
		Expect(shuffler.FileWalkerFunction(&window, "somepath", directory, nil)).NotTo(HaveOccurred())
	})
})
