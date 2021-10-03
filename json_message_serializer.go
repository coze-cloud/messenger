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

type jsonDeserializer struct {
	message string
}

func newJsonDeserializer(message string) *jsonDeserializer {
	return &jsonDeserializer{message: message}
}

func (j jsonDeserializer) Deserialize() (Message, error) {
	message := Message{}
	err := json.Unmarshal([]byte(j.message), &message)
	if err != nil {
		return Message{}, err
	}
	return message, nil
}

