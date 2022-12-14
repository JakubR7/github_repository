package github

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "a golang introduction repository",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}
	//Marshal takes an input interface and attempts to create a valid json string
	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t,
		`{"name":"golang introduction","description":"a golang introduction repository","homepage":"https://github.com","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`,
		string(bytes))

	var target CreateRepoRequest

	//Unmarshal takes an input byte array and a *pointer* that we're trying to fill using this json.
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
	fmt.Println(target)

}
