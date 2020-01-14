package integration_test

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os/exec"
	"testing"

	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var cliBin string

func BeforeSuite() {
	var err error
	cliBin, err = gexec.Build("github.com/mo-work/go-technical-test-for-claudia/cmd", "-mod=vendor")
	if err != nil {
		panic(err)
	}
}

func AfterSuite() {
	gexec.Terminate()
	gexec.CleanupBuildArtifacts()
}

func TestIntegration(t *testing.T) {
	BeforeSuite()
	defer AfterSuite()

	BeforeEach := func() (inBuf *gbytes.Buffer, cliCmd *exec.Cmd) {
		cliCmd = exec.Command(cliBin)
		inBuf = gbytes.NewBuffer()
		cliCmd.Stdin = inBuf
		return inBuf, cliCmd
	}

	t.Run("'I': setting image size", func(t *testing.T) {
		t.Run("when executing the program", func(t *testing.T) {
			t.Run("the user can enter a image size", func(t *testing.T) {
				inBuf, cliCmd := BeforeEach()

				_, err := io.WriteString(inBuf, "I 5 5")
				assert.NoError(t, err)

				bytes, err := cliCmd.CombinedOutput()
				assert.NoError(t, err)
				assert.Len(t, bytes, 0)
			})
		})

		t.Run("if the user specifies something other than a image size", func(t *testing.T) {
			// TODO: this can probably be changed to repeat the request
			t.Run("the program will complain and exit", func(t *testing.T) {
				inBuf, cliCmd := BeforeEach()

				_, err := io.WriteString(inBuf, "X 5 5")
				assert.NoError(t, err)

				bytes, err := cliCmd.CombinedOutput()
				assert.NoError(t, err)
				assert.Equal(t, "invalid image value: unrecognised command 'X', use 'I' for Image initialisation\n", string(bytes))
			})
		})
	})

	t.Run("'L': colouring an individual pixel", func(t *testing.T) {
		t.Run("sets the pixel at the given coordinates to a given colour", func(t *testing.T) {
			inBuf, cliCmd := BeforeEach()

			_, err := io.WriteString(inBuf, "I 5 5\nL 1 3 A\nS")
			assert.NoError(t, err)

			bytes, err := cliCmd.CombinedOutput()
			assert.NoError(t, err)
			assert.Equal(t, "OOOOO\nOOOOO\nAOOOO\nOOOOO\nOOOOO\n\n", string(bytes))
		})

		t.Run("if the action cannot be processed", func(t *testing.T) {
			t.Run("prints an error", func(t *testing.T) {
				inBuf, cliCmd := BeforeEach()

				_, err := io.WriteString(inBuf, "I 5 5\nL 1 6 A")
				assert.NoError(t, err)

				bytes, err := cliCmd.CombinedOutput()
				assert.NoError(t, err)
				assert.Equal(t, "given coordinate is beyond image grid\n", string(bytes))
			})
		})
	})

	t.Run("'V': colouring a vertical line", func(t *testing.T) {
		t.Run("sets the pixels between the specified coordinates", func(t *testing.T) {
			inBuf, cliCmd := BeforeEach()

			_, err := io.WriteString(inBuf, "I 5 5\nV 2 3 5 W\nS")
			assert.NoError(t, err)

			bytes, err := cliCmd.CombinedOutput()
			assert.NoError(t, err)
			assert.Equal(t, "OOOOO\nOOOOO\nOWOOO\nOWOOO\nOWOOO\n\n", string(bytes))
		})

		t.Run("if the action cannot be processed", func(t *testing.T) {
			t.Run("prints an error", func(t *testing.T) {
				inBuf, cliCmd := BeforeEach()

				_, err := io.WriteString(inBuf, "I 5 5\nV 7 2 5 W")
				assert.NoError(t, err)

				bytes, err := cliCmd.CombinedOutput()
				assert.NoError(t, err)
				assert.Equal(t, "given coordinate is beyond image grid\n", string(bytes))
			})
		})
	})

	t.Run("'H': colouring a horizontal line", func(t *testing.T) {
		t.Run("sets the pixels between the specified coordinates", func(t *testing.T) {
			inBuf, cliCmd := BeforeEach()

			_, err := io.WriteString(inBuf, "I 5 5\nH 3 5 2 Z\nS")
			assert.NoError(t, err)

			bytes, err := cliCmd.CombinedOutput()
			assert.NoError(t, err)
			assert.Equal(t, "OOOOO\nOOZZZ\nOOOOO\nOOOOO\nOOOOO\n\n", string(bytes))
		})

		t.Run("if the action cannot be processed", func(t *testing.T) {
			t.Run("prints an error", func(t *testing.T) {
				inBuf, cliCmd := BeforeEach()

				_, err := io.WriteString(inBuf, "I 5 5\nH 3 5 7 Z")
				assert.NoError(t, err)

				bytes, err := cliCmd.CombinedOutput()
				assert.NoError(t, err)
				assert.Equal(t, "given coordinate is beyond image grid\n", string(bytes))
			})
		})
	})

	t.Run("'S': showing the image", func(t *testing.T) {
		t.Run("whenever the user inputs the 'S' command", func(t *testing.T) {
			t.Run("the image is printed in its current state", func(t *testing.T) {
				inBuf, cliCmd := BeforeEach()

				_, err := io.WriteString(inBuf, "I 5 5\nS")
				assert.NoError(t, err)

				bytes, err := cliCmd.CombinedOutput()
				assert.NoError(t, err)
				assert.Equal(t, "OOOOO\nOOOOO\nOOOOO\nOOOOO\nOOOOO\n\n", string(bytes))
			})
		})
	})

	t.Run("'C': clearing the image", func(t *testing.T) {
		t.Run("the image pixels are cleared", func(t *testing.T) {
			inBuf, cliCmd := BeforeEach()

			_, err := io.WriteString(inBuf, "I 2 2\nL 1 1 A\nS\nC\nS")
			assert.NoError(t, err)

			bytes, err := cliCmd.CombinedOutput()
			assert.NoError(t, err)
			assert.Equal(t, "AO\nOO\n\nOO\nOO\n\n", string(bytes))
		})
	})

	t.Run("any other attempted action", func(t *testing.T) {
		t.Run("complains", func(t *testing.T) {
			inBuf, cliCmd := BeforeEach()

			_, err := io.WriteString(inBuf, "I 2 2\nP 1 1 A\n")
			assert.NoError(t, err)

			bytes, err := cliCmd.CombinedOutput()
			assert.NoError(t, err)
			assert.Equal(t, "invalid action\n", string(bytes))
		})
	})
}
