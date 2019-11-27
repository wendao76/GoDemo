package main

func main() {

}

type Response struct {
    Msg string
}

type IResponse interface {
    Message() string
}

func (resp *Response) Message() string {
	return resp.Msg
}

//仅返回需要用到的部分
func NewResp() IResponse {
	return &Response {Msg : "1234"}
}
