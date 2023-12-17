// pkg/utils/utils.go
package utils

import (
	"encoding/json"
	"github.com/gofr-dev/gofr"
	"io/ioutil"
)

func ParseBody(ctx *gofr.Context) {
	body, err := ioutil.ReadAll(ctx.Request().Body)
	if err == nil {
		err = json.Unmarshal([]byte(body), ctx.RequestBody)
	}
}
