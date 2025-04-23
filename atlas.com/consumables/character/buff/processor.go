package buff

import (
	"atlas-consumables/character/buff/stat"
	buff2 "atlas-consumables/kafka/message/character/buff"
	"atlas-consumables/kafka/producer"
	"context"
	_map "github.com/Chronicle20/atlas-constants/map"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/sirupsen/logrus"
)

type Processor struct {
	l   logrus.FieldLogger
	ctx context.Context
}

func NewProcessor(l logrus.FieldLogger, ctx context.Context) *Processor {
	p := &Processor{
		l:   l,
		ctx: ctx,
	}
	return p
}

func (p *Processor) Apply(m _map.Model, fromId uint32, sourceId int32, duration int32, statups []stat.Model) model.Operator[uint32] {
	return func(characterId uint32) error {
		return producer.ProviderImpl(p.l)(p.ctx)(buff2.EnvCommandTopic)(applyCommandProvider(m, characterId, fromId, sourceId, duration, statups))
	}
}
