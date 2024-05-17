package tests

import (
	"alexnerotd/blog/utils"
	"testing"
)

func TestConnectDb(t *testing.T) {
	_, err := utils.ConnectDB()
	if err != nil {
		t.Errorf("Error conectiong to the database: %v", err)
	}
}
