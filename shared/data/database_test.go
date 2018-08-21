package data

import (
	"testing"

	"github.com/currency-converter/shared/config"
	"github.com/stretchr/testify/assert"
)

//TestDabaseInstance : Test case for database instance creation
func TestDatabaseInstance(t *testing.T) {
	cfg, _ := config.New("../config/")

	dbInstance, err := NewDbFactory(cfg)

	//assert database factory
	assert.Nil(t, err)
	assert.NotNil(t, dbInstance)
}

func TestDatabaseConnection(t *testing.T) {
	cfg, _ := config.New("../config/")

	dbInstance, err := NewDbFactory(cfg)

	//assert database factory
	assert.Nil(t, err)
	assert.NotNil(t, dbInstance)

	//test connection
	conn, err := dbInstance.DBConnection()
	assert.Nil(t, err)
	assert.NotNil(t, conn)
}
