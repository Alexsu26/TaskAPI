package handler

func SuccessResp(data map[string]any) map[string]any {
	resp := make(map[string]any)
	resp["status"] = "ok"
	if data != nil {
		resp["data"] = data
	}
	return resp
}

func FailResp(data map[string]any) map[string]any {
	resp := make(map[string]any)
	resp["status"] = "error"
	resp["error"] = data
	return resp
}
