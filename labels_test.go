package logentries_goclient

import (
	"fmt"
	"github.com/dikhan/http_goclient/testutils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestLabels_GetLabels(t *testing.T) {

	expectedLabels := []Label{
		{
			Id:       "label-uuid",
			Name:     "Login Failure",
			Reserved: false,
			Color:    "007afb",
			SN:       1056,
		},
	}

	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, "/management/labels", nil, http.StatusOK, &labelsCollection{expectedLabels})
	labels := getLabelsClient(requestMatcher)

	returnedLabels, err := labels.GetLabels()
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(expectedLabels, returnedLabels))
}

func TestTags_GetLabel(t *testing.T) {

	expectedLabel := Label{
		Id:       "label-uuid",
		Name:     "Login Failure",
		Reserved: false,
		Color:    "007afb",
		SN:       1056,
	}

	url := fmt.Sprintf("/management/labels/%s", expectedLabel.Id)
	requestMatcher := testutils.NewRequestMatcher(http.MethodGet, url, nil, http.StatusOK, &getLabel{expectedLabel})

	labels := getLabelsClient(requestMatcher)

	returnedLabel, err := labels.GetLabel(expectedLabel.Id)
	assert.Nil(t, err)
	assert.EqualValues(t, expectedLabel, returnedLabel)

}

func TestTags_GetLabelErrorsIfTagIdIsEmpty(t *testing.T) {
	labels := Labels{nil}
	_, err := labels.GetLabel("")
	assert.NotNil(t, err)
	assert.Error(t, err, "labelId input parameter is mandatory")
}

func getLabelsClient(requestMatcher testutils.TestRequestMatcher) Labels {
	c := getTestClient(requestMatcher)
	return newLabels(c)
}
