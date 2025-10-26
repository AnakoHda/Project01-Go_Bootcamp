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

func findFIO(all map[string][]visit, find string) (bool, error) {
	if _, exist := all[find]; !exist {
		return false, PatientNotFoundError{}
	}
	return true, nil
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	visits := make(map[string][]visit)
	var input string
	for {
		input = readLine(reader)

		switch input {
		case "Save":
			/*
				Save
				Ivanov Ivan Ivanovich
				orthopedist
				2024-04-13
			*/
			save(visits, reader)

		case "GetHistory":
			/*
				GetHistory
				Ivanov Ivan Ivanovich
			*/
			if err := getHistory(visits, reader); err != nil {
				continue
			}

		case "GetLastVisit":
			/*
				GetLastVisit
				Ivanov Ivan Ivanovich
				orthopedist
			*/
			if err := getLastVisit(visits, reader); err != nil {
				continue
			}

		}
	}
}

func readLine(reader *bufio.Reader) string {
	str, _ := reader.ReadString('\n')
	return strings.TrimSpace(str)
}

func save(visits map[string][]visit, reader *bufio.Reader) {
	FIO := readLine(reader)
	spec := readLine(reader)
	date := readLine(reader)

	t, _ := time.Parse("2006-01-02", date)
	visits[FIO] = append(visits[FIO], visit{spec, t})
}
func getHistory(visits map[string][]visit, reader *bufio.Reader) error {
	FIO := readLine(reader)

	if _, err := findFIO(visits, FIO); err != nil {
		fmt.Println(err.Error())
		return err
	}

	for _, value := range visits[FIO] {
		fmt.Printf("%s \n%s\n", value.Specialization, value.Date.Format("2006-01-02"))
	}
	return nil
}
func getLastVisit(visits map[string][]visit, reader *bufio.Reader) error {
	FIO := readLine(reader)

	if _, err := findFIO(visits, FIO); err != nil {
		fmt.Println(err.Error())
		return err
	}

	spec := readLine(reader)

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
	return nil
}
