package integration_test

import (
	"io"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Integration", func() {
	var (
		inBuf  *gbytes.Buffer
		cliCmd *exec.Cmd
	)

	BeforeEach(func() {
		cliCmd = exec.Command(cliBin)
		inBuf = gbytes.NewBuffer()
		cliCmd.Stdin = inBuf
	})

	Describe("'I': setting image size", func() {
		Context("when executing the program", func() {
			It("the user can enter a image size", func() {
				_, err := io.WriteString(inBuf, "I 5 5")
				Expect(err).NotTo(HaveOccurred())

				session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session.Err.Contents).Should(HaveLen(0))
			})
		})

		Context("if the user specifies something other than a image size", func() {
			// TODO: this can probably be changed to repeat the request
			It("the program will complain and exit", func() {
				_, err := io.WriteString(inBuf, "X 5 5")
				Expect(err).NotTo(HaveOccurred())

				session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("invalid image value: unrecognised command 'X', use 'I' for Image initialisation"))
			})
		})
	})

	Describe("'L': colouring an individual pixel", func() {
		It("sets the pixel at the given coordinates to a given colour", func() {
			_, err := io.WriteString(inBuf, "I 5 5\nL 1 3 A\nS")
			Expect(err).NotTo(HaveOccurred())

			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("O O O O O\nO O O O O\nA O O O O\nO O O O O\nO O O O O\n"))
		})

		XContext("if the action cannot be processed", func() {
			It("prints an error", func() {
				_, err := io.WriteString(inBuf, "I 5 5\nL 1 6 A")
				Expect(err).NotTo(HaveOccurred())

				session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("given coordinate is beyond image grid"))
			})
		})
	})

	Describe("'S': showing the image", func() {
		Context("whenever the user inputs the 'S' command", func() {
			It("the image is printed in its current state", func() {
				_, err := io.WriteString(inBuf, "I 5 5\nS")
				Expect(err).NotTo(HaveOccurred())

				session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session.Out).Should(gbytes.Say("O O O O O\nO O O O O\nO O O O O\nO O O O O\nO O O O O\n"))
			})
		})
	})

	Describe("'C': clearing the image", func() {
		It("the image pixels are cleared", func() {
			_, err := io.WriteString(inBuf, "I 2 2\nL 1 1 A\nS\nC\nS")
			Expect(err).NotTo(HaveOccurred())

			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("A O\nO O\n\nO O\nO O\n"))
		})
	})

	Context("any other attempted action", func() {
		It("complains", func() {
			_, err := io.WriteString(inBuf, "I 2 2\nP 1 1 A\n")
			Expect(err).NotTo(HaveOccurred())

			session, err := gexec.Start(cliCmd, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session.Out).Should(gbytes.Say("invalid action"))
		})
	})
})
