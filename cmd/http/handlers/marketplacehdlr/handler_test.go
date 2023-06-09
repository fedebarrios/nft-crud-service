package marketplacehdlr

import (
	"github.com/golang/mock/gomock"
	"nft-crud-service/internal/application/marketplacesrvs"
	mocks "nft-crud-service/mocks/mockgen"
	"testing"
)

type dependencies struct {
	marketplaceRepository *mocks.MockMarketplaceRepository
}

type errResponse struct {
	Messages []string `json:"messages"`
	Code     string   `json:"code"`
}

func makeDependencies(t *testing.T) *dependencies {
	return &dependencies{
		marketplaceRepository: mocks.NewMockMarketplaceRepository(gomock.NewController(t)),
	}
}

func buildHandler(dep *dependencies) *handler {
	srvs := marketplacesrvs.NewServices(dep.marketplaceRepository)
	return NewMarketplaceHandler(srvs)
}
