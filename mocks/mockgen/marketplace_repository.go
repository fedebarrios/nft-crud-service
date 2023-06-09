// Code generated by MockGen. DO NOT EDIT.
// Source: ./repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	domain "ssr/internal/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockMarketplaceRepository is a mock of MarketplaceRepository interface.
type MockMarketplaceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMarketplaceRepositoryMockRecorder
}

// MockMarketplaceRepositoryMockRecorder is the mock recorder for MockMarketplaceRepository.
type MockMarketplaceRepositoryMockRecorder struct {
	mock *MockMarketplaceRepository
}

// NewMockMarketplaceRepository creates a new mock instance.
func NewMockMarketplaceRepository(ctrl *gomock.Controller) *MockMarketplaceRepository {
	mock := &MockMarketplaceRepository{ctrl: ctrl}
	mock.recorder = &MockMarketplaceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMarketplaceRepository) EXPECT() *MockMarketplaceRepositoryMockRecorder {
	return m.recorder
}

// FindNfts mocks base method.
func (m *MockMarketplaceRepository) FindNfts(ctx context.Context, next, took *int) (*domain.NftList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindNfts", ctx, next, took)
	ret0, _ := ret[0].(*domain.NftList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindNfts indicates an expected call of FindNfts.
func (mr *MockMarketplaceRepositoryMockRecorder) FindNfts(ctx, next, took interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindNfts", reflect.TypeOf((*MockMarketplaceRepository)(nil).FindNfts), ctx, next, took)
}

// GetNftById mocks base method.
func (m *MockMarketplaceRepository) GetNftById(ctx context.Context, nftId string) (*domain.Nft, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftById", ctx, nftId)
	ret0, _ := ret[0].(*domain.Nft)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNftById indicates an expected call of GetNftById.
func (mr *MockMarketplaceRepositoryMockRecorder) GetNftById(ctx, nftId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftById", reflect.TypeOf((*MockMarketplaceRepository)(nil).GetNftById), ctx, nftId)
}

// GetUserById mocks base method.
func (m *MockMarketplaceRepository) GetUserById(ctx context.Context, userId string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, userId)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockMarketplaceRepositoryMockRecorder) GetUserById(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockMarketplaceRepository)(nil).GetUserById), ctx, userId)
}

// SaveNft mocks base method.
func (m *MockMarketplaceRepository) SaveNft(ctx context.Context, nft domain.Nft) (*domain.Nft, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNft", ctx, nft)
	ret0, _ := ret[0].(*domain.Nft)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveNft indicates an expected call of SaveNft.
func (mr *MockMarketplaceRepositoryMockRecorder) SaveNft(ctx, nft interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNft", reflect.TypeOf((*MockMarketplaceRepository)(nil).SaveNft), ctx, nft)
}

// UpdateUsersBalances mocks base method.
func (m *MockMarketplaceRepository) UpdateUsersBalances(ctx context.Context, balances map[string]float64) ([]domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUsersBalances", ctx, balances)
	ret0, _ := ret[0].([]domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUsersBalances indicates an expected call of UpdateUsersBalances.
func (mr *MockMarketplaceRepositoryMockRecorder) UpdateUsersBalances(ctx, balances interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUsersBalances", reflect.TypeOf((*MockMarketplaceRepository)(nil).UpdateUsersBalances), ctx, balances)
}
