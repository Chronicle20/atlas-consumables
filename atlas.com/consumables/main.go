package main

import (
	"atlas-consumables/kafka/consumer/character"
	"atlas-consumables/kafka/consumer/consumable"
	"atlas-consumables/kafka/consumer/inventory"
	"atlas-consumables/logger"
	"atlas-consumables/service"
	"atlas-consumables/tracing"
	"github.com/Chronicle20/atlas-kafka/consumer"
)

const serviceName = "atlas-consumables"
const consumerGroupId = "Consumables Service"

func main() {
	l := logger.CreateLogger(serviceName)
	l.Infoln("Starting main service.")

	tdm := service.GetTeardownManager()

	tc, err := tracing.InitTracer(l)(serviceName)
	if err != nil {
		l.WithError(err).Fatal("Unable to initialize tracer.")
	}

	cmf := consumer.GetManager().AddConsumer(l, tdm.Context(), tdm.WaitGroup())
	inventory.InitConsumers(l)(cmf)(consumerGroupId)
	character.InitConsumers(l)(cmf)(consumerGroupId)
	consumable.InitConsumers(l)(cmf)(consumerGroupId)
	character.InitHandlers(l)(consumer.GetManager().RegisterHandler)
	consumable.InitHandlers(l)(consumer.GetManager().RegisterHandler)

	tdm.TeardownFunc(tracing.Teardown(l)(tc))

	tdm.Wait()
	l.Infoln("Service shutdown.")
}
