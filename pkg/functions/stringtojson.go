package functions

import (
	"encoding/json"
	"git-connector/pkg/model"
)

func StringToJson(requeststring string) (*model.Requestjsonmodel, error) {

	var jsonModel *model.Requestjsonmodel
	err := json.Unmarshal([]byte(requeststring), jsonModel)

	return jsonModel, err
}
