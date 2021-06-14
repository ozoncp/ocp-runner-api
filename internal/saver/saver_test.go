package saver_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-runner-api/internal/mocks"
	"github.com/ozoncp/ocp-runner-api/internal/models"
	"github.com/ozoncp/ocp-runner-api/internal/saver"
)

const (
	capacity = 10
)

var _ = Describe("Saver", func() {
	var (
		err error

		ctrl *gomock.Controller
		ctx  context.Context

		mockFlusher *mocks.MockFlusher
		mockAlarm   *mocks.MockAlarm

		runner *models.Runner
		s      saver.Saver

		alarms chan struct{}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.Background()

		mockFlusher = mocks.NewMockFlusher(ctrl)
		mockAlarm = mocks.NewMockAlarm(ctrl)

		alarms = make(chan struct{})
		mockAlarm.EXPECT().Alarm().Return(alarms).AnyTimes()

		s = saver.NewSaver(capacity, mockAlarm, mockFlusher)
	})

	JustBeforeEach(func() {
		s.Init(ctx)
		err = s.Save(ctx, runner)
	})

	AfterEach(func() {
		s.Close()
		ctrl.Finish()
	})

	Context("if context cancel called", func() {
		var (
			cancelFunc context.CancelFunc
		)

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			mockFlusher.EXPECT().Flush(gomock.Any(), gomock.Any()).Return(nil)
		})

		JustBeforeEach(func() {
			cancelFunc()
		})

		It("flusher returns no errors", func() {
			Expect(err).Should(BeNil())
		})
	})

	Context("if alarm is occurring", func() {
		var (
			cancelFunc context.CancelFunc
		)

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			mockFlusher.EXPECT().Flush(gomock.Any(), gomock.Any()).Return(nil).MinTimes(1).MaxTimes(2)
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
			cancelFunc()
		})

		It("flusher flush data without error", func() {
			Expect(err).Should(BeNil())
		})
	})
})

func TestSaver_Save(t *testing.T) {

	t.Run("save should not return error when flushing gourutine running", func(t *testing.T) {
		var (
			ctrl        = gomock.NewController(GinkgoT())
			mockAlarm   = mocks.NewMockAlarm(ctrl)
			sv          = saver.NewSaver(10, mockAlarm, nil)
			ctx, cancel = context.WithCancel(context.Background())

			alarmCh chan struct{}
		)
		mockAlarm.EXPECT().Alarm().Return(alarmCh)

		defer cancel()
		sv.Init(ctx)

		ctx2 := context.Background()
		err := sv.Save(ctx2, &models.Runner{Arch: "foo", Os: "bar", Guid: "some_id"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

}
