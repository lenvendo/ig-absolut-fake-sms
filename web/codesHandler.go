package web

import (
	"net/http"
	"fmt"
	"errors"

	"gorm.io/gorm"
	"go.uber.org/zap"

	"github.com/lenvendo/ig-absolut-fake-sms/service/model"
)

type CodesHandler struct {
	Gorm   *gorm.DB
	Logger *zap.Logger
}

func NewCodesHandler(gorm *gorm.DB, logger *zap.Logger) *CodesHandler {
	return &CodesHandler{Gorm: gorm, Logger: logger}
}

func (h *CodesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	codes := make([]*model.Code, 0)
	err := h.Gorm.
		Table("code").
		Raw(`SELECT * FROM "code" order by id desc;`).
		Scan(&codes).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Fprintf(w, "Not found")
			return
		}
		h.Logger.Error("code select failed")
		return
	}

	for _, v := range codes {
		fmt.Fprintln(w, fmt.Sprintf("%s - %d", v.Phone, v.Code))
	}
}
