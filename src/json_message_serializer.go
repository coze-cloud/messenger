package messenger

import "encoding/json"

type jsonMessageSerializer struct {
	message Message
}

func newJsonMessageSerializer(message Message) *jsonMessageSerializer {
	return &jsonMessageSerializer{message: message}
}

func (j jsonMessageSerializer) Serialize() (string, error) {
	serializedMessage, err := json.Marshal(j.message)
	if err != nil {
		return "", err
	}
	return string(serializedMessage), nil
}

// ---

type jsonMessageDeserializer struct {
	message string
}

func newJsonMessageDeserializer(message string) *jsonMessageDeserializer {
	return &jsonMessageDeserializer{message: message}
}

func (j jsonMessageDeserializer) Deserialize() (Message, error) {
	message := Message{}
	err := json.Unmarshal([]byte(j.message), &message)
	if err != nil {
		return Message{}, err
	}
	return message, nil
}

