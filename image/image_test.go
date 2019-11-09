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
			Expect(m.Set(2, 1, "R")).To(Succeed())
			Expect(m.Grid).To(Equal(expected))
		})

		Context("if the x coordinate is out of range", func() {
			It("fails", func() {
				err := m.Set(5, 2, "R")
				Expect(err).To(MatchError("given coordinate is beyond image grid"))
			})
		})

		Context("if the y coordinate is out of range", func() {
			It("fails", func() {
				err := m.Set(2, 5, "R")
				Expect(err).To(MatchError("given coordinate is beyond image grid"))
			})
		})
	})

	Describe("SetMultiY", func() {
		var m image.Image

		BeforeEach(func() {
			m = image.New(2, 3)
		})

		It("sets the given y-axis bits to a value", func() {
			expected := [][]string{{"O", "G"}, {"O", "G"}, {"O", "G"}}
			Expect(m.SetMultiY(2, 1, 3, "G")).To(Succeed())
			Expect(m.Grid).To(Equal(expected))
		})

		Context("if a y coordinate is out of range", func() {
			It("fails", func() {
				err := m.SetMultiY(2, 1, 5, "G")
				Expect(err).To(MatchError("given coordinate is beyond image grid"))
			})
		})
	})

	Describe("Clear", func() {
		var m image.Image

		BeforeEach(func() {
			m = image.New(2, 3)
			m.Set(2, 1, "C")
		})

		It("resets the grid to 'blank'", func() {
			set := [][]string{{"O", "C"}, {"O", "O"}, {"O", "O"}}
			Expect(m.Grid).To(Equal(set))

			cleared := [][]string{{"O", "O"}, {"O", "O"}, {"O", "O"}}
			m.Clear()
			Expect(m.Grid).To(Equal(cleared))
		})
	})

	Describe("Pretty", func() {
		It("prettifies the grid for printing", func() {
			m := image.New(3, 2)
			Expect(m.Pretty()).To(Equal("O O O\nO O O\n"))
		})
	})
})
