package mock

import (
	"context"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
)

// AccountClientMock is a hand-written mock of the AccountClient gRPC client.
type AccountClientMock struct {
	GetAccountByIdFunc func(ctx context.Context, accountId int64) (*v1.GetAccountByIdResponse, error)
	CreateAccountFunc  func(ctx context.Context, name string, email string, password string) error
	LoginFunc          func(ctx context.Context, email string, password string) (*v1.LoginResponse, error)

	GetAccountByIdCalls int
	CreateAccountCalls  int
	LoginCalls          int
}

func (m *AccountClientMock) GetAccountById(ctx context.Context, accountId int64) (*v1.GetAccountByIdResponse, error) {
	m.GetAccountByIdCalls++

	if m.GetAccountByIdFunc == nil {
		return nil, nil
	}

	return m.GetAccountByIdFunc(ctx, accountId)
}

func (m *AccountClientMock) CreateAccount(ctx context.Context, name string, email string, password string) error {
	m.CreateAccountCalls++

	if m.CreateAccountFunc == nil {
		return nil
	}

	return m.CreateAccountFunc(ctx, name, email, password)
}

func (m *AccountClientMock) Login(ctx context.Context, email string, password string) (*v1.LoginResponse, error) {
	m.LoginCalls++

	if m.LoginFunc == nil {
		return nil, nil
	}

	return m.LoginFunc(ctx, email, password)
}

// ProfileClientMock is a hand-written mock of the ProfileClient gRPC client.
type ProfileClientMock struct {
	GetProfileByIdFunc func(ctx context.Context, id int64) (*v1.GetProfileByIdResponse, error)

	GetProfileByIdCalls int
	GetProfileByIdId    int64
}

func (m *ProfileClientMock) GetProfileById(ctx context.Context, id int64) (*v1.GetProfileByIdResponse, error) {
	m.GetProfileByIdCalls++
	m.GetProfileByIdId = id

	if m.GetProfileByIdFunc == nil {
		return nil, nil
	}

	return m.GetProfileByIdFunc(ctx, id)
}
