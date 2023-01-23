package api

import (
	"os"
	"testing"
	"time"

	db "github.com/cindy-cyber/simpleBank/db/sqlc"
	"github.com/cindy-cyber/simpleBank/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}


func TestMain(m *testing.M) { // entry point of all unit tests inside a specific golang package (in this case, package db)
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}