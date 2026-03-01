package postgres

var (
	client *Client
)

// NewClient instantiate a new postegres client
func NewClient(params ConnParams) *Client {
	client = &Client{
		ConnParams: params,
	}
	return client
}

// GetClient get a postgres client
func GetClient() *Client {
	return client
}

type StorageType int

const (
	Mongo StorageType = iota + 1
	Postgres
)

type ConnParams struct {
	Type     StorageType
	Host     string
	Port     uint
	Database string

	AuthWithUserAndPassword bool
	UserName                string
	UserPassword            string
}
