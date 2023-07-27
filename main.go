package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"salt-academy_learn_week2/domain/repository"
	"salt-academy_learn_week2/model"
	"salt-academy_learn_week2/model/mapper"
	repo "salt-academy_learn_week2/model/repository"
	"salt-academy_learn_week2/pkg/database/mysql"
	"salt-academy_learn_week2/presenter/response"
	"salt-academy_learn_week2/request"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

/*
= mysql.InitMysqlDB()
= repository.NewMahasiswaRepositoryMysql(mysqlConnection)
*/
var (
	ctx             = context.Background()
	mysqlConnection *sql.DB
	repositoryMysql repository.MahasiswaTemplate
)

func main() {
	_ = godotenv.Load()

	mysqlConnection = mysql.InitMysqlDB()
	repositoryMysql = repo.NewMahasiswaRepositoryMysql(mysqlConnection)

	fmt.Println(os.Getenv("MYSQL_DB_NAME"))

	err := mysqlConnection.Ping()
	if err != nil {
		panic(err.Error())
		//fmt.Println(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/list-mahasiswa", GetListMahasiwa).Methods(http.MethodGet)
	r.HandleFunc("/addmahasiswa", CreateMahasiswa).Methods(http.MethodPost)
	http.ListenAndServe(":8080", r)
}

func CreateMahasiswa(w http.ResponseWriter, r *http.Request) {
	var req request.BodyMahasiswa
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN CREATE MAHASISWA",
		}})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseStatusError)
		return
	}

	DBirth, err := time.Parse("2006-01-02", req.BirthDay)
	if err != nil {
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "DATE NOT MATCH",
		}})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseStatusError)
		return
	}

	err = repositoryMysql.AddMahasiswa(ctx, model.Mahasiswa{
		Name:      req.Nama,
		Gender:    req.Gender,
		BirthDate: DBirth,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN CREATE MAHASISWA",
		}})
		w.Header().Set("Content-Type", "application/json")

		w.Write(responseStatusError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetListMahasiwa(w http.ResponseWriter, router *http.Request) {

	listMahasiswa, err := GetListMahasiswaUC(ctx, repositoryMysql)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN GET LIST MAHASISWA",
		}})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseStatusError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(mapper.MapperCollectionModelToCollectionJSON(listMahasiswa))
	return
}

func GetListMahasiswaUC(ctx context.Context, repo repository.MahasiswaTemplate) ([]*model.Mahasiswa, error) {
	listMahasiswa, err := repositoryMysql.GetListMahasiswa(ctx)
	if err != nil {
		return nil, err
	}
	return listMahasiswa, nil
}
