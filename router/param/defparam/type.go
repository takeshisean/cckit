package defparam

import (
	// "github.com/takeshisean/cckit/router"
	// "github.com/takeshisean/cckit/router/param"
	
	"github.com/takeshisean/cckit/router"
	"github.com/takeshisean/cckit/router/param"
)

func Proto(target interface{}, argPoss ...int) router.MiddlewareFunc {
	return param.Proto(router.DefaultParam, target, argPoss...)
}
