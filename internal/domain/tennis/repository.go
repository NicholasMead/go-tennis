package tennis

type Repository interface {
	Get(string) Tennis
	Save(Tennis) string
}
