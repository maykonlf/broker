package subscriber

type PrefetchQos struct {
	Count    int
	Size     int
	IsGlobal bool
}
