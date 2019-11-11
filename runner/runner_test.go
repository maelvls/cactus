package runner_test

import (
	"bufio"
	"errors"
	"io"

	"github.com/mo-work/go-technical-test-for-claudia/runner"
	"github.com/mo-work/go-technical-test-for-claudia/runner/runnerfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Runner", func() {
	var (
		inBuf           *gbytes.Buffer
		outBuf          *gbytes.Buffer
		r               runner.Runner
		fakeImageEditor *runnerfakes.FakeImageEditor
	)

	BeforeEach(func() {
		inBuf = gbytes.NewBuffer()
		scanner := bufio.NewScanner(inBuf)
		outBuf = gbytes.NewBuffer()
		fakeImageEditor = new(runnerfakes.FakeImageEditor)
		r = runner.New(scanner, outBuf, fakeImageEditor)
	})

	Describe("ProcessImageSize", func() {
		It("receives the image size from stdin", func() {
			_, err := io.WriteString(inBuf, "I 5 7")
			Expect(err).NotTo(HaveOccurred())

			Expect(r.ProcessImageSize()).To(Succeed())
			Expect(fakeImageEditor.CreateImageCallCount()).To(Equal(1))
			rows, cols := fakeImageEditor.CreateImageArgsForCall(0)
			Expect(rows).To(Equal(5))
			Expect(cols).To(Equal(7))
		})

		It("upcases the image command, and doesn't fail", func() {
			_, err := io.WriteString(inBuf, "i 5 7")
			Expect(err).NotTo(HaveOccurred())

			Expect(r.ProcessImageSize()).To(Succeed())
			Expect(fakeImageEditor.CreateImageCallCount()).To(Equal(1))
			rows, cols := fakeImageEditor.CreateImageArgsForCall(0)
			Expect(rows).To(Equal(5))
			Expect(cols).To(Equal(7))
		})

		Context("if the argument for the x azis cannot be translated into an integer", func() {
			It("fails", func() {
				_, err := io.WriteString(inBuf, "I x 7")
				Expect(err).NotTo(HaveOccurred())

				Expect(r.ProcessImageSize()).To(MatchError("could not parse non-integer 'x'"))
				Expect(fakeImageEditor.CreateImageCallCount()).To(Equal(0))
			})
		})

		Context("if the argument for the y azis cannot be translated into an integer", func() {
			It("fails", func() {
				_, err := io.WriteString(inBuf, "I 7 y")
				Expect(err).NotTo(HaveOccurred())

				Expect(r.ProcessImageSize()).To(MatchError("could not parse non-integer 'y'"))
				Expect(fakeImageEditor.CreateImageCallCount()).To(Equal(0))
			})
		})
		//
		Context("if the argument for the x axis size is less than the min value", func() {
			It("fails", func() {
				_, err := io.WriteString(inBuf, "I 0 7")
				Expect(err).NotTo(HaveOccurred())

				Expect(r.ProcessImageSize()).To(MatchError("image axis out of range: 1 <= M,N <= 1024"))
				Expect(fakeImageEditor.CreateImageCallCount()).To(Equal(0))
			})
		})

		Context("if the argument for the y axis size is less than the min value", func() {
			It("fails", func() {
				_, err := io.WriteString(inBuf, "I 7 0")
				Expect(err).NotTo(HaveOccurred())

				Expect(r.ProcessImageSize()).To(MatchError("image axis out of range: 1 <= M,N <= 1024"))
				Expect(fakeImageEditor.CreateImageCallCount()).To(Equal(0))
			})
		})

		Context("if the argument for the x axis size is greater than the max value", func() {
			It("fails", func() {
				_, err := io.WriteString(inBuf, "I 1025 7")
				Expect(err).NotTo(HaveOccurred())

				Expect(r.ProcessImageSize()).To(MatchError("image axis out of range: 1 <= M,N <= 1024"))
				Expect(fakeImageEditor.CreateImageCallCount()).To(Equal(0))
			})
		})

		Context("if the argument for the y axis size is greater than the max value", func() {
			It("fails", func() {
				_, err := io.WriteString(inBuf, "I 7 1025")
				Expect(err).NotTo(HaveOccurred())

				Expect(r.ProcessImageSize()).To(MatchError("image axis out of range: 1 <= M,N <= 1024"))
				Expect(fakeImageEditor.CreateImageCallCount()).To(Equal(0))
			})
		})
	})

	Describe("ProcessEditActions", func() {
		It("forwards Set instructions to the editor", func() {
			_, err := io.WriteString(inBuf, "L 1 3 A")
			Expect(err).NotTo(HaveOccurred())

			r.ProcessEditActions()

			Expect(fakeImageEditor.SetCallCount()).To(Equal(1))
			x, y, char := fakeImageEditor.SetArgsForCall(0)
			Expect(x).To(Equal(1))
			Expect(y).To(Equal(3))
			Expect(char).To(Equal("A"))
		})

		Context("if calling Set on the editor fails", func() {
			BeforeEach(func() {
				fakeImageEditor.SetReturns(errors.New("EXPLODE"))
			})

			It("forwards the error", func() {
				_, err := io.WriteString(inBuf, "L 1 3 A")
				Expect(err).NotTo(HaveOccurred())

				r.ProcessEditActions()
				Expect(fakeImageEditor.SetCallCount()).To(Equal(1))
				Expect(outBuf).To(gbytes.Say("EXPLODE"))
			})
		})

		It("forwards SetMultiX instructions to the editor", func() {
			_, err := io.WriteString(inBuf, "H 3 5 2 Z")
			Expect(err).NotTo(HaveOccurred())

			r.ProcessEditActions()

			Expect(fakeImageEditor.SetMultiXCallCount()).To(Equal(1))
			x1, x2, y, char := fakeImageEditor.SetMultiXArgsForCall(0)
			Expect(x1).To(Equal(3))
			Expect(x2).To(Equal(5))
			Expect(y).To(Equal(2))
			Expect(char).To(Equal("Z"))
		})

		Context("if calling SetMultiXCallCount on the editor fails", func() {
			BeforeEach(func() {
				fakeImageEditor.SetMultiXReturns(errors.New("EXPLODE"))
			})

			It("forwards the error", func() {
				_, err := io.WriteString(inBuf, "H 3 5 2 Z")
				Expect(err).NotTo(HaveOccurred())

				r.ProcessEditActions()
				Expect(fakeImageEditor.SetMultiXCallCount()).To(Equal(1))
				Expect(outBuf).To(gbytes.Say("EXPLODE"))
			})
		})

		It("forwards SetMultiY instructions to the editor", func() {
			_, err := io.WriteString(inBuf, "V 2 3 5 W")
			Expect(err).NotTo(HaveOccurred())

			r.ProcessEditActions()

			Expect(fakeImageEditor.SetMultiYCallCount()).To(Equal(1))
			x, y1, y2, char := fakeImageEditor.SetMultiYArgsForCall(0)
			Expect(x).To(Equal(2))
			Expect(y1).To(Equal(3))
			Expect(y2).To(Equal(5))
			Expect(char).To(Equal("W"))
		})

		Context("if calling SetMultiYCallCount on the editor fails", func() {
			BeforeEach(func() {
				fakeImageEditor.SetMultiYReturns(errors.New("EXPLODE"))
			})

			It("forwards the error", func() {
				_, err := io.WriteString(inBuf, "V 2 3 5 W")
				Expect(err).NotTo(HaveOccurred())

				r.ProcessEditActions()
				Expect(fakeImageEditor.SetMultiYCallCount()).To(Equal(1))
				Expect(outBuf).To(gbytes.Say("EXPLODE"))
			})
		})

		It("forwards Show instructions to the editor", func() {
			_, err := io.WriteString(inBuf, "S")
			Expect(err).NotTo(HaveOccurred())

			r.ProcessEditActions()

			Expect(fakeImageEditor.PrettyCallCount()).To(Equal(1))
		})

		It("forwards Clear instructions to the editor", func() {
			_, err := io.WriteString(inBuf, "C")
			Expect(err).NotTo(HaveOccurred())

			r.ProcessEditActions()

			Expect(fakeImageEditor.ClearCallCount()).To(Equal(1))
		})

		It("upcases the command action and char", func() {
			_, err := io.WriteString(inBuf, "l 1 3 a")
			Expect(err).NotTo(HaveOccurred())

			r.ProcessEditActions()

			Expect(fakeImageEditor.SetCallCount()).To(Equal(1))
			x, y, char := fakeImageEditor.SetArgsForCall(0)
			Expect(x).To(Equal(1))
			Expect(y).To(Equal(3))
			Expect(char).To(Equal("A"))
		})

		Context("if any coordinate argument cannot be translated into an integer", func() {
			It("prints the error", func() {
				_, err := io.WriteString(inBuf, "L p 3 A")
				Expect(err).NotTo(HaveOccurred())

				r.ProcessEditActions()
				Expect(outBuf).To(gbytes.Say("could not parse non-integer 'p'"))
			})
		})
	})
})
