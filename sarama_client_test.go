package kafka

import (
	"reflect"
	"testing"

	"github.com/Shopify/sarama"
)

func TestNewConsumerClient(t *testing.T) {
	type args struct {
		brokers string
		group   string
		topics  string
		oldest  bool
		verbose bool
	}
	tests := []struct {
		name string
		args args
		want *consumerClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConsumerClient(tt.args.brokers, tt.args.group, tt.args.topics, tt.args.oldest, tt.args.verbose); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConsumerClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClientWithApollo(t *testing.T) {
	tests := []struct {
		name string
		want *consumerClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClientWithApollo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClientWithApollo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_consumerClient_Consume(t *testing.T) {
	type fields struct {
		brokers  string
		group    string
		topics   string
		oldest   bool
		verbose  bool
		version  string
		consumer sarama.ConsumerGroupHandler
	}
	type args struct {
		consumer sarama.ConsumerGroupHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *consumerClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &consumerClient{
				brokers:  tt.fields.brokers,
				group:    tt.fields.group,
				topics:   tt.fields.topics,
				oldest:   tt.fields.oldest,
				verbose:  tt.fields.verbose,
				version:  tt.fields.version,
				consumer: tt.fields.consumer,
			}
			if got := c.Consume(tt.args.consumer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("consumerClient.Consume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_consumerClient_ConsumeWithCommit(t *testing.T) {
	type fields struct {
		brokers  string
		group    string
		topics   string
		oldest   bool
		verbose  bool
		version  string
		consumer sarama.ConsumerGroupHandler
	}
	type args struct {
		handler func(message *sarama.ConsumerMessage) error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *consumerClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &consumerClient{
				brokers:  tt.fields.brokers,
				group:    tt.fields.group,
				topics:   tt.fields.topics,
				oldest:   tt.fields.oldest,
				verbose:  tt.fields.verbose,
				version:  tt.fields.version,
				consumer: tt.fields.consumer,
			}
			if got := c.ConsumeWithCommit(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("consumerClient.ConsumeWithCommit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_consumerClient_ClusterVersion(t *testing.T) {
	type fields struct {
		brokers  string
		group    string
		topics   string
		oldest   bool
		verbose  bool
		version  string
		consumer sarama.ConsumerGroupHandler
	}
	type args struct {
		version string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *consumerClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &consumerClient{
				brokers:  tt.fields.brokers,
				group:    tt.fields.group,
				topics:   tt.fields.topics,
				oldest:   tt.fields.oldest,
				verbose:  tt.fields.verbose,
				version:  tt.fields.version,
				consumer: tt.fields.consumer,
			}
			if got := c.ClusterVersion(tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("consumerClient.ClusterVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_consumerClient_Start(t *testing.T) {
	type fields struct {
		brokers  string
		group    string
		topics   string
		oldest   bool
		verbose  bool
		version  string
		consumer sarama.ConsumerGroupHandler
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &consumerClient{
				brokers:  tt.fields.brokers,
				group:    tt.fields.group,
				topics:   tt.fields.topics,
				oldest:   tt.fields.oldest,
				verbose:  tt.fields.verbose,
				version:  tt.fields.version,
				consumer: tt.fields.consumer,
			}
			c.Start()
		})
	}
}

func Test_dc_Setup(t *testing.T) {
	type fields struct {
		handler func(message *sarama.ConsumerMessage) error
	}
	type args struct {
		session sarama.ConsumerGroupSession
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := &dc{
				handler: tt.fields.handler,
			}
			if err := dc.Setup(tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("dc.Setup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dc_Cleanup(t *testing.T) {
	type fields struct {
		handler func(message *sarama.ConsumerMessage) error
	}
	type args struct {
		session sarama.ConsumerGroupSession
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := &dc{
				handler: tt.fields.handler,
			}
			if err := dc.Cleanup(tt.args.session); (err != nil) != tt.wantErr {
				t.Errorf("dc.Cleanup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_dc_ConsumeClaim(t *testing.T) {
	type fields struct {
		handler func(message *sarama.ConsumerMessage) error
	}
	type args struct {
		session sarama.ConsumerGroupSession
		claim   sarama.ConsumerGroupClaim
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dc := &dc{
				handler: tt.fields.handler,
			}
			if err := dc.ConsumeClaim(tt.args.session, tt.args.claim); (err != nil) != tt.wantErr {
				t.Errorf("dc.ConsumeClaim() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
