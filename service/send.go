package service

import (
	"encoding/json"
	"gorm.io/gorm"
	"math/rand"

	"go.uber.org/zap"
	"github.com/nats-io/nats.go"

	"github.com/lenvendo/ig-absolut-fake-sms/service/model"
)

type WorkerService struct {
	Gorm   *gorm.DB
	Logger *zap.Logger
}

func NewWorkerService(gorm *gorm.DB, logger *zap.Logger) *WorkerService {
	return &WorkerService{Gorm: gorm, Logger: logger}
}

func (ws *WorkerService) SendHandler(m *nats.Msg) {
	ws.Logger.Info("got task", zap.Any("message", m))
	msg := &model.Message{}
	err := json.Unmarshal(m.Data, msg)
	if err != nil {
		ws.Logger.Error("error json unmarshalling", zap.Error(err))
	}

	ws.SendCode(msg)
	err = m.Ack()
	if err != nil {
		ws.Logger.Error("error while acking", zap.Error(err))
	}
}

func (ws *WorkerService) SendCode(message *model.Message) error {

	code := model.NewCodeModel()
	code.Phone = message.Phone
	code.Code = rand.Intn(9999-1000) + 1000
	err := ws.Gorm.
		Table("code").
		Save(code).Error
	if err != nil {
		ws.Logger.Error("code insert failed", zap.Error(err))
		return err
	}
	return nil
}
