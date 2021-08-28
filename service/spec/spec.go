package spec

type Spec struct {
	Input  Input  `json:"input" validate:"required"`
	Resize Resize `json:"resize" validate:"required"`
	Crop   Crop   `json:"crop" validate:"required"`
	Effect Effect `json:"effect" validate:"required"`
}

func NewSpec() *Spec {
	return &Spec{}
}

type Input struct {
	Url string `json:"url" validate:"required,url"`
}

type Resize struct {
	Width           int  `json:"width" validate:"required_with=Height,numeric,min=0"`
	Height          int  `json:"height" validate:"required_with=Width,numeric,min=0"`
	KeepAspectRatio bool `json:"keepAspectRatio" validate:"omitempty"`
}

type Crop struct {
	Anchor string `json:"anchor" validate:"required,oneof=none top center bottom"`
}

type Effect struct {
	Type string `json:"type" validate:"required,oneof=none blur sharpening brightness"`
}
