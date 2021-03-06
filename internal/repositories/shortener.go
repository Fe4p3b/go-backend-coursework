package repositories

type ShortenerRepository interface {
	Find(string) (string, error)
	Save(string, string) error
	AddCount(string) error
	GetVisitorCounter(string) (int, error)
}
