package api

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

func TestMain(m *testing.M) { // entry point of all unit tests inside a specific golang package (in this case, package db)
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}