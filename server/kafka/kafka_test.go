package kafka_test

import (
	"context"
	"encoding/json"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	kafgo "github.com/segmentio/kafka-go"
	"github.com/walmartdigital/katalog/domain"
	"github.com/walmartdigital/katalog/mocks/mock_kafka"
	"github.com/walmartdigital/katalog/mocks/mock_repositories"
	"github.com/walmartdigital/katalog/mocks/mock_server"
	"github.com/walmartdigital/katalog/server"
	"github.com/walmartdigital/katalog/server/kafka"
)

var ctrl *gomock.Controller

func TestAll(t *testing.T) {
	ctrl = gomock.NewController(t)
	defer ctrl.Finish()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kafka")
}

var _ = Describe("Run Consumer on 'created' topic", func() {
	var (
		fakeReaderFactory  *mock_kafka.MockReaderFactory
		fakeReader         *mock_kafka.MockReader
		fakeRepoFactory    *mock_repositories.MockRepositoryFactory
		fakeRepo           *mock_repositories.MockRepository
		fakeMetricsFactory *mock_server.MockMetricsFactory
		fakeMetrics        *mock_server.MockMetrics
		consumer           *kafka.Consumer
		ctx                context.Context
		cancel             context.CancelFunc
		wg                 *sync.WaitGroup
		service            server.Service
	)

	BeforeEach(func() {
		// Initialize the mocked Kafka related objects
		fakeReaderFactory = mock_kafka.NewMockReaderFactory(ctrl)
		fakeReader = mock_kafka.NewMockReader(ctrl)
		fakeReaderFactory.EXPECT().Create(gomock.Any(), gomock.Any()).Return(
			fakeReader,
		).Times(1)

		// Initialize the mocked Repository related objects
		fakeRepoFactory = mock_repositories.NewMockRepositoryFactory(ctrl)
		fakeRepo = mock_repositories.NewMockRepository(ctrl)
		fakeRepoFactory.EXPECT().Create().Return(
			fakeRepo,
		).Times(1)

		// Initialize the mocked Metrics related objects
		fakeMetricsFactory = mock_server.NewMockMetricsFactory(ctrl)
		fakeMetrics = mock_server.NewMockMetrics(ctrl)
		fakeMetrics.EXPECT().IncrementCounter(gomock.Any(), gomock.Any()).AnyTimes()
		fakeMetricsFactory.EXPECT().Create().Return(
			fakeMetrics,
		).Times(1)

		ctx, cancel = context.WithCancel(context.Background())
		service = server.MakeService(fakeRepoFactory.Create(), fakeMetricsFactory)
		wg = new(sync.WaitGroup)
		consumer = kafka.CreateConsumer(ctx, wg, "", "", "created", fakeReaderFactory, &service)
	})

	It("should create a consumer", func() {
		Expect(consumer).NotTo(BeNil())
	})

	It("should create a Deployment", func() {
		wg.Add(1)
		defer wg.Wait()

		var testwg sync.WaitGroup
		testwg.Add(1)
		defer testwg.Wait()

		ss := domain.Deployment{
			ID:         "276797fa-b207-11e9-8527-000d3af9d6b6",
			Name:       "queue-node",
			Generation: 7,
			Namespace:  "amida",
			Labels: map[string]string{
				"HEAD":                   "569de2ecd9f9357b3380664f43c90d07ec6acaff",
				"app":                    "nats",
				"fluxcd.io/sync-gc-mark": "sha256.0fRlq9kqkh2eSDRqXANMzgN8_8jeguja3eDLoE5E0Xo",
			},
			Containers: map[string]string{
				"nats-exporter":  "synadia/prometheus-nats-exporter:0.4.0",
				"nats-streaming": "nats-streaming:0.15.1",
			},
		}

		ssbytes, _ := json.Marshal(ss)

		message := kafgo.Message{
			Topic:     "_katalog.artifact.created",
			Partition: 1,
			Offset:    5,
			Key:       []byte("/deployments/276797fa-b207-11e9-8527-000d3af9d6b6"),
			Value:     ssbytes,
			Headers:   nil,
			Time:      time.Now(),
		}

		resource := domain.Resource{K8sResource: &ss}

		fakeReader.EXPECT().Close().Times(1)
		fakeRepo.EXPECT().CreateResource(resource).Times(1).Do(
			func(r domain.Resource) {
				testwg.Done()
			},
		)
		fakeReader.EXPECT().ReadMessage(ctx).Return(message, nil).Times(1).Do(
			func(c context.Context) {
				cancel()
			},
		)
		Expect(consumer).NotTo(BeNil())
		go consumer.Run()
	})

	It("should create a StatefulSet", func() {
		wg.Add(1)
		defer wg.Wait()

		var testwg sync.WaitGroup
		testwg.Add(1)
		defer testwg.Wait()

		ss := domain.StatefulSet{
			ID:         "276797fa-b207-11e9-8527-000d3af9d6b6",
			Name:       "queue-node",
			Generation: 7,
			Namespace:  "amida",
			Labels: map[string]string{
				"HEAD":                   "569de2ecd9f9357b3380664f43c90d07ec6acaff",
				"app":                    "nats",
				"fluxcd.io/sync-gc-mark": "sha256.0fRlq9kqkh2eSDRqXANMzgN8_8jeguja3eDLoE5E0Xo",
			},
			Containers: map[string]string{
				"nats-exporter":  "synadia/prometheus-nats-exporter:0.4.0",
				"nats-streaming": "nats-streaming:0.15.1",
			},
		}

		ssbytes, _ := json.Marshal(ss)

		message := kafgo.Message{
			Topic:     "_katalog.artifact.created",
			Partition: 1,
			Offset:    5,
			Key:       []byte("/statefulsets/276797fa-b207-11e9-8527-000d3af9d6b6"),
			Value:     ssbytes,
			Headers:   nil,
			Time:      time.Now(),
		}

		resource := domain.Resource{K8sResource: &ss}

		fakeReader.EXPECT().Close().Times(1)
		fakeRepo.EXPECT().CreateResource(resource).Times(1).Do(
			func(r domain.Resource) {
				testwg.Done()
			},
		)
		fakeReader.EXPECT().ReadMessage(ctx).Return(message, nil).Times(1).Do(
			func(c context.Context) {
				cancel()
			},
		)
		Expect(consumer).NotTo(BeNil())
		go consumer.Run()
	})

	It("should create a Service", func() {
		wg.Add(1)
		defer wg.Wait()

		var testwg sync.WaitGroup
		testwg.Add(1)
		defer testwg.Wait()

		i := domain.Instance{Address: "hello"}
		ss := domain.Service{
			ID:         "276797fa-b207-11e9-8527-000d3af9d6b6",
			Name:       "queue-node",
			Port:       1212,
			Address:    "someservice",
			Generation: 7,
			Namespace:  "amida",
			Instances:  []domain.Instance{i},
		}

		ssbytes, _ := json.Marshal(ss)

		message := kafgo.Message{
			Topic:     "_katalog.artifact.created",
			Partition: 1,
			Offset:    5,
			Key:       []byte("/services/276797fa-b207-11e9-8527-000d3af9d6b6"),
			Value:     ssbytes,
			Headers:   nil,
			Time:      time.Now(),
		}

		resource := domain.Resource{K8sResource: &ss}

		fakeReader.EXPECT().Close().Times(1)
		fakeRepo.EXPECT().CreateResource(resource).Times(1).Do(
			func(r domain.Resource) {
				testwg.Done()
			},
		)
		fakeReader.EXPECT().ReadMessage(ctx).Return(message, nil).Times(1).Do(
			func(c context.Context) {
				cancel()
			},
		)
		Expect(consumer).NotTo(BeNil())
		go consumer.Run()
	})

	AfterEach(func() {
	})
})

