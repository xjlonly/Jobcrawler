package model


import (
	"strconv"
	"time"
)

type JobInfo struct {
	Id        int
	Name      string
	Company   string
	City      string
	Salary    string
	Education string
	WorkExp   string
	Jd        string
	Welfare   string
	Url       string
	CreatedAt time.Time
}

func (j *JobInfo) String() string {
	return "id:" + strconv.Itoa(j.Id) + ";name:" + j.Name
}
