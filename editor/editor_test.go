package editor_test

import (
	"github.com/mo-work/go-technical-test-for-claudia/editor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Editor", func() {
	var e editor.Editor

	Describe("New", func() {
		It("generates a 'blank' grid", func() {
			expected := [][]string{{"O", "O"}, {"O", "O"}}

			e = editor.New(2, 2)
			Expect(e.Grid).To(Equal(expected))
		})
	})

	Describe("Set", func() {

		BeforeEach(func() {
			e = editor.New(2, 3)
		})

		It("sets the given bit to a value", func() {
			expected := [][]string{{"O", "R"}, {"O", "O"}, {"O", "O"}}
			Expect(e.Set(2, 1, "R")).To(Succeed())
			Expect(e.Grid).To(Equal(expected))
		})

		Context("if the x coordinate is out of range", func() {
			It("fails", func() {
				err := e.Set(5, 2, "R")
				Expect(err).To(MatchError("given coordinate is beyond image grid"))
			})
		})

		Context("if the y coordinate is out of range", func() {
			It("fails", func() {
				err := e.Set(2, 5, "R")
				Expect(err).To(MatchError("given coordinate is beyond image grid"))
			})
		})
	})

	Describe("SetMultiY", func() {
		BeforeEach(func() {
			e = editor.New(2, 3)
		})

		It("sets the given y-axis bits to a value", func() {
			expected := [][]string{{"O", "G"}, {"O", "G"}, {"O", "G"}}
			Expect(e.SetMultiY(2, 1, 3, "G")).To(Succeed())
			Expect(e.Grid).To(Equal(expected))
		})

		Context("if y1 is greater than y2", func() {
			It("a line will still be drawn", func() {
				expected := [][]string{{"O", "G"}, {"O", "G"}, {"O", "G"}}
				Expect(e.SetMultiY(2, 3, 1, "G")).To(Succeed())
				Expect(e.Grid).To(Equal(expected))
			})
		})

		Context("if a y coordinate is out of range", func() {
			It("fails", func() {
				err := e.SetMultiY(2, 1, 5, "G")
				Expect(err).To(MatchError("given coordinate is beyond image grid"))
			})
		})
	})

	Describe("SetMultiX", func() {
		BeforeEach(func() {
			e = editor.New(3, 2)
		})

		It("sets the given x-axis bits to a value", func() {
			expected := [][]string{{"B", "B", "B"}, {"O", "O", "O"}}
			Expect(e.SetMultiX(1, 3, 1, "B")).To(Succeed())
			Expect(e.Grid).To(Equal(expected))
		})

		Context("if x1 is greater than x2", func() {
			It("a line will still be drawn", func() {
				expected := [][]string{{"B", "B", "B"}, {"O", "O", "O"}}
				Expect(e.SetMultiX(3, 1, 1, "B")).To(Succeed())
				Expect(e.Grid).To(Equal(expected))
			})
		})

		Context("if an x coordinate is out of range", func() {
			It("fails", func() {
				err := e.SetMultiX(1, 5, 1, "B")
				Expect(err).To(MatchError("given coordinate is beyond image grid"))
			})
		})
	})

	Describe("Clear", func() {
		BeforeEach(func() {
			e = editor.New(2, 3)
			e.Set(2, 1, "C")
		})

		It("resets the grid to 'blank'", func() {
			set := [][]string{{"O", "C"}, {"O", "O"}, {"O", "O"}}
			Expect(e.Grid).To(Equal(set))

			cleared := [][]string{{"O", "O"}, {"O", "O"}, {"O", "O"}}
			e.Clear()
			Expect(e.Grid).To(Equal(cleared))
		})
	})

	Describe("Pretty", func() {
		It("prettifies the grid for printing", func() {
			e = editor.New(3, 2)
			Expect(e.Pretty()).To(Equal("OOO\nOOO\n"))
		})
	})
})
