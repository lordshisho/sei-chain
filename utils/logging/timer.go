package logging

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	"time"
)

type Timer struct {
	name   string
	start  time.Time
	logger log.Logger
}

func NewTimer(name string, ctx sdk.Context) *Timer {
	return &Timer{
		name:   name,
		start:  time.Now(),
		logger: ctx.Logger(),
	}
}

func (t *Timer) Stop() {
	t.logger.Info(fmt.Sprintf("[Timer] %s took %dms", t.name, time.Since(t.start).Microseconds()))
}
