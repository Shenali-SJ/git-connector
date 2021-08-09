package ARF_CreateGitRepo

import util "git-connector/pkg/util"

import (
	"git-connector/pkg/functions"
	"git-connector/pkg/model"
)

type CFGContext struct {
	Requestjsonmodel *model.Requestjsonmodel
	repoData         *model.Createrepodata
	gitToken         string
	ruleReturnString string
	_context         *util.CustomContext
	_ruleConfig      map[string]interface{}
	_returnValue     interface{}
	_errorValue      interface{}
}

func NewCFGContext(pRequestjsonmodel *model.Requestjsonmodel, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Requestjsonmodel: pRequestjsonmodel,
		_context:         pContext,
		_ruleConfig:      pRuleConfig,
	}
}
func ARF_CreateGitRepo(pRequestjsonmodel *model.Requestjsonmodel, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pRequestjsonmodel, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiHybridFunctionNodeM01()
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM01() error {
	var err error
	cfg.ruleReturnString, err = functions.CreateRepo(cfg.gitToken)
	cfg._returnValue = cfg.ruleReturnString
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.repoData = &model.Createrepodata{}
	cfg.gitToken = cfg.Requestjsonmodel.Token

	return nil
}
