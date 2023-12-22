package web

import (
	// "log"
	"strconv"
	"time"

	"github.com/aceberg/ExerciseDiary/internal/models"
)

func generateHeatMap() (heatMap []models.HeatMapData) {
	var heat models.HeatMapData

	w := 52 // weeks to show

	max := time.Now()
	min := max.AddDate(0, 0, -7*w)

	startDate := weekStartDate(min)
	countMap := countHeat()

	dow := []string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}

	for _, day := range dow {

		heat.Y = day

		for i := 0; i < w+1; i++ {

			heat.X = strconv.Itoa(i)
			heat.D = startDate.AddDate(0, 0, 7*i).Format("2006-01-02")
			heat.V = countMap[heat.D]
			heatMap = append(heatMap, heat)
		}

		startDate = startDate.AddDate(0, 0, 1)
	}

	return heatMap
}

func weekStartDate(date time.Time) time.Time {
	offset := (int(time.Monday) - int(date.Weekday()) - 7) % 7
	result := date.Add(time.Duration(offset*24) * time.Hour)
	return result
}

func countHeat() map[string]int {
	countMap := make(map[string]int)

	for _, ex := range exData.Sets {
		reps, exists := countMap[ex.Date]
		if exists {
			countMap[ex.Date] = reps + 1
		} else {
			countMap[ex.Date] = 1
		}
	}

	return countMap
}
