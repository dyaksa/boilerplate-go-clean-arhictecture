package response

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
type Wrapper struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []Error     `json:"errors,omitempty"`
	Details string      `json:"details,omitempty"`
}

func Send(msg, status string) *Wrapper {
	return &Wrapper{
		Status:  status,
		Message: msg,
	}
}

func (w *Wrapper) WithErrors(errs []Error) *Wrapper {
	w.Errors = errs
	return w
}

func (w *Wrapper) WithDetails(details string) *Wrapper {
	w.Details = details
	return w
}

func (w *Wrapper) WithData(data interface{}) *Wrapper {
	w.Data = data
	return w
}
