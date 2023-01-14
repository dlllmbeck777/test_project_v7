package errorsApp

const defaultErr = 99

var listErrors = map[int]string{
	1: "unknow route 123",
	2: "error with unmarshal request body",
	3: "error with create user",
	4: "id not found",
	5: "error with get user data",
}

func GetErr(num int, systemError string) (int, string, string) {
	err, ok := listErrors[num]
	if !ok {
		return defaultErr, listErrors[defaultErr], ""
	}
	//if len(str) > 0 {
	//	err = err + " : " + str[0]
	//}
	return num, err, systemError
}
