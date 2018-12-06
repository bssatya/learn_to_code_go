package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from the user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record Cleo win from the user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}

		cli := &CLI{playerStore, in}
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Cleo")
	})

}
