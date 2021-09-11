package command

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/ssym0614/go-bto-bot/internal/app/domain"
	"github.com/ssym0614/go-bto-bot/internal/app/repository"
	"log"
	"strings"
)

type Worker struct {
	r repository.Receiver
}

func (w *Worker) Start(ctx context.Context) error {
	go func() {
		if err := w.r.Start(ctx); err != nil && !errors.Is(err, context.Canceled) {
			log.Println(err)
		}
	}()
	commandChannel := w.r.MessageReceptionChannel()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case commandInteractor, ok := <-commandChannel:
			if !ok {
				return errors.New("command channel has been closed")
			}
			if err := w.execute(commandInteractor); err != nil {
				log.Println(err)
			}
		}
	}
}

func (w *Worker) execute(i domain.MessageInteractor) error {
	splitMessage := strings.Split(i.Message(), " ")
	commandName := splitMessage[0]
	args := splitMessage[1:]
	fs := flag.NewFlagSet(commandName, flag.ContinueOnError)
	wordPtr := fs.String("word", "foo", "a string")

	buf := bytes.NewBufferString("")

	fs.SetOutput(buf)
	_ = fs.Parse(args)
	if buf.Len() == 0 {
		response := fmt.Sprintf("value of word option: %s, remaining args: %+v", *wordPtr, fs.Args())
		if err := i.Send(response); err != nil {
			return err
		}

	} else {
		if err := i.Send(buf.String()); err != nil {
			return err
		}
	}
	return nil
}

func NewWorker(r repository.Receiver) *Worker {
	return &Worker{r: r}
}
