package impl

import (
	"context"
	"encoding/json"
	"log"
	"vehicle-log-graphql-api/internal/common/publisher"
)

func (p publishers) LogIn(ctx context.Context, req publisher.MessageLogIn) error {
	b, err := json.Marshal(&req)
	err = p.producer.Publish(publisher.TopicLogIn, b)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
