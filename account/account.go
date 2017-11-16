package account

type AccountType int

const (
	Password AccountType = 1 << iota
)

type Account struct {
	Type     AccountType
	Name     string
	Username string
	Password string
}

func MakeSecret(username, password string) []byte {
	str := username + ":" + password
	data := []byte(str)
	length := len(data)
	if length < 32 {
		diff := 32 - length
		for i := diff; i > 0; i-- {
			data = append(data, 0x0)
		}
	}
	return data
}
