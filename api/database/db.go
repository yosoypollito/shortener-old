package database

import (
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
	"shortener/api/config"
)

func Getbase(name string) (*base.Base, error) {

	key := config.EnvVariable("DETA_PROJECT_KEY")

	d, err := deta.New(deta.WithProjectKey(key))

	if err != nil {
		return nil, err
	}

	db, err := base.New(d, name)

	if err != nil{
		return nil, err
	}

	return db, nil
}

func GetFromBase[K any](name string, query base.Query) ([]*K, error) {

	db, err := Getbase(name)

	if err != nil {
		return nil, err
	}

	var results []*K

	_, err = db.Fetch(&base.FetchInput{
		Q: query,
		Dest: &results,
	})

	if err != nil {
		return results, err
	}

	return results, nil
}
