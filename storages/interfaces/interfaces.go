package interfaces

type Storage interface {
	Add(data int64) (index int64)
	Get(index int64) (data int64)
	Delete(index int64)
	String() string
}
