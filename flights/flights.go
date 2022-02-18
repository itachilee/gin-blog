package flights

import (
	"log"
	"time"
)

type Flight struct {
	Flight          string
	Scheduled       string
	Departure       string
	Arrival         string
	WeekdaysRegular int
	WeekdaysReward  int
	WeekendsRegular int
	WeekendsReward  int
}

const (
	sian    = "Xi’an"
	chengdu = "Chengdu"
)

var flights [6]Flight = [6]Flight{
	{
		Flight:          "GD2501",
		Scheduled:       "08:00",
		Departure:       sian,
		Arrival:         chengdu,
		WeekdaysRegular: 1100,
		WeekdaysReward:  800,
		WeekendsRegular: 900,
		WeekendsReward:  500,
	},

	{
		Flight:          "GD2506",
		Scheduled:       "12:25",
		Departure:       sian,
		Arrival:         chengdu,
		WeekdaysRegular: 1600,
		WeekdaysReward:  1100,
		WeekendsRegular: 600,
		WeekendsReward:  500,
	},

	{
		Flight:          "GD8732",
		Scheduled:       "19:30",
		Departure:       sian,
		Arrival:         chengdu,
		WeekdaysRegular: 2200,
		WeekdaysReward:  1000,
		WeekendsRegular: 1500,
		WeekendsReward:  400,
	},

	{
		Flight:          "GD2502",
		Scheduled:       "12:00",
		Departure:       chengdu,
		Arrival:         sian,
		WeekdaysRegular: 1700,
		WeekdaysReward:  800,
		WeekendsRegular: 900,
		WeekendsReward:  800,
	},

	{
		Flight:          "GD2607",
		Scheduled:       "16:25",
		Departure:       chengdu,
		Arrival:         sian,
		WeekdaysRegular: 1600,
		WeekdaysReward:  1100,
		WeekendsRegular: 600,
		WeekendsReward:  500,
	},

	{
		Flight:          "GD8733",
		Scheduled:       "16:00",
		Departure:       chengdu,
		Arrival:         sian,
		WeekdaysRegular: 1600,
		WeekdaysReward:  1500,
		WeekendsRegular: 1000,
		WeekendsReward:  400,
	},
}

const (
	REWARD  string = "REWARD"
	REGULAR string = "REGULAR"
)

const layout string = "20060102"

// 获取机票包
func GetFlightsPackge(customerType, departureDate, arrivalDate string) []Flight {
	switch customerType {
	case REWARD:
		dd, err := timePrase(departureDate)
		if err != nil {
			log.Fatalf("DepartureDate err:%v", err)
		}
		ad, err := timePrase(departureDate)
		if err != nil {
			log.Fatalf("ArrivalDate err:%v", err)
		}
		ddt := dd.Weekday()
	case REGULAR:
	default:
		log.Fatalln("数据的[CUSTOMER_TYPE]不正确,请检查输入")
	}
	return nil
}

func timePrase(input string) (time.Time, error) {
	inputStr := input[0 : len(input)-3]
	t, err := time.Parse(layout, inputStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil

}

func GetFlights(customerType string, t time.Weekday) *Flight {

	return nil
}
