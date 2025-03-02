package temper

type LogArg interface {
	[]string | LogTime
}

type LogTime struct {
	Format string
	Color  int
}
