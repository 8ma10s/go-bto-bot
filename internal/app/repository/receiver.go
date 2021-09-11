package repository

import (
	"context"
	"github.com/ssym0614/go-bto-bot/internal/app/domain"
)

type Receiver interface {
	Start(ctx context.Context) error
	MessageReceptionChannel() chan domain.MessageInteractor
}
