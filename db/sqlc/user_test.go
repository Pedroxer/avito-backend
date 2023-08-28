package sqlc

import (
	"context"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

var key = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func TestCreateUser(t *testing.T) {
	res := ""
	for i := 0; i < 6; i++ {
		res += string(key[rand.Intn(len(key))])
	}

	arg := res
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, user)
	require.Equal(t, arg, user.Name)
}
