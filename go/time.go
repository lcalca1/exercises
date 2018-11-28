package main

import "time"
import "fmt"

var t time.Time = time.Now()
var week time.Duration

func main() {
	fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t, t.Location(), t)
	fmt.Printf("时间: %v, 时区:  %v,  时间类型: %T\n", t.UTC(), t.UTC().Location(), t)
	t_by_unix := time.Unix(1487780010, 0)
	fmt.Println("'time.Unix': ", t_by_unix)

	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)

	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	fmt.Println(t)
	s := t.Format("20050102")
	fmt.Println(t, "=>", s)
}
