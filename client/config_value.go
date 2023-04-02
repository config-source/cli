package client

type ConfigValue struct {
	Model

	EnvironmentId      int64 `json:"environment_id"`
	ConfigurationKeyId int64 `json:"configuration_key_id"`

	Key        string  `json:"key"`
	ValueType  string  `json:"value_type"`
	StrValue   string  `json:"str_value"`
	IntValue   int64   `json:"int_value"`
	FloatValue float64 `json:"float_value"`
	BoolValue  bool    `json:"bool_value"`
}

func (cv ConfigValue) IsBool() bool {
	return cv.ValueType == "boolean"
}

func (cv ConfigValue) IsInt() bool {
	return cv.ValueType == "integer"
}

func (cv ConfigValue) IsFloat() bool {
	return cv.ValueType == "float"
}

func (cv ConfigValue) IsStr() bool {
	return cv.ValueType == "string"
}

func (cv ConfigValue) Value() interface{} {
	switch {
	case cv.IsBool():
		return cv.BoolValue
	case cv.IsStr():
		return cv.StrValue
	case cv.IsFloat():
		return cv.FloatValue
	case cv.IsInt():
		return cv.IntValue
	default:
		return nil
	}
}
