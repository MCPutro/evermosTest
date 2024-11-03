package response

type Response struct {
	Code         int    `json:"code"`
	Success      bool   `json:"success"`
	Message      string `json:"message,omitempty"`
	TotalRecords int    `json:"total_records,omitempty"`
	//ErrorMessage string      `json:"error_message,omitempty"`
	ErrorDetail interface{} `json:"error_detail,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}
