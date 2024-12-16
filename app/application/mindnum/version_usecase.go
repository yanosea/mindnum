package mindnum

import ()

// VersionUseCase is a struct that contains the use case of the version.
type VersionUseCase struct{}

// NewVersionUseCase returns a new instance of the VersionUseCase struct.
func NewVersionUseCase() *VersionUseCase {
	return &VersionUseCase{}
}

// VersionUsecaseOutputDto is a DTO struct that contains the output data of the VersionUseCase.
type VersionUsecaseOutputDto struct {
	Version string
}

// Run returns the output of the VersionUseCase.
func (uc *VersionUseCase) Run(version string) *VersionUsecaseOutputDto {
	return &VersionUsecaseOutputDto{
		Version: version,
	}
}
