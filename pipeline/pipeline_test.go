package pipeline

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const channelLength = 10
const writeNumber = 500

func TestPipeline(t *testing.T) {
	t.Run("pipeline", func(t *testing.T) {
		in := make(chan string, 100)
		out1 := make(chan string)
		out2 := make(chan string)

		go step1(in, out1)
		go step2(out1, out2)
		out3 := step3(out2)

		data := `
	Это текст с    лишними пробелами.
    Он содержит.    несколько предложений.
	в том числе без заглавных букв.
    `
		for _, line := range strings.Split(data, "\n") {
			in <- line
		}
		close(in)

		result := <-out3
		require.Equal(t, "Это текст с лишними пробелами", result)
		result = <-out3
		require.Equal(t, "Он содержит", result)
		result = <-out3
		require.Equal(t, "Несколько предложений", result)
		result = <-out3
		require.Equal(t, "В том числе без заглавных букв", result)
	})
}
