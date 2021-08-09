package functions

import (
	"fmt"
	"git-connector/pkg/model"
	"github.com/go-resty/resty/v2"
)

func CreateRepo(gitaccesstoken string) (string, error) {

	repoData := model.Createrepodata{
		Name:              "connector-generated-repo",
		Autoinit:          true,
		Private:           true,
		Gitignoretemplate: "nanoc",
	}

	client := resty.New()
	resp, err := client.R().
		SetHeader("Authorization", gitaccesstoken).
		SetBody(repoData).
		Post("https://api.github.com/user/repos")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status())

	return resp.Status(), nil
}
