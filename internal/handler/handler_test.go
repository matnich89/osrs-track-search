package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"osrs-track-search/internal/client"
	"osrs-track-search/internal/client/mocks"
	"osrs-track-search/internal/process"
	"osrs-track-search/model"
	"testing"
)

const happyResponse = "7215,1475,25505097\n10117,63,388102\n9482,61,309518\n12656,64,416412\n12100,66,508922\n14136,60,274719\n7889,53,138457\n9357,69,709323\n4787,80,1991640\n13249,65,452814\n10324,61,328597\n6599,82,2433319\n12425,74,1143541\n11296,56,195885\n8456,58,227931\n16487,52,127762\n5649,60,283914\n10619,70,803067\n2200,99,13403954\n7654,56,199804\n7575,65,471485\n9028,47,78204\n11619,48,85098\n6994,66,532629\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n4182,143\n1614,101\n4043,34\n6429,8\n-1,-1\n-1,-1\n-1,-1\n13249,500\n-1,-1\n-1,-1\n5595,20\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n3116,102\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n-1,-1\n8301,74\n-1,-1\n-1,-1\n"

func TestSearchIronman(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := mocks.NewMockClient(mockCtrl)
	processor := process.NewProcessor()

	successResponseFile, err := os.Open("./testData/expectedSuccessResponse.json")

	assert.NoError(t, err)

	defer successResponseFile.Close()

	fileBytes, err := io.ReadAll(successResponseFile)

	assert.NoError(t, err)

	tests := []struct {
		name           string
		character      string
		responseFunc   func()
		expectedStatus int
		expectedBody   []byte
	}{
		{
			name:      "success",
			character: "validCharacter",
			responseFunc: func() {
				mockClient.EXPECT().GetHighScores(gomock.Eq("validCharacter"), gomock.Eq(client.Ironman)).Return(&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(happyResponse)),
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   fileBytes,
		},
		{
			name:      "not found",
			character: "notFound",
			responseFunc: func() {
				mockClient.EXPECT().GetHighScores(gomock.Eq("notFound"), gomock.Eq(client.Ironman)).Return(nil, client.ErrNotFound)
			},
		},
		{
			name:      "other error",
			character: "otherError",
			responseFunc: func() {
				mockClient.EXPECT().GetHighScores(gomock.Eq("otherError"), gomock.Eq(client.Ironman)).Return(nil, errors.New("some bad error"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.responseFunc != nil {
				tt.responseFunc()
			}

			req, _ := http.NewRequest("GET", "/search_ironman?character="+tt.character, nil)
			rr := httptest.NewRecorder()

			handler := NewHandler(mockClient)
			handler.processor = processor

			handler.SearchIronman(rr, req)

			if tt.expectedBody != nil {
				var actualStats model.CharStats
				var expectedStats model.CharStats
				json.Unmarshal(tt.expectedBody, &expectedStats)
				json.Unmarshal(rr.Body.Bytes(), &actualStats)

				assert.Equal(t, expectedStats, actualStats)
			}

		})
	}

}
