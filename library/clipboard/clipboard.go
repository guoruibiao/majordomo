package clipboard

type clipboard interface {
	Get() (string, error)
	Set(content string) error
}