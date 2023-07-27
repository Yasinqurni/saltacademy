package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"salt-academy_learn_week2/domain/entity"
	"salt-academy_learn_week2/domain/repository"
	"salt-academy_learn_week2/model"
	"salt-academy_learn_week2/model/mapper"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

type mahasiswaInteractor struct {
	conn *sql.DB
}

func NewMahasiswaRepositoryMysql(connMysql *sql.DB) repository.MahasiswaTemplate {
	return &mahasiswaInteractor{conn: connMysql}
}

func (mhs *mahasiswaInteractor) GetMahasiswa(ctx context.Context) (*model.Mahasiswa, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s LIMIT 1 ORDER BY created_at DESC`, model.Mahasiswa{}.GetTableName())
	opts := &dbq.Options{SingleResult: true, ConcreteStruct: model.Mahasiswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	result := dbq.MustQ(ctx, mhs.conn, stmt, opts)

	if result != nil {
		return result.(*model.Mahasiswa), nil
	}

	return nil, errors.New("DATA EMPTY")
}

func (mhs *mahasiswaInteractor) GetListMahasiswa(ctx context.Context) ([]*entity.Mahasiswa, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, model.Mahasiswa{}.GetTableName())
	opts := &dbq.Options{SingleResult: false, ConcreteStruct: model.Mahasiswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	result := dbq.MustQ(ctx, mhs.conn, stmt, opts)

	if result != nil || len(result.([]*model.Mahasiswa)) == 0 {

		return mapper.MapperListModelToEntityMahasiswa(result.([]*model.Mahasiswa)), nil

	}

	return nil, errors.New("DATA EMPTY")
}

func (mhs *mahasiswaInteractor) AddMahasiswa(ctx context.Context, Mahasiswa model.Mahasiswa) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	stmt := fmt.Sprintf(`INSERT INTO %s (nama, tanggal_lahir, jenis_kelamin) VALUES (?, ?, ?)`, model.Mahasiswa{}.GetTableName())
	ops := &dbq.Options{SingleResult: true, ConcreteStruct: model.Mahasiswa{}, DecoderConfig: dbq.StdTimeConversionConfig()}
	_, err := dbq.E(
		ctx,
		mhs.conn,
		stmt,
		ops,
		Mahasiswa.Name,
		Mahasiswa.BirthDate,
		Mahasiswa.Gender,
	)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// func isZero(v interface{}) bool {
// 	if v == nil {
// 		return true
// 	}
// 	return reflect.ValueOf(v).IsZero()
// }
