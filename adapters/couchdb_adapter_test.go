package adapters_test

import (
	"context"
)

func TestCouchDB(t *testing.T) {
	ctx := context.Background()
	mongodbUrl := "http://admin:1234@localhost:5984/"
	adapter, err := adapters.GetCouchdbConnection(ctx, mongodbUrl)
	if err != nil {
		t.Fatalf("Error opening database connection: %s", err)
	}

	testHelper(t, ctx, adapter)
}
