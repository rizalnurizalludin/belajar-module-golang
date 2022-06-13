package belajar_module

/**
a is number1 as integer
b is number2 as integer
opr is operator as string
return is integer and error
*/
func Calculate(a, b int, opr string)(int, string) {
	var result int
	var err string
	switch opr{
	case "+" ,"tambah":
		result=a+b
	case "-", "kurang":
		result=a-b
	case "*", "kali":
		result=a*b
	case "/" , "bagi":
		result=a/b
	default:
		err="error operator name"
	}

	return result, err
}
