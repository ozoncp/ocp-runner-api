package api_test

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"

	"github.com/ozoncp/ocp-runner-api/internal/api"
	"github.com/ozoncp/ocp-runner-api/internal/mocks"
	"github.com/ozoncp/ocp-runner-api/internal/repo"
	server "github.com/ozoncp/ocp-runner-api/pkg/ocp-runner-api"
)

var _ = Describe("Api", func() {
	var (
		err error

		ctrl *gomock.Controller
		ctx  context.Context

		db       *sqlx.DB
		dbMock   sqlxmock.Sqlmock
		prodMock *mocks.MockProducer

		r       repo.Repo
		service server.OcpRunnerServiceServer
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		db, dbMock, err = sqlxmock.Newx()
		r = repo.New(db)

		prodMock = mocks.NewMockProducer(ctrl)
		prodMock.EXPECT().SendMessage(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

		service = api.NewRunnerApi(r, prodMock)
	})

	AfterEach(func() {
		ctrl.Finish()
		db.Close()
	})

	Context("if create method called", func() {
		var (
			request  *server.CreateRunnerRequest
			response *server.CreateRunnerResponse
		)

		BeforeEach(func() {
			request = &server.CreateRunnerRequest{
				Os:   "some_os",
				Arch: "some_arch",
			}

			dbMock.
				ExpectExec("INSERT INTO runners").
				WithArgs(sqlxmock.AnyArg(), request.Os, request.Arch).
				WillReturnResult(sqlxmock.NewResult(0, 1))

			response, err = service.CreateRunner(ctx, request)
		})

		It("should be added a new runner", func() {
			Expect(response).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})
	})

	Context("if update method called", func() {
		var (
			request  *server.UpdateRunnerRequest
			response *server.UpdateRunnerResponse
		)

		BeforeEach(func() {
			request = &server.UpdateRunnerRequest{
				Guid: "test-11eb-be48-0242ac160006",
				Os:   "new_val",
				Arch: "new_val",
			}

			dbMock.
				ExpectExec("UPDATE runners").
				WithArgs(request.Os, request.Arch, request.Guid).
				WillReturnResult(sqlxmock.NewResult(0, 1))

			response, err = service.UpdateRunner(ctx, request)
		})

		It("should be updated existing runner", func() {
			Expect(response).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})
	})

	Context("if update method called with one field", func() {
		var (
			request  *server.UpdateRunnerRequest
			response *server.UpdateRunnerResponse
		)

		BeforeEach(func() {
			request = &server.UpdateRunnerRequest{
				Guid: "test-11eb-be48-0242ac160006",
				Os:   "new_val",
			}

			dbMock.
				ExpectExec("UPDATE runners").
				WithArgs(request.Os, request.Guid).
				WillReturnResult(sqlxmock.NewResult(0, 1))

			response, err = service.UpdateRunner(ctx, request)
		})

		It("should be updated existing runner", func() {
			Expect(response).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})
	})

	Context("if update method called without guid", func() {
		var (
			request  *server.UpdateRunnerRequest
			response *server.UpdateRunnerResponse
		)

		BeforeEach(func() {
			request = &server.UpdateRunnerRequest{}
			response, err = service.UpdateRunner(ctx, request)
		})

		It("should be error", func() {
			Expect(response).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("if update method called without any new fields", func() {
		var (
			request  *server.UpdateRunnerRequest
			response *server.UpdateRunnerResponse
		)

		BeforeEach(func() {
			request = &server.UpdateRunnerRequest{Guid: "test-11eb-be48-0242ac160006"}

			dbMock.
				ExpectExec("UPDATE runners").
				WithArgs(request.Guid, request.Os, request.Arch).
				WillReturnResult(sqlxmock.NewResult(0, 1))

			response, err = service.UpdateRunner(ctx, request)
		})

		It("should be error", func() {
			Expect(response).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("if remove method called with guid", func() {
		var (
			request  *server.RemoveRunnerRequest
			response *server.RemoveRunnerResponse
		)

		BeforeEach(func() {
			request = &server.RemoveRunnerRequest{Guid: "test-11eb-be48-0242ac160006"}

			dbMock.
				ExpectExec("DELETE FROM runners").
				WithArgs(request.Guid).
				WillReturnResult(sqlxmock.NewResult(0, 1))

			response, err = service.RemoveRunner(ctx, request)
		})

		It("should be removed runner", func() {
			Expect(response).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})
	})

	Context("if remove method called without guid", func() {
		var (
			request  *server.RemoveRunnerRequest
			response *server.RemoveRunnerResponse
		)

		BeforeEach(func() {
			request = &server.RemoveRunnerRequest{}

			dbMock.
				ExpectExec("DELETE FROM runners").
				WithArgs(request.Guid).
				WillReturnResult(sqlxmock.NewResult(0, 1))

			response, err = service.RemoveRunner(ctx, request)
		})

		It("should be error", func() {
			Expect(response).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("if describe method called", func() {
		var (
			request  *server.DescribeRunnerRequest
			response *server.DescribeRunnerResponse
		)

		BeforeEach(func() {
			const guid string = "test-11eb-be48-0242ac160006"
			request = &server.DescribeRunnerRequest{Guid: guid}

			rows := sqlxmock.
				NewRows([]string{"guid", "os", "arch"}).
				AddRow(guid, "windows", "x64")
			dbMock.
				ExpectQuery("^SELECT (.+) FROM runners").
				WillReturnRows(rows)

			response, err = service.DescribeRunner(ctx, request)
		})

		It("should be no errors", func() {
			Expect(response).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})
	})

	Context("if describe method called without guid", func() {
		var (
			request  *server.DescribeRunnerRequest
			response *server.DescribeRunnerResponse
		)

		BeforeEach(func() {
			request = &server.DescribeRunnerRequest{}
			response, err = service.DescribeRunner(ctx, request)
		})

		It("should be error", func() {
			Expect(response).Should(BeNil())
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("if list method called without filters", func() {
		var (
			request  *server.ListFiltersRequest
			response *server.RunnersListResponse
		)

		BeforeEach(func() {
			request = &server.ListFiltersRequest{}

			rows := sqlxmock.
				NewRows([]string{"guid", "os", "arch"}).
				AddRow("test-11eb-be48-0242ac160006", "windows", "x64").
				AddRow("test-21eb-be48-0242ac160006", "linux", "x64").
				AddRow("test-31eb-be48-0242ac160006", "macOS", "x64")
			dbMock.
				ExpectQuery("^SELECT (.+) FROM runners").
				WillReturnRows(rows)

			response, err = service.ListRunners(ctx, request)
		})

		It("should be listed all runners", func() {
			Expect(response).ShouldNot(BeNil())
			Expect(len(response.Runners)).Should(BeEquivalentTo(3))
			Expect(err).Should(BeNil())
		})
	})

	Context("if list method called with guid filter", func() {
		var (
			request  *server.ListFiltersRequest
			response *server.RunnersListResponse
		)

		BeforeEach(func() {
			const guid string = "test-11eb-be48-0242ac160006"
			request = &server.ListFiltersRequest{Guids: []string{guid}}

			rows := sqlxmock.
				NewRows([]string{"guid", "os", "arch"}).
				AddRow(guid, "windows", "x64")
			dbMock.
				ExpectQuery("^SELECT (.+) FROM runners").
				WithArgs(guid).
				WillReturnRows(rows)

			response, err = service.ListRunners(ctx, request)
		})

		It("should be listed single runner", func() {
			Expect(response).ShouldNot(BeNil())
			Expect(len(response.Runners)).Should(BeEquivalentTo(1))
			Expect(err).Should(BeNil())
		})
	})
})
