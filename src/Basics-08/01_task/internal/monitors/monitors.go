package monitors

type Monitor interface {
	Type() string
	Run() error
}
