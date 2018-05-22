package nsq_test

import (
	"log"
	"strconv"
	"sync"
	"testing"
)

func newProducer() *nsq.Producer {
	prod, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())
	if err != nil {
		log.Fatalln(err)
	}
	return prod
}

func newConsumer() *nsq.Consumer {
	cons, err := nsq.NewConsumer("test", "test_channel", nsq.NewConfig())
	if err != nil {
		log.Fatalln(err)
	}
	return cons
}

func send(prod *nsq.Producer, topic string, msg []byte) {
	if err := prod.Publish(topic, msg); err != nil {
		log.Fatalln(err)
	}
}

type conshandle struct{}

func (*conshandle) HandleMessage(msg *nsq.Message) error {
	defer wg.Done()
	log.Println("receive from:", msg.NSQDAddress, "message id :", string(msg.ID[:]), "message body:", string(msg.Body))
	return nil
}

func receive(cons *nsq.Consumer) {
	cons.AddHandler(&conshandle{})
	if err := cons.ConnectToNSQLookupd("127.0.0.1:4161"); err != nil {
		log.Fatalln(err)
	}
}

var wg sync.WaitGroup

func TestNSQ(t *testing.T) {
	prod, cons := newProducer(), newConsumer()
	wg.Add(10)
	receive(cons)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			send(prod, "test", []byte(strconv.FormatInt(int64(i), 10)))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
