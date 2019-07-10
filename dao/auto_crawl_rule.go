package dao

import "jobcrawler/model"

type autoRuleRepository struct {
}

var DefaultAutoRuleRepository  = autoRuleRepository{}

func (rep autoRuleRepository) GetList() ([]model.AutoCrawlRule,error){
	all := make([]model.AutoCrawlRule, 0)
	err := dbEngine.Find(&all)
	return all, err
}

func (rep autoRuleRepository)GetModel(Id int) (*model.AutoCrawlRule, bool)  {
	model := new (model.AutoCrawlRule)
	result,err := dbEngine.ID(Id).Get(model)
	if(err != nil){
		panic(err)
	}

	return  model, result
}
