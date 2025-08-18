package banco

import (
	"database/sql"
	"fmt"
)

func Conectar() (*sql.DB, error) {
	fmt.Print("Conectando ao banco de dados")
}
