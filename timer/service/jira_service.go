package service

import (
	"encoding/json"
	"github.com/andygrunwald/go-jira"
	"go_study/config"
	"go_study/logger"
	"io/ioutil"
)

var jiraClient *jira.Client

type Field struct {
	FieldAsHtml   string
	fieldCssClass string
}

type Issue struct {
	Key      string
	Position int
	Fields   []Field
}

type UnSignedSupport struct {
	Displayed int
	Total     int
	Start     int
	End       int
	IssueHash string
	Issues    []Issue
	//Columns               []map[string]string
	IsUsingDefaultSorting bool
}

type JiraService struct {
}

func NewJiraService() *JiraService {
	return &JiraService{}
}

func (c *JiraService) GetJiraClient() (*jira.Client, error) {
	tp := jira.BasicAuthTransport{
		Username: config.Cfg.UserName,
		Password: config.Cfg.Password,
	}
	var err error
	if jiraClient == nil {
		jiraClient, err = jira.NewClient(tp.Client(), config.Cfg.JiraHostUrl)
	}
	if err != nil {
		return nil, err
	}
	return jiraClient, nil
}

func (c *JiraService) respToStr(resp *jira.Response) string {
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		logger.Logger.Error(err)
		return ""
	} else {
		return string(body)
	}
}

func (c *JiraService) Refresh() (res *UnSignedSupport, err error) {
	jiraClient, err := c.GetJiraClient()
	res = &UnSignedSupport{}
	if err != nil {
		return nil, err
	}
	req, _ := jiraClient.NewRequest("GET", config.Cfg.Url, nil)
	resp, err := jiraClient.Do(req, nil)
	respStr := c.respToStr(resp)
	if err != nil {
		logger.Logger.Error(respStr)
		return nil, err
	}
	logger.Logger.Info(respStr)
	if err := json.Unmarshal([]byte(respStr), res); err != nil {
		logger.Logger.Error(err)
	}
	return res, nil
}

func (c *JiraService) Assign(key, assignUser string) error {
	var (
		issue      *jira.Issue
		jiraClient *jira.Client
		err        error
	)
	if jiraClient, err = c.GetJiraClient(); err != nil {
		logger.Logger.Error(err)
		return err
	}
	if issue, _, err = jiraClient.Issue.Get(key, nil); err != nil {
		logger.Logger.Error(err)
		return err
	}
	issue.Fields.Assignee = &jira.User{
		Name: assignUser,
	}
	if _, _, err = jiraClient.Issue.Update(issue); err != nil {
		logger.Logger.Error(err)
		return err
	}
	return nil
}
