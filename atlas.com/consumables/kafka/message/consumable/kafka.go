package consumable

const (
	EnvCommandTopic = "COMMAND_TOPIC_CONSUMABLE"

	CommandRequestItemConsume = "REQUEST_ITEM_CONSUME"
	CommandRequestScroll      = "REQUEST_SCROLL"
)

type Command[E any] struct {
	CharacterId uint32 `json:"characterId"`
	Type        string `json:"type"`
	Body        E      `json:"body"`
}

type RequestItemConsumeBody struct {
	Source   int16  `json:"source"`
	ItemId   uint32 `json:"itemId"`
	Quantity int16  `json:"quantity"`
}

type RequestScrollBody struct {
	ScrollSlot      int16 `json:"scrollSlot"`
	EquipSlot       int16 `json:"equipSlot"`
	WhiteScroll     bool  `json:"whiteScroll"`
	LegendarySpirit bool  `json:"legendarySpirit"`
}
