package jdb

type DbAccesser interface {
	FetchOne()
	FetchAll()
	Update()
	Insert()
}

func getMaster(dbname string) {

}

func getSlave(dbname string) {

}
