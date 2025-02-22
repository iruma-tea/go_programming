package integration

import (
	"context"
	"go-api-arch-clean-template/adapter/controller/gin/presenter"
	"go-api-arch-clean-template/pkg"
	"net/http"
	"testing"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/suite"
)

type AlbumTestSuite struct {
	suite.Suite
}

func TestAlbumSuite(t *testing.T) {
	suite.Run(t, new(AlbumTestSuite))
}

func (suite *AlbumTestSuite) TestAlbumCreateGetDelete() {
	// Create
	baseEndPoint := pkg.GetEndpoint("/app/v1")
	apiClient, _ := presenter.NewClientWithResponses(baseEndPoint)
	createResponse, err := apiClient.CreateAlbumWithResponse(context.Background(), presenter.CreateAlbumJSONRequestBody{
		Title:       "test",
		Category:    presenter.Category{Name: presenter.Sports},
		ReleaseDate: openapi_types.Date{Time: time.Now()},
	})
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusCreated, createResponse.StatusCode())
	suite.Assert().Nil(err)
	suite.Assert().NotNil(createResponse.JSON201.Id)
	suite.Assert().Equal("test", createResponse.JSON201.Title)
	suite.Assert().Equal("sports", string(createResponse.JSON201.Category.Name))
	suite.Assert().NotNil(createResponse.JSON201.ReleaseDate.Date())
}
