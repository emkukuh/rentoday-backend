package constant

import "time"

var (
	OneDay24Hour int64 = time.Now().Add(time.Hour * 24).Unix()
)