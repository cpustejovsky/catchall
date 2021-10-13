package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/cpustejovsky/catchall/logger"
)

func ServerError(log logger.Logger, w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	log.ErrorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
