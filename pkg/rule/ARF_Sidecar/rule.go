package ARF_Sidecar

import util "git-connector/pkg/util"

import (
	"git-connector/pkg/dto"
	"git-connector/pkg/functions"
	"git-connector/pkg/model"
	"git-connector/pkg/rule/ARF_CreateGitRepo"
)

type CFGContext struct {
	Request      *dto.Request
	request      string
	requestJson  *model.Requestjsonmodel
	returnValue  string
	_context     *util.CustomContext
	_ruleConfig  map[string]interface{}
	_returnValue interface{}
	_errorValue  interface{}
}

func NewCFGContext(pRequest *dto.Request, pContext *util.CustomContext, pRuleConfig map[string]interface{}) *CFGContext {
	return &CFGContext{

		Request:     pRequest,
		_context:    pContext,
		_ruleConfig: pRuleConfig,
	}
}
func ARF_Sidecar(pRequest *dto.Request, pContext *util.CustomContext, pRuleConfig map[string]interface{}) interface{} {

	cfg := NewCFGContext(pRequest, pContext, pRuleConfig)
	cfg.r0()
	return cfg._returnValue
}
func (cfg *CFGContext) r0() error {

	cfg.xiModelPropertyNodeM0()

	cfg.xiHybridFunctionNodeM01()

	cfg.xiSubRuleNodeM12()
	return nil
}

func (cfg *CFGContext) xiSubRuleNodeM12() error {
	cfg.returnValue = ARF_CreateGitRepo.ARF_CreateGitRepo(cfg.requestJson, cfg._context, cfg._ruleConfig).(string)
	cfg._returnValue = cfg.returnValue
	return nil
}

func (cfg *CFGContext) xiHybridFunctionNodeM01() error {
	var err error
	cfg.requestJson, err = functions.StringToJson(cfg.request)
	cfg._returnValue = cfg.requestJson
	return err
}

func (cfg *CFGContext) xiModelPropertyNodeM0() error {
	cfg.request = cfg.Request.Data
	cfg.requestJson = &model.Requestjsonmodel{}

	return nil
}
