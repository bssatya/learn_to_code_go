package poker_test

import (
	"strings"
	"testing"
	"time"

	"github.com/bssatya/learn_to_code_go/learn_go_with_tests/http_server"
)

type SpyBlindAlerter struct {
	alerts []struct {
		scheduledAt time.Duration
		ammount     int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, ammount int) {
	s.alerts = append(s.alerts, struct {
		scheduledAt time.Duration
		ammount     int
	}{duration, ammount})
}

var dummySpyAlerter = &SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	t.Run("record Chris win from the user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record Cleo win from the user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(blindAlerter.alerts) != 1 {
			t.Fatalf("expected a blind alert to be scheduled")
		}
	})

}
