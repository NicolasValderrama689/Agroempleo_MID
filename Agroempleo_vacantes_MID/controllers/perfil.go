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
	Json_perfil1, _ := services.ProcessarJson(json_perfil)

	perfil_json := Json_perfil1["Consulta de id"]

	arreglo_perfil, _ := services.ConvertInterfaceToSliceMap(perfil_json)

	var Perfil_total []map[string]interface{}

	for i := range arreglo_perfil {

		Perfil_parcial := map[string]interface{}{
			"Nombre":            arreglo_perfil[i]["Nombre"],
			"Apellido":          arreglo_perfil[i]["Apellido"],
			"FNacimiento":       arreglo_perfil[i]["FechaNacimiento"],
			"NDocumento":        arreglo_perfil[i]["NDocumento"],
			"CorreoElectronico": arreglo_perfil[i]["CorreoElectronico"],
			"IdTipoDocumento":   arreglo_perfil[i]["IdTipoDocumentoTipoDocumento"].(map[string]interface{})["Nombre"],
			"IdTipoUsuario":     arreglo_perfil[i]["IdRolRol"].(map[string]interface{})["Nombre"],
			"pais":              arreglo_perfil[i]["Pais"],
			"departamento":      arreglo_perfil[i]["Departamento"],
			"ciudad":            arreglo_perfil[i]["Ciudad"],
		}
		Perfil_total = append(Perfil_total, Perfil_parcial)

	}

	json_laboral, _ := services.Metodo_get_all("host_api3", "laboral/")
	json_laboral1, _ := services.ProcessarJson(json_laboral)
	laboral_json := json_laboral1["usuarios consultados"]
	arreglo_laboral, _ := services.ConvertInterfaceToSliceMap(laboral_json)
	
	var Perfil_total_laboral []map[string]interface{}
	
	nombre := fmt.Sprintf("%v", jsonData["nombre"])
	apellido := fmt.Sprintf("%v", jsonData["apellido"])
	jsonData["nombre_completo"] = nombre + " " + apellido
	
	for i := range arreglo_laboral {

		Perfil_parcial_laboral := map[string]interface{}{
	   		"cargo": Perfil_parcial_laboral["Cargo"],
      		"empresa": Perfil_parcial_laboral["Empresa"],
      		"periodo": "2021 - 2023"


		}
	}

	json_academico, _ := services.Metodo_get_all("host_api4", "buscador/")
	json_academico1, _ := services.ProcessarJson(json_academico)
	academico_json := json_academico1["usuarios consultados"]
	arreglo_academico, _ := services.ConvertInterfaceToSliceMap(academico_json)

	c.Data["json"] = map[string]interface{}{"Succes": true, "Status": 200, "Message": "Perfil existosa", "Data": Perfil_total}
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
