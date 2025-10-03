func simplifyPath(path string) string {
	stack := make([]string, 0, len(path))
	fields := strings.Split(path, "/")

	for _, v := range fields {
		if v == "" {
			continue
		}

		switch v {
		case ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
        default:
		    stack = append(stack,v)
		}
	}

	return "/" + strings.Join(stack, "/")
}
