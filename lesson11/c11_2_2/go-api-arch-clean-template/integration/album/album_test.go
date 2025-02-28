package integration

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"github.com/stretchr/testify/suite"

	"go-api-arch-clean-template/adapter/controller/gin/presenter"
	"go-api-arch-clean-template/api"
	"go-api-arch-clean-template/pkg"
)

type AlbumTestSuite struct {
	suite.Suite
}

func TestAlbumSuite(t *testing.T) {
	suite.Run(t, new(AlbumTestSuite))
}

func (suite *AlbumTestSuite) TestAlbumCreateGetDelete() {
	// Create
	baseEndpoint := pkg.GetEndpoint(fmt.Sprintf("/api/%s", api.Version))
	apiClient, _ := presenter.NewClientWithResponses(baseEndpoint)
	kindAlbum := "album"
	createResponse, err := apiClient.CreateAlbumWithResponse(context.Background(), presenter.CreateAlbumJSONRequestBody{
		Kind:        &kindAlbum,
		Title:       "test",
		Category:    presenter.Category{Name: presenter.Sports},
		ReleaseDate: openapi_types.Date{Time: time.Now()},
	})
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusCreated, createResponse.StatusCode())
	suite.Assert().Nil(err)
	suite.Assert().NotNil(createResponse.JSON201.Data.Id)
	suite.Assert().Equal("album", createResponse.JSON201.Data.Kind)
	suite.Assert().Equal("test", createResponse.JSON201.Data.Title)
	suite.Assert().Equal("sports", string(createResponse.JSON201.Data.Category.Name))
	suite.Assert().NotNil(createResponse.JSON201.Data.ReleaseDate)

	// Get
	getResponse, err := apiClient.GetAlbumByIdWithResponse(context.Background(), createResponse.JSON201.Data.Id)
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusOK, getResponse.StatusCode())
	suite.Assert().Nil(err)
	suite.Assert().Equal(createResponse.JSON201.Data.Id, getResponse.JSON200.Data.Id)
	suite.Assert().Equal("album", getResponse.JSON200.Data.Kind)
	suite.Assert().Equal("test", getResponse.JSON200.Data.Title)
	suite.Assert().Equal("sports", string(getResponse.JSON200.Data.Category.Name))
	suite.Assert().NotNil(getResponse.JSON200.Data.ReleaseDate)

	// Update
	title := "updated"
	category := presenter.Category{
		Name: presenter.Food,
	}
	updateResponse, err := apiClient.UpdateAlbumByIdWithResponse(context.Background(), getResponse.JSON200.Data.Id, presenter.UpdateAlbumByIdJSONRequestBody{
		Title:    &title,
		Category: &category,
	})
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusOK, updateResponse.StatusCode())
	suite.Assert().Nil(err)
	suite.Assert().Equal("album", updateResponse.JSON200.Data.Kind)
	suite.Assert().Equal("updated", updateResponse.JSON200.Data.Title)
	suite.Assert().Equal("food", string(updateResponse.JSON200.Data.Category.Name))
	suite.Assert().NotNil(updateResponse.JSON200.Data.ReleaseDate)

	// Delete
	deleteResponse, err := apiClient.DeleteAlbumByIdWithResponse(context.Background(), updateResponse.JSON200.Data.Id)
	suite.Assert().Nil(err)
	suite.Assert().Equal(http.StatusNoContent, deleteResponse.StatusCode())
}
