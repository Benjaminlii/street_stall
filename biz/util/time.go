package util

import "time"

// GetYesterdayTime 获取昨天的一个时间
func GetYesterdayTime() time.Time {
	nowTime := time.Now()
	yesTime := nowTime.AddDate(0, 0, -1)
	return yesTime
}

// GetTodayFirstSecond 获取当天第一秒的time对象
func GetTodayFirstSecond() time.Time {
	currentTime := time.Now()
	todayFirstSecond := time.Date(
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		0,
		0,
		0,
		0,
		currentTime.Location(),
	)
	return todayFirstSecond
}

// GetTodayLastSecond 获取当天最后一秒的time对象
func GetTodayLastSecond() time.Time {
	currentTime := time.Now()
	todayLastSecond := time.Date(
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		23,
		59,
		59,
		0,
		currentTime.Location(),
	)
	return todayLastSecond
}

// GetTodayFirstSecond 获取某一天第一秒的time对象
func GetOneDayFirstSecond(oneDay time.Time) time.Time {
	firstSecond := time.Date(
		oneDay.Year(),
		oneDay.Month(),
		oneDay.Day(),
		0,
		0,
		0,
		0,
		oneDay.Location(),
	)
	return firstSecond
}

// GetOneDayLastSecond 获取某一天最后一秒的time对象
func GetOneDayLastSecond(oneDay time.Time) time.Time {
	lastSecond := time.Date(
		oneDay.Year(),
		oneDay.Month(),
		oneDay.Day(),
		23,
		59,
		59,
		0,
		oneDay.Location(),
	)
	return lastSecond
}
