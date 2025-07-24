package sample

import (
	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//var activityMd = activity.ToMetadata(&Settings{}, &Input{})

// New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {
	// No settings needed for simple logging
	act := &Activity{}
	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	logger := ctx.Logger()

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}
	logger.Debugf("Input: %s", input.LogText)

	// Log at different levels based on message content or always log all levels
	logger.Debug("DEBUG: Processing log message: ", input.LogText)
	logger.Info("INFO: ", input.LogText)
	logger.Warn("WARN: This is a warning with message: ", input.LogText)

	// You can also conditionally log at error level
	if input.LogText == "error" {
		logger.Error("ERROR: Error message detected: ", input.LogText)
	}

	// Set the output parameter with the result
	output := &Output{
		Result: "Logging completed successfully for: " + input.LogText,
	}
	err = ctx.SetOutputObject(output)
	if err != nil {
		logger.Errorf("Error setting output object: %v", err)
		return false, err
	}

	return true, nil
}
