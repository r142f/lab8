package events_statistic

import (
	"fmt"
	"lab8/clock"
	"time"
)

type IEventsStatistic interface {
	IncEvent(name string)
	GetEventStatisticByName(name string) int
	GetAllEventsStatistic() map[string]int
	PrintStatistic()
}

type EventsStatistic struct {
	Statistic map[string][]time.Time
	Clock     clock.Clock
}

func (es *EventsStatistic) removeOldStatistics() {
	for name, statistic := range es.Statistic {
		curFrom := 0
		for statistic[len(statistic)-1].Sub(statistic[curFrom]) >= time.Hour {
			curFrom++
		}
		
		es.Statistic[name] = es.Statistic[name][curFrom:]
	}
}

func (es *EventsStatistic) IncEvent(name string) {
	_, ok := es.Statistic[name]
	if ok {
		es.Statistic[name] = append(es.Statistic[name], es.Clock.Now())
	} else {
		es.Statistic[name] = []time.Time{es.Clock.Now()}
	}

	es.removeOldStatistics()
}

func (es *EventsStatistic) GetEventStatisticByName(name string) int {
	es.removeOldStatistics()

	return len(es.Statistic[name])
}

func (es *EventsStatistic) GetAllEventsStatistic() map[string]int {
	es.removeOldStatistics()

	statistic := make(map[string]int)
	for name := range es.Statistic {
		statistic[name] = es.GetEventStatisticByName(name)
	}

	return statistic
}

func (es *EventsStatistic) PrintStatistic() {
	es.removeOldStatistics()

	for name, count := range es.GetAllEventsStatistic() {
		fmt.Println(name, count)
	}
}
