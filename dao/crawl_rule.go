package dao

import "jobcrawler/model"

type ruleRepository struct {
}

var DefaultRuleRepository  = ruleRepository{}

func (rep ruleRepository) GetList() ([]model.CrawlRule,error){
	all := make([]model.CrawlRule, 0)
	err := dbEngine.Find(&all)
	return all, err
}

func (rep ruleRepository)GetModel(Id int) (*model.CrawlRule, bool)  {
	model := new (model.CrawlRule)
	result,err := dbEngine.ID(Id).Get(model)
	if(err != nil){
		panic(err)
	}

	return  model, result
}
