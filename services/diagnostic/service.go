package diagnostic

import (
	"github.com/influxdata/kapacitor"
	alertservice "github.com/influxdata/kapacitor/services/alert"
	"github.com/influxdata/kapacitor/services/alerta"
	"github.com/influxdata/kapacitor/services/hipchat"
	"github.com/influxdata/kapacitor/services/httpd"
	"github.com/influxdata/kapacitor/services/pagerduty"
	"github.com/influxdata/kapacitor/services/reporting"
	"github.com/influxdata/kapacitor/services/slack"
	"github.com/influxdata/kapacitor/services/smtp"
	"github.com/influxdata/kapacitor/services/storage"
	"github.com/influxdata/kapacitor/services/task_store"
	udfservice "github.com/influxdata/kapacitor/services/udf"
	"github.com/influxdata/kapacitor/services/victorops"
	"go.uber.org/zap"
)

type Service interface {
	NewVictorOpsHandler() victorops.Diagnostic
	NewSlackHandler() slack.Diagnostic
	NewAlertaHandler() alerta.Diagnostic
	NewHipChatHandler() hipchat.Diagnostic
	NewPagerDutyHandler() pagerduty.Diagnostic
	NewSMTPHandler() smtp.Diagnostic

	NewStorageHandler() storage.Diagnostic
	NewTaskStoreHandler() task_store.Diagnostic
	NewReportingHandler() reporting.Diagnostic
	NewHTTPDHandler() httpd.Diagnostic
	NewKapacitorHandler() kapacitor.Diagnostic
	NewAlertServiceHandler() alertservice.Diagnostic
	NewUDFServiceHandler() udfservice.Diagnostic
}

type service struct {
	logger *zap.Logger
}

func NewService() Service {
	// TODO: change
	l := zap.NewExample()
	return &service{
		logger: l,
	}
}

func (s *service) NewVictorOpsHandler() victorops.Diagnostic {
	return &VictorOpsHandler{
		l: s.logger.With(zap.String("service", "victorops")),
	}
}

func (s *service) NewSlackHandler() slack.Diagnostic {
	return &SlackHandler{
		l: s.logger.With(zap.String("service", "slack")),
	}
}

func (s *service) NewTaskStoreHandler() task_store.Diagnostic {
	return &TaskStoreHandler{
		l: s.logger.With(zap.String("service", "task_store")),
	}
}

func (s *service) NewReportingHandler() reporting.Diagnostic {
	return &ReportingHandler{
		l: s.logger.With(zap.String("service", "reporting")),
	}
}

func (s *service) NewStorageHandler() storage.Diagnostic {
	return &StorageHandler{
		l: s.logger.With(zap.String("service", "storage")),
	}
}

func (s *service) NewHTTPDHandler() httpd.Diagnostic {
	return &HTTPDHandler{
		l: s.logger.With(zap.String("service", "http")),
	}
}

func (s *service) NewAlertaHandler() alerta.Diagnostic {
	return &AlertaHandler{
		l: s.logger.With(zap.String("service", "alerta")),
	}
}

func (s *service) NewKapacitorHandler() kapacitor.Diagnostic {
	return &KapacitorHandler{
		l: s.logger.With(zap.String("service", "kapacitor")), // TODO: what here
	}
}

func (s *service) NewAlertServiceHandler() alertservice.Diagnostic {
	return &AlertServiceHandler{
		l: s.logger.With(zap.String("service", "alert")),
	}
}

func (s *service) NewHipChatHandler() hipchat.Diagnostic {
	return &HipChatHandler{
		l: s.logger.With(zap.String("service", "hipchat")),
	}
}

func (s *service) NewPagerDutyHandler() pagerduty.Diagnostic {
	return &PagerDutyHandler{
		l: s.logger.With(zap.String("service", "pagerduty")),
	}
}

func (s *service) NewSMTPHandler() smtp.Diagnostic {
	return &SMTPHandler{
		l: s.logger.With(zap.String("service", "smtp")),
	}
}

func (s *service) NewUDFServiceHandler() udfservice.Diagnostic {
	return &UDFServiceHandler{
		l: s.logger.With(zap.String("service", "udf")),
	}
}
