package impl

import (
	"context"
	"encoding/json"
	"log"
	"vehicle-log-graphql-api/internal/common/publisher"
)

func (p publishers) LogOut(ctx context.Context, req publisher.MessageLogOut) error {
	b, err := json.Marshal(&req)
	err = p.producer.Publish(publisher.TopicLogOut, b)
	if err != nil {
		log.Fatal(err)
	}

	return err
}
