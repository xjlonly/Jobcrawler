package model


import "time"

type CrawlRule struct {
	Id           int
	Domain       string
	Name         string
	JobName      string
	JobCompany   string
	JobCity      string
	JobWorkExp   string
	JobSalary    string
	JobEducation string
	JobJd        string
	JobWelfare   string
	CreatedAt    time.Time
}