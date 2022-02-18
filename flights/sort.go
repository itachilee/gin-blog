package flights

import (
	"fmt"
	"sort"
)

// 结构体定义
type test struct {
	value int
	str   string
}

func Sort() {

	s := make([]Flight, 6)
	s[0] = Flight{
		Flight:          "GD2501",
		Scheduled:       "08:00",
		Departure:       sian,
		Arrival:         chengdu,
		WeekdaysRegular: 1100,
		WeekdaysReward:  800,
		WeekendsRegular: 900,
		WeekendsReward:  500,
	}
	s[1] = Flight{
		Flight:          "GD2506",
		Scheduled:       "12:25",
		Departure:       sian,
		Arrival:         chengdu,
		WeekdaysRegular: 1600,
		WeekdaysReward:  1100,
		WeekendsRegular: 600,
		WeekendsReward:  500,
	}
	s[2] = Flight{
		Flight:          "GD8732",
		Scheduled:       "19:30",
		Departure:       sian,
		Arrival:         chengdu,
		WeekdaysRegular: 2200,
		WeekdaysReward:  1000,
		WeekendsRegular: 1500,
		WeekendsReward:  400,
	}
	s[3] = Flight{
		Flight:          "GD2502",
		Scheduled:       "12:00",
		Departure:       chengdu,
		Arrival:         sian,
		WeekdaysRegular: 1700,
		WeekdaysReward:  800,
		WeekendsRegular: 900,
		WeekendsReward:  800,
	}
	s[4] = Flight{
		Flight:          "GD2607",
		Scheduled:       "16:25",
		Departure:       chengdu,
		Arrival:         sian,
		WeekdaysRegular: 1600,
		WeekdaysReward:  1100,
		WeekendsRegular: 600,
		WeekendsReward:  500,
	}
	s[5] = Flight{
		Flight:          "GD8733",
		Scheduled:       "16:00",
		Departure:       chengdu,
		Arrival:         sian,
		WeekdaysRegular: 1600,
		WeekdaysReward:  1500,
		WeekendsRegular: 1000,
		WeekendsReward:  400,
	}

	fmt.Println("初始化结果:")
	fmt.Println(s)

	// 从小到大排序(稳定排序)
	sort.SliceStable(s, func(i, j int) bool {
		if s[i].WeekdaysRegular < s[j].WeekdaysRegular {
			return true
		}
		return false
	})
	fmt.Println("\n从小到大排序结果:")
	fmt.Println(s)
}
