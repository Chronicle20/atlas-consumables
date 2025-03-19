package consumable

import (
	"atlas-consumables/consumable"
	consumer2 "atlas-consumables/kafka/consumer"
	consumable2 "atlas-consumables/kafka/message/consumable"
	"context"
	"github.com/Chronicle20/atlas-constants/item"
	"github.com/Chronicle20/atlas-kafka/consumer"
	"github.com/Chronicle20/atlas-kafka/handler"
	"github.com/Chronicle20/atlas-kafka/message"
	"github.com/Chronicle20/atlas-kafka/topic"
	"github.com/Chronicle20/atlas-model/model"
	"github.com/sirupsen/logrus"
)

func InitConsumers(l logrus.FieldLogger) func(func(config consumer.Config, decorators ...model.Decorator[consumer.Config])) func(consumerGroupId string) {
	return func(rf func(config consumer.Config, decorators ...model.Decorator[consumer.Config])) func(consumerGroupId string) {
		return func(consumerGroupId string) {
			rf(consumer2.NewConfig(l)("consumable_command")(consumable2.EnvCommandTopic)(consumerGroupId), consumer.SetHeaderParsers(consumer.SpanHeaderParser, consumer.TenantHeaderParser))
		}
	}
}

func InitHandlers(l logrus.FieldLogger) func(rf func(topic string, handler handler.Handler) (string, error)) {
	return func(rf func(topic string, handler handler.Handler) (string, error)) {
		var t string
		t, _ = topic.EnvProvider(l)(consumable2.EnvCommandTopic)()
		_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleRequestItemConsume)))
		_, _ = rf(t, message.AdaptHandler(message.PersistentConfig(handleRequestScroll)))
	}
}

func handleRequestItemConsume(l logrus.FieldLogger, ctx context.Context, c consumable2.Command[consumable2.RequestItemConsumeBody]) {
	if c.Type != consumable2.CommandRequestItemConsume {
		return
	}
	_ = consumable.RequestItemConsume(l)(ctx)(c.CharacterId, c.Body.Source, item.Id(c.Body.ItemId), c.Body.Quantity)
	// TODO issue error
}

func handleRequestScroll(l logrus.FieldLogger, ctx context.Context, c consumable2.Command[consumable2.RequestScrollBody]) {
	if c.Type != consumable2.CommandRequestScroll {
		return
	}
	_ = consumable.RequestScroll(l)(ctx)(c.CharacterId, c.Body.ScrollSlot, c.Body.EquipSlot, c.Body.WhiteScroll, c.Body.LegendarySpirit)
	// TODO issue error
}
