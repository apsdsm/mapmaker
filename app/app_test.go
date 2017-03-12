package app_test

import (
	"github.com/apsdsm/mapmaker/app"
	"github.com/apsdsm/mapmaker/fakes"
	"github.com/apsdsm/mapmaker/input"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {

	var ()

	Context("Application loop", func() {
		It("checks for input each update", func() {
			s := fakes.NewScreen()
			i := fakes.NewInputProvider()
			a := app.NewApp(s, i)

			a.Update()

			Expect(i.Received("GetInput")).To(BeTrue())
		})
	})

	Context("Quitting", func() {
		It("sets the done flag to true if receives quit input", func() {
			s := fakes.NewScreen()
			i := fakes.NewInputProvider()
			a := app.NewApp(s, i)
			i.GetInputReturns(input.Quit)

			a.Update()

			Expect(a.Done).To(Equal(true))
		})
	})
})
