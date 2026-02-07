package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	strs := strings.Split(data, ",")
	if len(strs) != 2 {
		return 0, 0, errors.New("wrong split func")
	}
	steps, err := strconv.Atoi(strs[0])
	if err != nil{
		log.Println(err)
		return 0, 0, err
	}
	if steps <= 0{
		return 0, 0, errors.New("have not steps")
	}
	activityTime, err := time.ParseDuration(strs[1])
	if err != nil{
		log.Println(err)
		return 0, 0, err
	}
	if activityTime <= 0{
		return 0, 0, errors.New("time = 0")
	}
	return steps, activityTime, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, activityTime, err := parsePackage(data)
	if err != nil{
		log.Println(err)
		fmt.Println(err)
		return ""
	}
	//сделана проверка в parsePackage, но по тз надо еще раз
	if steps <= 0{
		return ""
	}
	distanсe := float64(steps) * stepLength / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, activityTime)
	if err != nil{
		log.Println(err)
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distanсe, calories)
}
