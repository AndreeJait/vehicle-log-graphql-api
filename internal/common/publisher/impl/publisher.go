package impl

import (
	"github.com/nsqio/go-nsq"
	"vehicle-log-graphql-api/internal/common/publisher"
)

type publishers struct {
	producer *nsq.Producer
}

func New(producer *nsq.Producer) publisher.Publisher {
	return &publishers{
		producer: producer,
	}
}
