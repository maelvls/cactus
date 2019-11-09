package image_test

import (
	"github.com/mo-work/go-technical-test-for-claudia/image"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Image", func() {
	Describe("New", func() {
		It("generates a 'blank' grid", func() {
			expected := [][]string{{"O", "O"}, {"O", "O"}}

			m := image.New(2, 2)
			Expect(m.Grid).To(Equal(expected))
		})
	})

	Describe("Set", func() {
		var m image.Image

		BeforeEach(func() {
			m = image.New(2, 3)
		})

		It("sets the given bit to a value", func() {
			expected := [][]string{{"O", "R"}, {"O", "O"}, {"O", "O"}}
			m.Set(2, 1, "R")
			Expect(m.Grid).To(Equal(expected))
		})
	})

	Describe("Pretty", func() {
		It("prettifies the grid for printing", func() {
			m := image.New(3, 2)
			Expect(m.Pretty()).To(Equal("O O O\nO O O\n"))
		})
	})
})
