package factory

// ApplicationServiceFactory creates application layer services
type ApplicationServiceFactory struct{}

// NewApplicationServiceFactory creates a new application service factory
func NewApplicationServiceFactory() *ApplicationServiceFactory {
	return &ApplicationServiceFactory{}
}
