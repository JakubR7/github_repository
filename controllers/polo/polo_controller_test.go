package polo

import (
	"golang_mvc/git_repo/src/api/utils/test_utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstatns(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestMarco(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/marco", nil)
	c := test_utils.GetMockedContext(request, response)

	Marco(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())

}
