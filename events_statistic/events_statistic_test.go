package events_statistic

import (
	"lab8/clock"
	"lab8/utils"
	"testing"
	"time"
)

func TestEventsStatisticWithNormalClock(t *testing.T) {
	normalClock := &clock.NormalClock{}

	eventsStatistic := EventsStatistic{
		Statistic: make(map[string][]time.Time),
		Clock:     normalClock,
	}

	names := []string{"event1", "event2"}
	for _, name := range names {
		for i := 0; i < 10; i++ {
			eventsStatistic.IncEvent(name)
		}
	}

	for name, count := range eventsStatistic.GetAllEventsStatistic() {
		utils.AssertEqual(count, 10)
		utils.AssertEqual(eventsStatistic.GetEventStatisticByName(name), 10)
	}
}

func TestEventsStatisticWithSetableClock(t *testing.T) {
	setableClock := &clock.SetableClock{}

	eventsStatistic := EventsStatistic{
		Statistic: make(map[string][]time.Time),
		Clock:     setableClock,
	}

	startTime := time.Date(2022, time.February, 24, 0, 0, 0, 0, time.UTC)
	setableClock.SetNow(startTime)

	names := []string{"event1", "event2"}
	for _, name := range names {
		for i := 0; i < 10; i++ {
			eventsStatistic.IncEvent(name)
		}
	}

	setableClock.SetNow(startTime.Add(time.Hour))

	for name, count := range eventsStatistic.GetAllEventsStatistic() {
		utils.AssertEqual(count, 0)
		utils.AssertEqual(eventsStatistic.GetEventStatisticByName(name), 0)
	}
}
