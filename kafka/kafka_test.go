package kafka_test

import (
	"testing"
	"time"

	"github.com/Shopify/sarama"
)

func TestKafka(t *testing.T) {
	cf := sarama.NewConfig()
	cf.Producer.Return.Successes = true
	cf.Consumer.Return.Errors = true

	p, err := sarama.NewSyncProducer([]string{"localhost:9092"}, cf)
	if err != nil {
		t.Error(err)
	}
	defer p.Close()
	tp := "test"
	for i := 0; i < 10; i++ {
		go func() {
			p.SendMessage(&sarama.ProducerMessage{
				Topic: tp,
				Key:   sarama.StringEncoder("foo"),
				Value: sarama.StringEncoder("bar"),
			})
		}()
	}

	c, err := sarama.NewConsumer([]string{"localhost:9092"}, cf)
	if err != nil {
	}
	defer c.Close()
	pc, err := c.ConsumePartition(tp, 0, sarama.OffsetNewest)
	if err != nil {
		t.Error(err)
	}
	timeout := time.After(1 * time.Second)
	for {
		select {
		case err := <-pc.Errors():
			t.Log(err)
		case m := <-pc.Messages():
			t.Log("Received messages ", string(m.Key), string(m.Value), m.Offset)
		case <-timeout:
			return
		}
	}
}
