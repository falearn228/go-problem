func simplifyPath(path string) string {
	stack := []string{}
	fields := strings.Split(path, "/")

	for i := range fields {
		if len(stack) > 0 && fields[i] == ".." {
			stack = stack[:len(stack)-1]
		} else if fields[i] == "" || fields[i] == "." || fields[i] == ".."  {
			continue
		} else {
			stack = append(stack, fields[i])
		}
	}

    var preres string
    if  len(stack) > 0 && stack[0] != "/" {
        preres = "/" + strings.Join(stack, "/")
    } else {
	    preres = strings.Join(stack, "/")
    }

    if len(preres) == 0 {
        return "/"
    }

	if len(preres) > 0 && preres[len(preres)-1] == '/' {
		return preres[:len(preres)-1]
	} 

	return preres
}
