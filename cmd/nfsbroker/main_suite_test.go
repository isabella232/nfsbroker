package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

var binaryPath string

func TestLocalbroker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Knfsbroker Main Suite")
}

var _ = SynchronizedBeforeSuite(func() []byte {
	var err error
	binaryPath, err = gexec.Build("github.com/cloudfoundry-community/nfsbroker/cmd/nfsbroker", "-race")
	Expect(err).NotTo(HaveOccurred())

	return []byte(binaryPath)
}, func(bytes []byte) {
	binaryPath = string(bytes)
})