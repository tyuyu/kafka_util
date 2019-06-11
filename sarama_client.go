package kafka

import (
	"context"
	"git.code.oa.com/SNG_EDU_COMMON_PKG/bingo/config/apollo"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type consumerClient struct {
	brokers  string
	group    string
	topics   string
	oldest   bool
	verbose  bool
	version  string
	consumer sarama.ConsumerGroupHandler
}

/**
    example code
    kafka.NewConsumerClient("10.101.202.58:9092",
	"dts_client_demo",
	"ketang_user_message_binlog_notify",
	true, true).
	ConsumeWithCommit(handler).
	Start()
*/
//new consumer client
func NewConsumerClient(brokers, group, topics string, oldest, verbose bool) *consumerClient {
	c := &consumerClient{
		brokers: brokers,
		group:   group,
		topics:  topics,
		oldest:  oldest,
		verbose: verbose,
		version: "0.10.2.0", //连云端ckafka版本必须是这个，没事别乱改
	}
	return c
}

func NewClientWithApollo() *consumerClient {

	cc, _ := apollo.GetConfig()

	c := &consumerClient{
		brokers: cc.GetStringValue("kafka.sarama.brokers", ""),
		group:   cc.GetStringValue("kafka.sarama.group", ""),
		topics:  cc.GetStringValue("kafka.sarama.topics", ""),
		oldest:  cc.GetBool("kafka.sarama.oldest", true),
		verbose: cc.GetBool("kafka.sarama.verbose", true),
		version: cc.GetStringValue("kafka.sarama.version", "0.10.2.0"), //连云端kafka版本必须是这个，没事别乱改
	}
	return c
}

func (c *consumerClient) Consume(consumer sarama.ConsumerGroupHandler) *consumerClient {
	c.consumer = consumer
	return c
}

func (c *consumerClient) ConsumeWithCommit(handler func(message *sarama.ConsumerMessage) error) *consumerClient {
	c.consumer = &dc{
		handler: handler,
	}
	return c
}

func (c *consumerClient) ClusterVersion(version string) *consumerClient {
	c.version = version
	return c
}

func (c *consumerClient) Start() {

	log.Printf("Starting a new Sarama consumer %#v \n", c)

	if c.verbose {
		sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	}

	version, err := sarama.ParseKafkaVersion(c.version)
	if err != nil {
		panic(err)
	}

	/**
	 * Construct a new Sarama configuration.
	 * The Kafka cluster version has to be defined before the consumer/producer is initialized.
	 */
	config := sarama.NewConfig()
	config.Version = version

	if c.oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	/**
	 * Setup a new Sarama consumer group
	 */

	ctx := context.Background()
	client, err := sarama.NewConsumerGroup(strings.Split(c.brokers, ","), c.group, config)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			err := client.Consume(ctx, strings.Split(c.topics, ","), c.consumer)
			if err != nil {
				panic(err)
			}
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	<-sigterm // Await a sigterm signal before safely closing the consumer

	err = client.Close()
	if err != nil {
		panic(err)
	}
}

type dc struct {
	handler func(message *sarama.ConsumerMessage) error
}

func (dc *dc) Setup(session sarama.ConsumerGroupSession) error {
	log.Println("init success ")
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (dc *dc) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (dc *dc) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for message := range claim.Messages() {
		if c, err := apollo.GetConfig(); err != nil || c.GetBool("kafka.consume.log.enable", true) {
			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s  \n", string(message.Value), message.Timestamp, message.Topic)
		}
		if err := dc.handler(message); err != nil {
			log.Println("handle message fail", err)
		}
		session.MarkMessage(message, "")
	}
	return nil
}
