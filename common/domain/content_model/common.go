package content_model

//获取map第一层string结构的数据
func getStringVariables(variables map[string]interface{}) map[string]string {
	var newVariables = make(map[string]string, len(variables))
	for key, variable := range variables {
		if v, ok := variable.(string); ok {
			newVariables[key] = v
		}
	}
	return newVariables
}
