package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	strs := strings.Split(data, ",")
	if len(strs) != 3{
		return 0, "", 0, errors.New("wrong split return")
	}
	steps, err := strconv.Atoi(strs[0])
	if err != nil{
		return 0, "", 0, err
	}
	if steps <= 0{
		return 0, "", 0, errors.New("wrong")
	}
	activityTime, err := time.ParseDuration(strs[2])
	if err != nil{
		return 0, "", 0, err
	}
	if activityTime <= 0{
		return 0, "", 0, errors.New("wrong")
	}
	return steps, strs[1], activityTime, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLen := height * stepLengthCoefficient
	return float64(steps) * stepLen / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)
	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	var dist, speed, calories float64
	steps, activityName, activityTime, err := parseTraining(data)
	if err != nil{
		log.Println(err)
		return "", err
	}
	if weight <= 0 || height <= 0 {
		return "", errors.New("have not data")
	}
	switch{
	case activityName == "Ходьба":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, activityTime)
		calories, err = WalkingSpentCalories(steps, weight, height, activityTime)
		if err != nil{
			return "", err
		}
	case activityName == "Бег":
		dist = distance(steps, height)
		speed = meanSpeed(steps, height, activityTime)
		calories, err = RunningSpentCalories(steps, weight, height, activityTime)
		if err != nil{
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activityName, activityTime.Hours(), dist, speed, calories), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0{
		return 0, errors.New("have not data")
	}
	speed := meanSpeed(steps, height, duration)
	return (weight * speed * duration.Minutes()) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0{
		return 0, errors.New("have not data")
	}
	speed := meanSpeed(steps, height, duration)
	return (weight * speed * duration.Minutes()) / minInH * walkingCaloriesCoefficient, nil
}
