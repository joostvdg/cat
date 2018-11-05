package application

type Application struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	UUID        string        `json:"uuid"`
	Namespace   string        `json:"namespace"`
	ArtifactIDs []string      `json:"artifactIDs"`
	Sources     []string      `json:"sources"`
	Components  []Application `json:"components"`
	Labels      []Label       `json:"labels"`
	Annotations []Annotation  `json:"annotations"`
}
