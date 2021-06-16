package controller

import (
	"echo-tutor/config"
	//"echo-tutor/config"
	"echo-tutor/model"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
)

var pool = config.NewConnectionPools()

func GetMhs(c echo.Context) error {
	db := config.GetDBEngine()
	mhs, err := db.Query("SELECT * FROM mahasiswa")
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	var mahasiswas []model.Mahasiswa
	for mhs.Next() {
		var mahasiswa model.Mahasiswa
		mhs.Scan(&mahasiswa.ID, &mahasiswa.Jurusan, &mahasiswa.Name, &mahasiswa.NoTlp, &mahasiswa.Nim)
		mahasiswas = append(mahasiswas, mahasiswa)
	}

	return c.JSON(http.StatusOK, mahasiswas)
}

func GetMhsByNim(c echo.Context) error {
	connection := pool.Get()
	db := config.GetDBEngine()
	nim := c.Param("nim")

	sqlStatement := `SELECT * FROM mahasiswa WHERE nim=$1;`
	mahas, err := db.Query(sqlStatement, nim)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	mhs := model.Mahasiswa{}
	for mahas.Next(){
		var mahasiswa model.Mahasiswa
		mahas.Scan(&mahasiswa.ID, &mahasiswa.Jurusan, &mahasiswa.Name, &mahasiswa.NoTlp, &mahasiswa.Nim)
		mhs.Nim = nim
		mhs.ID=mahasiswa.ID
		mhs.Name=mahasiswa.Name
		mhs.Jurusan = mahasiswa.Jurusan
		mhs.NoTlp = mahasiswa.NoTlp
	}

	result, err := redis.String(connection.Do("GET", nim))
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	logrus.Info(result)

	defer db.Close()
	return c.JSON(http.StatusOK, mhs)
}

func CreateMhs(c echo.Context) error {
	connection := pool.Get()
	db := config.GetDBEngine()

	mhs := new(model.Mahasiswa)
	if err := c.Bind(mhs); err != nil {
		logrus.Error(err)
		panic(err)
	}

	sqlStatement := `insert into mahasiswa (id, jurusan, name, no_tlp, nim) values ($1, $2, $3, $4, $5);`

	mhs.ID = uuid.NewString()
	_, err := db.Exec(sqlStatement, mhs.ID, mhs.Jurusan, mhs.Name, mhs.NoTlp, mhs.Nim)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	mahas, err := json.Marshal(mhs)

	_, err = connection.Do("SET", mhs.Nim, mahas)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	defer db.Close()
	return c.JSON(http.StatusCreated, mhs)
}

func UpdateMhs(c echo.Context) error {
	connection := pool.Get()
	db := config.GetDBEngine()
	nim := c.Param("nim")
	mhs := new(model.Mahasiswa)
	if err := c.Bind(mhs); err != nil {
		logrus.Error(err)
		panic(err)
	}

	getdata := `SELECT * FROM mahasiswa WHERE nim=$1;`
	mahas, err := db.Query(getdata, nim)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	var mahasiswa model.Mahasiswa
	for mahas.Next(){
		mahas.Scan(&mahasiswa.ID, &mahasiswa.Jurusan, &mahasiswa.Name, &mahasiswa.NoTlp, &mahasiswa.Nim)
	}

	sqlStatement := `UPDATE mahasiswa SET  jurusan=$1, name=$2, no_tlp=$3 WHERE nim=$4;`

	_, err = db.Query(sqlStatement, mhs.Jurusan, mhs.Name, mhs.NoTlp, nim)
	mhs.Nim = nim
	mhs.ID=mahasiswa.ID
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	mahasis, err := json.Marshal(mhs)

	_, err = connection.Do("SET", mhs.Nim, mahasis)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	defer db.Close()
	return c.JSON(http.StatusOK, mhs)
}

func DeleteMhs(c echo.Context) error {
	connection := pool.Get()
	db := config.GetDBEngine()
	nim := c.Param("nim")

	sqlStatement := `DELETE FROM mahasiswa WHERE nim=$1;`

	_, err := db.Query(sqlStatement, nim)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	_, err = connection.Do("SET", nim, nil)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	defer db.Close()
	return c.JSON(http.StatusOK, nim+" DELETED")
}