package flusher_test

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-runner-api/internal/flusher"
	"github.com/ozoncp/ocp-runner-api/internal/mocks"
	"github.com/ozoncp/ocp-runner-api/internal/models"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	errDeadlineExceeded = errors.New("mock error")
)

var _ = Describe("Flusher", func() {
	var (
		err error

		ctrl *gomock.Controller
		ctx  context.Context

		mockRepo *mocks.MockRepo

		runners []models.Runner
		rest    []models.Runner

		f flusher.Flusher

		chunkSize int
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		mockRepo = mocks.NewMockRepo(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.New(mockRepo, chunkSize)
		rest = f.Flush(ctx, runners)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo saves all runners", func() {
		BeforeEach(func() {
			chunkSize = 2
			runners = []models.Runner{{}}

			mockRepo.EXPECT().AddRunners(gomock.Any(), gomock.Any()).Return(nil).MinTimes(1)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeNil())
		})
	})

	Context("repo don't saves any runner", func() {
		BeforeEach(func() {
			chunkSize = 2
			runners = []models.Runner{{}, {}}

			mockRepo.EXPECT().AddRunners(gomock.Any(), gomock.Len(chunkSize)).Return(errDeadlineExceeded)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeEquivalentTo(runners))
		})
	})

	Context("repo saves half runners", func() {
		BeforeEach(func() {
			runners = []models.Runner{{}, {}}
			chunkSize = len(runners) / 2

			gomock.InOrder(
				mockRepo.EXPECT().AddRunners(gomock.Any(), gomock.Len(chunkSize)).Return(nil),
				mockRepo.EXPECT().AddRunners(gomock.Any(), gomock.Any()).Return(errDeadlineExceeded),
			)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeEquivalentTo(runners[chunkSize:]))
		})
	})
})