var _ = Describe("Run Consumer on 'updated' topic", func() {
	var (
		fakeReaderFactory  *mock_kafka.MockReaderFactory
		fakeReader         *mock_kafka.MockReader
		fakeRepoFactory    *mock_repositories.MockRepositoryFactory
		fakeRepo           *mock_repositories.MockRepository
		fakeMetricsFactory *mock_server.MockMetricsFactory
		fakeMetrics        *mock_server.MockMetrics
		upconsumer         *kafka.Consumer
		ctx                context.Context
		cancel             context.CancelFunc
		wg                 *sync.WaitGroup
		service            server.Service
	)

	BeforeEach(func() {
		// Initialize the mocked Kafka related objects
		fakeReaderFactory = mock_kafka.NewMockReaderFactory(ctrl)
		fakeReader = mock_kafka.NewMockReader(ctrl)
		fakeReaderFactory.EXPECT().Create(gomock.Any(), gomock.Any()).Return(
			fakeReader,
		).Times(1)

		// Initialize the mocked Repository related objects
		fakeRepoFactory = mock_repositories.NewMockRepositoryFactory(ctrl)
		fakeRepo = mock_repositories.NewMockRepository(ctrl)
		fakeRepoFactory.EXPECT().Create().Return(
			fakeRepo,
		).Times(1)

		// Initialize the mocked Metrics related objects
		fakeMetricsFactory = mock_server.NewMockMetricsFactory(ctrl)
		fakeMetrics = mock_server.NewMockMetrics(ctrl)
		fakeMetrics.EXPECT().IncrementCounter(gomock.Any(), gomock.Any()).AnyTimes()
		fakeMetricsFactory.EXPECT().Create().Return(
			fakeMetrics,
		).Times(1)

		ctx, cancel = context.WithCancel(context.Background())
		service = server.MakeService(fakeRepoFactory.Create(), fakeMetricsFactory)
		wg = new(sync.WaitGroup)
		upconsumer = kafka.CreateConsumer(ctx, wg, "", "", "updated", fakeReaderFactory, &service)
	})

	It("should create a consumer", func() {
		Expect(upconsumer).NotTo(BeNil())
	})

	It("should update a Deployment", func() {
		wg.Add(1)
		defer wg.Wait()

		var testwg sync.WaitGroup
		testwg.Add(1)
		defer testwg.Wait()

		ss := domain.Deployment{
			ID:         "276797fa-b207-11e9-8527-000d3af9d6b6",
			Name:       "queue-node",
			Generation: 8,
			Namespace:  "amida",
			Labels: map[string]string{
				"HEAD":                   "569de2ecd9f9357b3380664f43c90d07ec6acaff",
				"app":                    "nats",
				"fluxcd.io/sync-gc-mark": "sha256.0fRlq9kqkh2eSDRqXANMzgN8_8jeguja3eDLoE5E0Xo",
				"some":                   "change",
			},
			Containers: map[string]string{
				"nats-exporter":  "synadia/prometheus-nats-exporter:0.4.0",
				"nats-streaming": "nats-streaming:0.15.1",
			},
		}

		ssbytes, _ := json.Marshal(ss)

		message := kafgo.Message{
			Topic:     "_katalog.artifact.updated",
			Partition: 1,
			Offset:    5,
			Key:       []byte("/deployments/276797fa-b207-11e9-8527-000d3af9d6b6"),
			Value:     ssbytes,
			Headers:   nil,
			Time:      time.Now(),
		}

		resource := domain.Resource{K8sResource: &ss}
		fakeReader.EXPECT().Close().Times(1)
		fakeRepo.EXPECT().UpdateResource(resource).Return(&resource, nil).Times(1).Do(
			func(r domain.Resource) {
				testwg.Done()
			},
		)
		fakeReader.EXPECT().ReadMessage(ctx).Return(message, nil).Times(1).Do(
			func(c context.Context) {
				cancel()
			},
		)
		go upconsumer.Run()
	})

	It("should update a Statefulset", func() {
		wg.Add(1)
		defer wg.Wait()

		var testwg sync.WaitGroup
		testwg.Add(1)
		defer testwg.Wait()

		ss := domain.StatefulSet{
			ID:         "276797fa-b207-11e9-8527-000d3af9d6b6",
			Name:       "queue-node",
			Generation: 8,
			Namespace:  "amida",
			Labels: map[string]string{
				"HEAD":                   "569de2ecd9f9357b3380664f43c90d07ec6acaff",
				"app":                    "nats",
				"fluxcd.io/sync-gc-mark": "sha256.0fRlq9kqkh2eSDRqXANMzgN8_8jeguja3eDLoE5E0Xo",
				"some":                   "change",
			},
			Containers: map[string]string{
				"nats-exporter":  "synadia/prometheus-nats-exporter:0.4.0",
				"nats-streaming": "nats-streaming:0.15.1",
			},
		}

		ssbytes, _ := json.Marshal(ss)

		message := kafgo.Message{
			Topic:     "_katalog.artifact.updated",
			Partition: 1,
			Offset:    5,
			Key:       []byte("/statefulsets/276797fa-b207-11e9-8527-000d3af9d6b6"),
			Value:     ssbytes,
			Headers:   nil,
			Time:      time.Now(),
		}

		resource := domain.Resource{K8sResource: &ss}

		fakeReader.EXPECT().Close().Times(1)
		fakeRepo.EXPECT().UpdateResource(resource).Return(&resource, nil).Times(1).Do(
			func(r domain.Resource) {
				testwg.Done()
			},
		)
		fakeReader.EXPECT().ReadMessage(ctx).Return(message, nil).Times(1).Do(
			func(c context.Context) {
				cancel()
			},
		)
		go upconsumer.Run()
	})

	It("should update a Service", func() {
		wg.Add(1)
		defer wg.Wait()

		var testwg sync.WaitGroup
		testwg.Add(1)
		defer testwg.Wait()

		i := domain.Instance{Address: "hello"}
		ss := domain.Service{
			ID:         "276797fa-b207-11e9-8527-000d3af9d6b6",
			Name:       "queue-node",
			Port:       1212,
			Address:    "someservice",
			Generation: 7,
			Namespace:  "amida",
			Instances:  []domain.Instance{i},
		}

		ssbytes, _ := json.Marshal(ss)

		message := kafgo.Message{
			Topic:     "_katalog.artifact.updated",
			Partition: 1,
			Offset:    5,
			Key:       []byte("/services/276797fa-b207-11e9-8527-000d3af9d6b6"),
			Value:     ssbytes,
			Headers:   nil,
			Time:      time.Now(),
		}

		resource := domain.Resource{K8sResource: &ss}

		fakeReader.EXPECT().Close().Times(1)
		fakeRepo.EXPECT().UpdateResource(resource).Return(&resource, nil).Times(1).Do(
			func(r domain.Resource) {
				testwg.Done()
			},
		)
		fakeReader.EXPECT().ReadMessage(ctx).Return(message, nil).Times(1).Do(
			func(c context.Context) {
				cancel()
			},
		)
		go upconsumer.Run()
	})

	AfterEach(func() {
	})
})

