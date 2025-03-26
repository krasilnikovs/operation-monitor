package kernel

type JobLoader interface {
	LoadJob()
}

type DefaultJobLoader struct {
	loaders []JobLoader
}

func NewDefaultJobLoader(loaders []JobLoader) DefaultJobLoader {
	return DefaultJobLoader{loaders: loaders}
}

func (l DefaultJobLoader) LoadJob() {
	for _, loader := range l.loaders {
		loader.LoadJob()
	}
}
