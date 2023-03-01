package Backend

import (
	"fmt"
	"net/http"
	"time"
)

// checkIP executes check IP request
func (ca *Backend) clock(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, e := w.Write([]byte(fmt.Sprintf("%d", time.Now().Unix())))
	if e != nil {
		panic("error writing response")
	}
}
