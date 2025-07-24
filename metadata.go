package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASetting string `md:"aSetting,required"`
}

type Input struct {
	AnInput string `md:"anInput,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["anInput"])
	r.AnInput = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"anInput": r.AnInput,
	}
}

// Output is an empty struct since this activity has no outputs.
type Output struct{}

// ToMap is required by the data.StructValue interface.
func (o *Output) ToMap() map[string]interface{} {
	return nil
}

// FromMap is required by the data.StructValue interface.
func (o *Output) FromMap(values map[string]interface{}) error {
	return nil
}
