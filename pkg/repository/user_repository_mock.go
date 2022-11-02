package repository

import "github.com/stretchr/testify/mock"

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) GetUserByID(id string) (UserSchema, error) {
	args := m.Called()
	return args.Get(0).(UserSchema), nil
}

func (m *UserRepositoryMock) ResetMock() {
	m.Mock = mock.Mock{}
}
