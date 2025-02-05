package test

import (
	"fmt"
	"testing"

	db "github.com/cparamoPriv/libcore/database"
)

/*TestDataTable : Crea sql dinamicos*/
func TestDataTable(t *testing.T) {
	var (
		data db.DataTable
	)
	data.SetTable("sumcon")

	data.AddRow(db.StData{
		"campo1": 5,
		"campo2": 5,
		"campo3": 6,
	})
	data.AddIndex("campo1")
	data.AddIndex("campo2")
	rowi, _ := data.GenInserts()
	fmt.Printf("%v", rowi)
	rowu, _ := data.GenUpdates()
	fmt.Printf("%v", rowu)
	rowd, _ := data.GenDeletes()
	fmt.Printf("%v", rowd)
}

/*TestSqlLite : Se conecta a una base de datos  sqllite de configuracion*/
func TestSqlLite(t *testing.T) {
	var (
		conexion db.StConect
	)
	path := "config/sqllite.ini"
	t.Logf("Capturando path:%s", path)
	err := conexion.ConfigINI(path)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("Conexion:%s", conexion.Conexion.ToString())
	t.Logf("Probando...")
	t.Logf("prueba:%v", conexion.Test())
}

/*TestSqlLiteEncrip : Se conecta a una base de datos  sqllite de configuracion con clave aes256*/
func TestSqlLiteEncrip(t *testing.T) {
	var (
		conexion db.StConect
	)
	pass := "abc123"
	path := "config/sqllite.dbx"
	t.Logf("Capturando path:%s", path)
	err := conexion.ConfigDBX(path, pass)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("Conexion:%s", conexion.Conexion.ToString())
	t.Logf("Probando...")
	t.Logf("prueba:%v", conexion.Test())
}

/*TestSqlLite : crea una conexion encriptada*/
func TestCreateDB(t *testing.T) {
	var (
		conexion  db.StConect
		cnxencrip []byte
	)
	pass := "abc123"
	path := "config/sqllite.ini"
	t.Logf("Capturando path:%s", path)
	err := conexion.ConfigINI(path)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	cnxencrip, err = db.CreateDBConect(conexion.Conexion, pass)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("Data Encrip:%s", string(cnxencrip))
	conexion.ResetCnx()
	cnxdecrip, err := db.DecripConect(cnxencrip, pass)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	conexion.Conexion = cnxdecrip
	t.Logf("Conexion:%s", conexion.Conexion.ToString())
	t.Logf("Probando...")
	t.Logf("prueba:%v", conexion.Test())
}

/*TestPost : Se conecta a una base de datos  posgresql de configuracion*/
func TestPost(t *testing.T) {
	var (
		conexion db.StConect
	)
	path := "config/post.ini"
	t.Logf("Capturando path:%s", path)
	err := conexion.ConfigINI(path)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("Conexion:%s", conexion.Conexion.ToString())
	t.Logf("Probando...")
	t.Logf("prueba:%v", conexion.Test())
	t.Logf("prueba:%v", conexion.Test())
	t.Logf("prueba:%v", conexion.Test())
}

/*TestTestORAPost : Se conecta a una base de datos  ORACLE de configuracion*/
func TestORA(t *testing.T) {
	var (
		conexion db.StConect
	)
	path := "config/ora.ini"
	t.Logf("Capturando path:%s", path)
	err := conexion.ConfigINI(path)
	if err != nil {
		t.Errorf("Error:%s", err.Error())
	}
	t.Logf("Conexion:%s", conexion.Conexion.ToString())
	t.Logf("Probando...")
	t.Logf("prueba:%v", conexion.Test())
}
