package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	// No settings needed for simple logging
}

type Input struct {
	LogText string `md:"logtext,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["logtext"])
	r.LogText = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"logtext": r.LogText,
	}
}

// Output struct with exported fields
type Output struct {
	Result string `md:"result"`
}

// ToMap is required by the data.StructValue interface.
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}

// FromMap is required by the data.StructValue interface.
func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["result"])
	o.Result = strVal
	return nil
}
