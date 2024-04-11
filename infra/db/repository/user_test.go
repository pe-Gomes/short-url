package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/pe-Gomes/short-url/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	pass := util.RandomString(6)
	hashedPassword, err := util.HashPassword(pass)

	require.NoError(t, err)

	arg := CreateUserParams{
		Name:     util.RandomString(10),
		Email:    fmt.Sprintf("%s@email.com", util.RandomString(6)),
		Password: hashedPassword,
	}

	user, err := testStore.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)
	require.NotEqual(t, pass, user.Password)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testStore.GetUser(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testStore.GetUserByEmail(context.Background(), user1.Email)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestListUsers(t *testing.T) {
	_ = createRandomUser(t)
	_ = createRandomUser(t)

	arg := ListUsersParams{Limit: 2, Offset: 0}

	users, err := testStore.ListUsers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, users, 2)
	require.NotEmpty(t, users)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := testStore.DeleteUser(context.Background(), user1.ID)

	require.NoError(t, err)

	user2, err := testStore.GetUser(context.Background(), user1.ID)

	require.Error(t, err)
	require.Empty(t, user2)
}
