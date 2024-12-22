package response

// Структуры для запросов и ответов
type RequestBody struct {
	Expression string `json:"expression"`
}

type ResponseBody struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}
