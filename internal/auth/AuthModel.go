package auth

type IToken interface {
	CheckCorrectToken() error
}

type Token struct {

}

func NewToken() IToken {
	return &Token{}
}

func (t *Token) CheckCorrectToken() error {
	return nil
}