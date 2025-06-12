package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/services"
)

// PerfilController operations for Perfil
type PerfilController struct {
	beego.Controller
}

// URLMapping ...
func (c *PerfilController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Perfil
// @Param	body		body 	models.Perfil	true		"body for Perfil content"
// @Success 201 {object} models.Perfil
// @Failure 403 body is empty
// @router / [post]
func (c *PerfilController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Perfil by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Perfil
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PerfilController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Perfil
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Perfil
// @Failure 403
// @router / [get]
func (c *PerfilController) GetAll() {
	json_perfil, _ := services.Metodo_get_all("host_api2", "Usuarios/")
	json_perfil2, _ := services.ProcessarJson(json_perfil)
	perfil_json := json_perfil2["Consulta de id"]
	arreglo_perfil, _ := services.ConvertInterfaceToSliceMap(perfil_json)
	fmt.Println("hola", arreglo_perfil)

	var perfil []map[string]interface{}
	for i := range arreglo_perfil {
		perfil_parcial := map[string]interface{}{
			"Id":                arreglo_perfil[i]["Id"],
			"Nombre":            arreglo_perfil[i]["Nombre"],
			"Apellido":          arreglo_perfil[i]["Apellido"],
			"CorreoElectronico": arreglo_perfil[i]["CorreoElectronico"],
			"Telefono":          arreglo_perfil[i]["Telefono"],
			"FechaNacimiento":   arreglo_perfil[i]["FechaNacimiento"],
			"Ciudad":            arreglo_perfil[i]["Ciudad"],
			"Departamento":      arreglo_perfil[i]["Departamento"],
			"Pais":              arreglo_perfil[i]["Pais"],
			"tipo_documento":    arreglo_perfil[i]["IdTipoDocumentoTipoDocumento"].(map[string]interface{})["Nombre"],
			"NDocumento":        arreglo_perfil[i]["NDocumento"],
			"tipo_usuario":      arreglo_perfil[i]["IdRolRol"].(map[string]interface{})["Nombre"],
		}
		perfil = append(perfil, perfil_parcial)
	}
	c.Data["json"] = map[string]interface{}{"Succes": true, "Consulta de id": perfil}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Perfil
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Perfil	true		"body for Perfil content"
// @Success 200 {object} models.Perfil
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PerfilController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Perfil
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PerfilController) Delete() {

}
