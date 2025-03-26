package services

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/sena_2824182/Agroempleo_MID/Agroempleo_usuarios_MID/models"
)

// Desactivar usuario 
func DesactivarUsuario(idUsuario int) error {
	o := orm.NewOrm()

	// Iniciar transacción
	err := o.Begin()
	if err != nil {
		return fmt.Errorf("error al iniciar la transacción: %v", err)
	}

	// Desactivar el usuario
	_, err = o.QueryTable(new(models.Usuarios)).Filter("Id", idUsuario).Update(orm.Params{"Activo": false})
	if err != nil {
		o.Rollback()
		return fmt.Errorf("error al desactivar usuario %d: %v", idUsuario, err)
	}

	return nil
}

// AutoEliminarUsuario - Permite que un usuario elimine su cuenta 
func AutoEliminarUsuario(idUsuario int) error {
	return DesactivarUsuario(idUsuario) // Reutilizamos la función
}