var _ = Describe("Run Consumer on 'deleted' topic", func() {
	var (
		fakeReaderFactory  *mock_kafka.MockReaderFactory
		fakeReader         *mock_kafka.MockReader
		fakeRepoFactory    *mock_repositories.MockRepositoryFactory
		fakeRepo           *mock_repositories.MockRepository
		fakeMetricsFactory *mock_server.MockMetricsFactory
		fakeMetrics        *mock_server.MockMetrics
		consumer           *kafka.Consumer
		ctx                context.Context
		cancel             context.CancelFunc
		wg                 *sync.WaitGroup
		service            server.Service
	)

	BeforeEach(func() {
		// Initialize the mocked Kafka related objects
		fakeReaderFactory = mock_kafka.NewMockReaderFactory(ctrl)
		fakeReader = mock_kafka.NewMockReader(ctrl)
		fakeReaderFactory.EXPECT().Create(gomock.Any(), gomock.Any()).Return(
			fakeReader,
		).Times(1)

		// Initialize the mocked Repository related objects
		fakeRepoFactory = mock_repositories.NewMockRepositoryFactory(ctrl)
		fakeRepo = mock_repositories.NewMockRepository(ctrl)
		fakeRepoFactory.EXPECT().Create().Return(
			fakeRepo,
		).Times(1)

		// Initialize the mocked Metrics related objects
		fakeMetricsFactory = mock_server.NewMockMetricsFactory(ctrl)
		fakeMetrics = mock_server.NewMockMetrics(ctrl)
		fakeMetrics.EXPECT().IncrementCounter(gomock.Any(), gomock.Any()).AnyTimes()
		fakeMetricsFactory.EXPECT().Create().Return(
			fakeMetrics,
		).Times(1)

		ctx, cancel = context.WithCancel(context.Background())
		service = server.MakeService(fakeRepoFactory.Create(), fakeMetricsFactory)
		wg = new(sync.WaitGroup)
		consumer = kafka.CreateConsumer(ctx, wg, "", "", "deleted", fakeReaderFactory, &service)
	})

	It("should create a consumer", func() {
		Expect(consumer).NotTo(BeNil())
	})

	It("should delete a Deployment", func() {
		wg.Add(1)
		defer wg.Wait()

		var testwg sync.WaitGroup
		testwg.Add(1)
		defer testwg.Wait()

		ss := domain.Deployment{
			ID:         "276797fa-b207-11e9-8527-000d3af9d6b6",
			Name:       "queue-node",
			Generation: 7,
			Namespace:  "amida",
			Labels: map[string]string{
				"HEAD":                   "569de2ecd9f9357b3380664f43c90d07ec6acaff",
				"app":                    "nats",
				"fluxcd.io/sync-gc-mark": "sha256.0fRlq9kqkh2eSDRqXANMzgN8_8jeguja3eDLoE5E0Xo",
			},
			Containers: map[string]string{
				"nats-exporter":  "synadia/prometheus-nats-exporter:0.4.0",
				"nats-streaming": "nats-streaming:0.15.1",
			},
		}

		ssbytes, _ := json.Marshal(ss)

		message := kafgo.Message{
			Topic:     "_katalog.artifact.deleted",
			Partition: 1,
			Offset:    5,
			Key:       []byte("/deployments/276797fa-b207-11e9-8527-000d3af9d6b6"),
			Value:     ssbytes,
			Headers:   nil,
			Time:      time.Now(),
		}

		resource := domain.Resource{K8sResource: &ss}
		id := "276797fa-b207-11e9-8527-000d3af9d6b6"

		fakeReader.EXPECT().Close().Times(1)
		fakeRepo.EXPECT().GetResource(id).Return(resource, nil).Times(1)
		fakeRepo.EXPECT().DeleteResource(id).Return(nil).Times(1).Do(
			func(id string) {
				testwg.Done()
			},
		)
		fakeReader.EXPECT().ReadMessage(ctx).Return(message, nil).Times(1).Do(
			func(c context.Context) {
				cancel()
			},
		)
		go consumer.Run()
	})

	It("should delete a Statefulset", func() {
		wg.Add(1)
		defer wg.Wait()

		var testwg sync.WaitGroup
		testwg.Add(1)
		defer testwg.Wait()

		ss := domain.StatefulSet{
			ID:         "276797fa-b207-11e9-8527-000d3af9d6b6",
			Name:       "queue-node",
			Generation: 7,
			Namespace:  "amida",
			Labels: map[string]string{
				"HEAD":                   "569de2ecd9f9357b3380664f43c90d07ec6acaff",
				"app":                    "nats",
				"fluxcd.io/sync-gc-mark": "sha256.0fRlq9kqkh2eSDRqXANMzgN8_8jeguja3eDLoE5E0Xo",
			},
			Containers: map[string]string{
				"nats-exporter":  "synadia/prometheus-nats-exporter:0.4.0",
				"nats-streaming": "nats-streaming:0.15.1",
			},
		}

		ssbytes, _ := json.Marshal(ss)

		message := kafgo.Message{
			Topic:     "_katalog.artifact.deleted",
			Partition: 1,
			Offset:    5,
			Key:       []byte("/statefulsets/276797fa-b207-11e9-8527-000d3af9d6b6"),
			Value:     ssbytes,
			Headers:   nil,
			Time:      time.Now(),
		}

		resource := domain.Resource{K8sResource: &ss}
		id := "276797fa-b207-11e9-8527-000d3af9d6b6"

		fakeReader.EXPECT().Close().Times(1)
		fakeRepo.EXPECT().GetResource(id).Return(resource, nil).Times(1)
		fakeRepo.EXPECT().DeleteResource(id).Return(nil).Times(1).Do(
			func(id string) {
				testwg.Done()
			},
		)

		fakeReader.EXPECT().ReadMessage(ctx).Return(message, nil).Times(1).Do(
			func(c context.Context) {
				cancel()
			},
		)
		go consumer.Run()
	})

	It("should delete a Service", func() {
		wg.Add(1)
		defer wg.Wait()

		var testwg sync.WaitGroup
		testwg.Add(1)
		defer testwg.Wait()

		i := domain.Instance{Address: "hello"}
		ss := domain.Service{
			ID:         "276797fa-b207-11e9-8527-000d3af9d6b6",
			Name:       "queue-node",
			Port:       1212,
			Address:    "someservice",
			Generation: 7,
			Namespace:  "amida",
			Instances:  []domain.Instance{i},
		}

		ssbytes, _ := json.Marshal(ss)

		message := kafgo.Message{
			Topic:     "_katalog.artifact.deleted",
			Partition: 1,
			Offset:    5,
			Key:       []byte("/services/276797fa-b207-11e9-8527-000d3af9d6b6"),
			Value:     ssbytes,
			Headers:   nil,
			Time:      time.Now(),
		}

		resource := domain.Resource{K8sResource: &ss}
		id := "276797fa-b207-11e9-8527-000d3af9d6b6"

		fakeReader.EXPECT().Close().Times(1)
		fakeRepo.EXPECT().GetResource(id).Return(resource, nil).Times(1)
		fakeRepo.EXPECT().DeleteResource(id).Return(nil).Times(1).Do(
			func(id string) {
				testwg.Done()
			},
		)
		fakeReader.EXPECT().ReadMessage(ctx).Return(message, nil).Times(1).Do(
			func(c context.Context) {
				cancel()
			},
		)
		go consumer.Run()
	})

	AfterEach(func() {
	})
})
