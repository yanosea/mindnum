package mindnum

import ()

type VersionUseCase struct{}

func NewVersionUseCase() *VersionUseCase {
	return &VersionUseCase{}
}

type VersionUsecaseOutputDto struct {
	Version string
}

func (uc *VersionUseCase) Run(version string) *VersionUsecaseOutputDto {
	return &VersionUsecaseOutputDto{
		Version: version,
	}
}
