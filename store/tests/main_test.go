package tests

import (
	"context"
	"fmt"
	"linkgen/pkg/tests"
	"linkgen/store/mysql"
	"log"
	"os"
	"testing"
)

var (
	mysqlURI   = ""
	storageDsn = ""
)

func setupMysqlStore(dsn string) error {
	mysqlLinkStore, err := mysql.New(dsn)
	if err != nil {
		return err
	}
	stores = append(stores, mysqlLinkStore)
	return nil
}

func TestMain(m *testing.M) {
	var code = 1
	defer func() { os.Exit(code) }()

	ctx := context.Background()

	mysqlContainer, err := tests.Setup(ctx, ContainerMySQLImage)
	if err != nil {
		log.Println("Provider Test Failed - Create MySQL Container", "Error", err)
		return
	}

	mysqlURI = mysqlContainer.URIS[0]

	storageDsn = fmt.Sprintf("%v:%v@tcp(%v)/linkgen?parseTime=true", "root", "secret", mysqlURI)

	if err = setupMysqlStore(storageDsn); err != nil {
		log.Println("Provider Test Failed - Create MySQL Store", "Error", err)
		return
	}

	code = m.Run()
}
