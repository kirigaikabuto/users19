package users

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
)

type UsersAmqpEndpoints struct {
	store UsersStore
}

func NewUsersAmqpEndpoints(s UsersStore) UsersAmqpEndpoints {
	return UsersAmqpEndpoints{s}
}

func (u *UsersAmqpEndpoints) CreateUserAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		userInfo := &User{}
		err := json.Unmarshal(message.Body, &userInfo)
		if err != nil {
			panic(err)
			return nil
		}
		result, err := u.store.Create(userInfo)
		if err != nil {
			panic(err)
			return nil
		}
		dataJson, err := json.Marshal(result)
		if err != nil {
			panic(err)
			return nil
		}
		return &amqp.Message{Body: dataJson}
	}
}

func (u *UsersAmqpEndpoints) GetUserAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		userInfo := &User{}
		err := json.Unmarshal(message.Body, &userInfo)
		if err != nil {
			panic(err)
			return nil
		}
		result, err := u.store.Get(userInfo.Id)
		if err != nil {
			panic(err)
			return nil
		}
		dataJson, err := json.Marshal(result)
		if err != nil {
			panic(err)
			return nil
		}
		return &amqp.Message{Body: dataJson}
	}
}
