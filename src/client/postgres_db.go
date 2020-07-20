package client

import (
	"fmt"
	"os"

	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/helper/logger"

	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/client/migrate"
	"github.com/andodeki/code/HA/api.gridbackendapp.com/src/client/ping"

	resterrors "github.com/andodeki/code/HA/api.gridbackendapp.com/src/helper/utils/rest_errors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

var (
	// conn *sqlx.DB

	port       string
	dbhost     string
	dbport     string
	dbuser     string
	dbpass     string
	dboptions  string
	dbdatabase string
)

const DEV = false

var (
	Conn databaseInterface = &databaseStruct{}
)

type databaseStruct struct {
	conn *sqlx.DB
}

func (c *databaseStruct) setConn(conn *sqlx.DB) {
	c.conn = conn
}
func (d *databaseStruct) Close() error {
	return d.conn.Close()
}

// Database is an interface for database
type databaseInterface interface {
	setConn(*sqlx.DB)
	GetClient() *sqlx.DB
	Close() error
	// CreateUser(ctx context.Context, user *usersDomain.User) error
	// GetUserByID(ctx context.Context, userID *usersDomain.UserID) (*usersDomain.User, error)
}

func Init() {

	//db_url := os.Getenv("DATABASE_URL")

	if DEV {
		conn, err := sqlx.Open("postgres", envInit(dbhost, dbport, dbuser, dbpass, dboptions, dbdatabase))
		dbInstance(conn, err)
	} else {
		conn, err := sqlx.Open("postgres", os.Getenv("DATABASE_URL"))
		dbInstance(conn, err)
	}
	// conn, err := sqlx.Open("postgres", dataSourceName)
	//conn, err := sqlx.Open("postgres", db_url)

}

func envInit(dbhost string, dbport string, dbuser string, dbpass string, dboptions string, dbdatabase string) string {

	// flag.StringVar(&dbhost, "dbhost", "localhost", "Set the port for the application")
	// flag.StringVar(&dbport, "dbport", "5432", "Set the port for the application")
	// flag.StringVar(&dbuser, "dbuser", "postgres", "Set the port for the application")
	// flag.StringVar(&dbpass, "dbpass", "password", "Set the port for the application")
	// flag.StringVar(&dboptions, "dboptions", "sslmode=disable", "Set the port for the application")
	// flag.StringVar(&dbdatabase, "dbdatabase", "testdb", "Set the port for the application")

	// flag.Parse()

	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	if host := os.Getenv("DB_HOST"); len(host) > 0 {
		dbhost = host
	}
	if database := os.Getenv("DB_DATABASE"); len(database) > 0 {
		dbdatabase = database
	}
	if user := os.Getenv("DB_USER"); len(user) > 0 {
		dbuser = user
	}
	if password := os.Getenv("DB_PASSWORD"); len(password) > 0 {
		dbpass = password
	}
	if port := os.Getenv("DB_PORT"); len(port) > 0 {
		dbport = port
	}

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s %s",
		dbhost, dbport, dbuser, dbdatabase, dbpass, dboptions,
	)
	return dataSourceName
}

func dbInstance(conn *sqlx.DB, err error) {
	resterrors.HandleErr(err)

	conn.SetMaxOpenConns(32)

	//Check if database is running
	if err := ping.WaitForDB(conn.DB); err != nil {
		errors.Wrap(err, "could not connect to database")
		logger.Info("could not connect to database")
	}

	if err := migrate.MigrateDb(conn.DB); err != nil {
		errors.Wrap(err, "could not migrate database")
		logger.Info("Migration Done Successfully")
	}

	logger.Info("database successfully configured")

	Conn.setConn(conn)
}

func (c *databaseStruct) GetClient() *sqlx.DB {
	return c.conn
}

// func GetSession() *sqlx.DB {
// 	d := &dbRepository{
// 		conn: Client,
// 	}
// 	return d
// }

// // ConnectDB is a func
// func Connect(host string, port string, user string, pass string, database string, options string) *gorm.DB {
// 	// db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=griffinsapp password=password sslmode=disable")
// 	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=griffinsapp password=password sslmode=disable")
// 	utils.HandleErr(err)
// 	return db
// }
