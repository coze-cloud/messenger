package messenger

type MessageSerializer interface {
	Serialize() (string, error)
}

type MessageDeserializer interface {
	Deserialize() (Message, error)
}
