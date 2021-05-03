package testrepo

import (
	"context"

	"github.com/decred/dcrdata/pubsub/v4/psclient"
)

func TestFunc() {
	_, _ = psclient.New("", context.Background(), nil)
}
