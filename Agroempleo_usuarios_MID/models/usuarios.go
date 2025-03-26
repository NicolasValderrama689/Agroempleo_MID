package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Rol struct {
	Id     int    `orm:"column(Id);pk"`
	Nombre string `orm:"column(Nombre)"`
}

type Contraseñas struct {
	Id                int       `orm:"column(id_contraseñas);pk;auto"`
	Contraseña        string    `orm:"column(contraseña)"`
	Activo            bool      `orm:"column(activo)"`
	FechaCreacion              time.Time     `orm:"column(Fecha_creacion);type(timestamp with time zone);auto_now_add"`
	FechaModificacion          time.Time     `orm:"column(Fecha_modificacion);type(timestamp with time zone);auto_now"`
}

type Identificacion struct {
	Id                            int            `orm:"column(Id_identificacion);pk;auto"`
	Numero                        int            `orm:"column(Numero)"`
	IdIdentificacionTipoDocumento *TipoDocumento `orm:"column(id_identificacion_tipo_documento);rel(fk)"`
	Activo                        bool           `orm:"column(activo)"`
	FechaCreacion              time.Time     `orm:"column(Fecha_creacion);type(timestamp with time zone);auto_now_add"`
	FechaModificacion          time.Time     `orm:"column(Fecha_modificacion);type(timestamp with time zone);auto_now"`
}


type Usuarios struct {
	Id                             int           `orm:"column(Id_Usuario);pk;auto"`
	Nombre                         string        `orm:"column(Nombre)"`
	IdRolRol                       *Rol          `orm:"column(id_rol_rol);rel(fk)"`
	CorreoElectronico                         string        `orm:"column(CorreoElectronico)"`
	IdIdentificacionIdentificacion *Identificacion        `orm:"column(Id_identificacion_Identificacion);rel(fk)"`
	NumeroTelefono                 string        `orm:"column(Numero_Telefono)"`
	FotoPerfil                     string        `orm:"column(Foto_Perfil);type(text);null"`
	FechaCreacion                  time.Time     `orm:"column(Fecha_Creacion);type(timestamp with time zone);auto_now_add"`
	FechaModificacion              time.Time     `orm:"column(Fecha_Modificacion);type(timestamp with time zone);auto_now"`
	Activo                         bool          `orm:"column(Activo)"`
	IdContraseñasContraseñas     *Contraseñas `orm:"column(Id_Credenciales);rel(fk)"`
}

// Registrar el modelo en Beego ORM
func init() {
	orm.RegisterModel(new(Usuarios), new(Rol), new(Identificacion))
}
