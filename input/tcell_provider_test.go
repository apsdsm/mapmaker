package input_test

import (
	"github.com/apsdsm/mapmaker/fakes"
	. "github.com/apsdsm/mapmaker/input"
	"github.com/gdamore/tcell"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TCell Provider", func() {
	It("gets events from source", func() {
		source := fakes.NewInputSource()
		provider := NewTcellProvider(source)

		provider.GetInput()

		Expect(source.Received("PollEvent")).To(BeTrue())
	})

	It("converts 'q' to the Quit command", func() {
		source := fakes.NewInputSource()
		provider := NewTcellProvider(source)
		event := tcell.NewEventKey(tcell.KeyRune, 'q', tcell.ModNone)
		source.PollEventReturns(event)

		code := provider.GetInput()

		Expect(code).To(Equal(Quit))
	})
})
