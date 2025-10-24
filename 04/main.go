package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type visit struct {
	Specialization string
	Date           time.Time
}
type PatientNotFoundError struct{}

func (e PatientNotFoundError) Error() string {
	return "patient not found"
}

func findFIO(all map[string][]visit, find string) (string, error) {
	if _, exist := all[find]; !exist {
		return "patient NOT found", PatientNotFoundError{}
	}
	return "patient found", nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	visits := make(map[string][]visit)
	var input string
	var FIO, spec, date string
	for {
		fmt.Scan(&input)
		//Ivanov Ivan Ivanovich /n orthopedist /n 2024-04-13
		if input == "Save" {
			FIO, _ = reader.ReadString('\n')
			FIO = strings.TrimSpace(FIO)
			fmt.Scan(&spec)
			fmt.Scan(&date)

			t, _ := time.Parse("2006-01-02", date)
			visits[FIO] = append(visits[FIO], visit{spec, t})
		}
		//GetHistory /n Ivanov Ivan Ivanovich
		if input == "GetHistory" {
			FIO, _ = reader.ReadString('\n')
			FIO = strings.TrimSpace(FIO)
			_, err := findFIO(visits, FIO)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			for _, value := range visits[FIO] {
				fmt.Println(value.Specialization, "\n", value.Date.Format("2006-01-02"))
			}
		}
		//GetLastVisit /n Ivanov Ivan Ivanovich /n orthopedist
		if input == "GetLastVisit" {
			FIO, _ = reader.ReadString('\n')
			FIO = strings.TrimSpace(FIO)

			_, err := findFIO(visits, FIO)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Scan(&spec)
			first := true
			var last time.Time
			for _, value := range visits[FIO] {
				if value.Specialization == spec && first {
					last = value.Date
					first = false
				}
				if value.Date.After(last) && value.Specialization == spec {
					last = value.Date
				}
			}
			fmt.Println(last.Format("2006-01-02"))
		}
	}
}
