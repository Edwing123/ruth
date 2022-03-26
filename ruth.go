package ruth

type Message struct {
	Value string
}

func WhoIsRuth() *Message {
	return &Message{"Ruth is a girl I used to chat with."}
}

func WhoAmI() *Message {
	return &Message{"I just don't know"}
}
