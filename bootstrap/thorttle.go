package bootstrap

import (
	"github.com/juju/ratelimit"
	"time"
)

var Bucket *ratelimit.Bucket

func init() {
	Bucket = ratelimit.NewBucket(1*time.Second, 100)
}
