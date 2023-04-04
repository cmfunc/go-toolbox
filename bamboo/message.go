package bamboo

type Message interface {
	GetTraceID() string
	GetType() string
	GetBody() interface{}
	GetHeader() interface{}
}

type CommonMsg struct {
	TraceID string      `json:"trace_id"`
	Type    string      `json:"type"`
	Body    interface{} `json:"body"`
	Header  interface{} `json:"header"`
}

func (m *CommonMsg) GetTraceID() string {
	if m == nil {
		return ""
	}
	return m.TraceID
}
func (m *CommonMsg) GetType() string {
	if m == nil {
		return ""
	}
	return m.Type
}
func (m *CommonMsg) GetBody() interface{} {
	if m == nil {
		return nil
	}
	return m.Body
}
func (m *CommonMsg) GetHeader() interface{} {
	if m == nil {
		return nil
	}
	return m.Header
}
