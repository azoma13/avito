package configs

const (
	DateLayout  = "20060102"
	DefaultPort = "8080"
	WebDir      = "./web"
	host        = "localhost"
	port        = 5432
	user        = "postgres"
	password    = "pass"
	dbname      = "some-postgres"
)

var JwtKey = []byte("my_secret_key")

type Employee struct {
	ID           int         `json:"id"`
	Email        string      `json:"email"`
	Password     string      `json:"password"`
	Coins        int         `json:"coins"`
	Inventory    []Item      `json:"inventory"`
	CoinHistorys CoinHistory `json:"coinHistory"`
}

type Item struct {
	Merch    string `json:"merch"`
	Quantity int    `json:"quantity"`
}

type CoinHistory struct {
	Received []Transaction `json:"received"`
	Sent     []Transaction `json:"sent"`
}

type Transaction struct {
	FromUser string `json:"fromUser"`
	ToUser   string `json:"toUser"`
	Amount   int    `json:"amount"`
}
