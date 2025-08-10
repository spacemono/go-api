package config

//
//type Bytes []byte
//
//func (e *Bytes) UnmarshalText(text []byte) error {
//	value := Bytes(strings.ToLower(string(text)))
//
//	switch value {
//	case development, production:
//		*e = value
//		return nil
//	default:
//		return fmt.Errorf("unable to parse env: %s", text)
//	}
//}
