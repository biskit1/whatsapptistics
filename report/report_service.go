package report

import (
	"context"
	"log"
)

// ReportService represents a service for interacting with the Report Repository
type ReportService interface {
	Get(ctx context.Context, id string) (*Report, error)
	New(*Report) error
}

type reportService struct {
	rr     ReportRepository
	logger *log.Logger
}

// Get retrieves a report from an injected db
func (rs *reportService) Get(ctx context.Context, id string) (*Report, error) {
	r, err := rs.rr.Get(ctx, id)
	if err != nil {
		rs.logger.Print(err)
		return nil, err
	}

	return r, nil
}

// New creats a new report and stores it in an injected repository. // TODO: Add some form of notification in here
func (rs *reportService) New(r *Report) error {
	if err := rs.rr.Store(r); err != nil {
		rs.logger.Print(err)
		return err
	}
	return nil
}

// NewReportService returns a ReportService instance with the dependencies injected
func NewReportService(rr ReportRepository, logger *log.Logger) *reportService {
	return &reportService{
		rr:     rr,
		logger: logger,
	}
}
