package output

import "github.com/fatih/color"

func PrintError(value any) {
	intval, ok := value.(int)
	if ok {
		color.Red("Pas error: %d", intval)
		return
	}
	strValue, ok := value.(string)
	if ok {
		color.Red(strValue)
	}
	errorValue, ok := value.(error)
	if ok {
		color.Red(errorValue.Error())
	}
	color.Red("Undefind error")
	// switch t := value.(type) {
	// case string:
	// 	color.Red(t)
	// case int:
	// 	color.Red("Password error: %d", t)
	// case error:
	// 	color.Red(t.Error())
	// default:
	// 	color.Red("Undefind error")
	// }
}

