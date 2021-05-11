package testrepo

import (
	"context"

	"github.com/decred/dcrdata/v6/pubsub/psclient"
)

func TestFunc() {
	_, _ = psclient.New("", context.Background(), nil)
}
