package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/services"
)

// PostulacionesController operations for Postulaciones
type PostulacionesController struct {
	beego.Controller
}

// URLMapping ...
func (c *PostulacionesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Postulaciones
// @Param	body		body 	models.Postulaciones	true		"body for Postulaciones content"
// @Success 201 {object} models.Postulaciones
// @Failure 403 body is empty
// @router / [post]
func (c *PostulacionesController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Postulaciones by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Postulaciones
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PostulacionesController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Postulaciones
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Postulaciones
// @Failure 403
// @router / [get]
func (c *PostulacionesController) GetAll() {
	json_postulaciones, _ := services.Metodo_get_all("host_api5", "postulaciones")
	json_postulaciones2, _ := services.ProcessarJson(json_postulaciones)
	postulaciones_json := json_postulaciones2["usuarios consultados"]
	arreglo_postulaciones, _ := services.ConvertInterfaceToSliceMap(postulaciones_json)

	var postulaciones []map[string]interface{}
	for i := range arreglo_postulaciones {

		usuarioID := arreglo_postulaciones[i]["IdUsuarios"]
		usuarioID_string := fmt.Sprintf("%v", usuarioID)

		endpoint := "Usuarios/"+ usuarioID_string
	

		final,_ := services.Metodo_getid("host_api2", endpoint)
		Json_usuario, _ := services.ProcessarJson(final)
		nombre_usuario := Json_usuario["Consulta de id"].(map[string]interface{})["Nombre"]

		empleoID := arreglo_postulaciones[i]["IdEmpleo"]
		empleoID_string := fmt.Sprintf("%v", empleoID)

		endpoint2 := "Vacantes/"+ empleoID_string

		final2,_ := services.Metodo_getid("host_api", endpoint2)
		Json_vacante, _ := services.ProcessarJson(final2)
		titulo_puesto := Json_vacante["Consulta de id"].(map[string]interface{})["TituloPuesto"]
		IdEmpleador := Json_vacante["Consulta de id"].(map[string]interface{})["Id_usuarios"]
		usuarioID_string1 := fmt.Sprintf("%v", IdEmpleador)
		endpoint3 := "Usuarios/"+ usuarioID_string1
		final3,_ := services.Metodo_getid("host_api2", endpoint3)
		json_empleador, _ := services.ProcessarJson(final3)
		nombre_empleador := json_empleador["Consulta de id"].(map[string]interface{})["Nombre"]


		

		postulacion_parcial := map[string]interface{}{
			"Id":                arreglo_postulaciones[i]["Id"],
			"IdAspirante":            arreglo_postulaciones[i]["IdUsuarios"],
			"IdEmpleador":        IdEmpleador,
			"NombrePostulado":            nombre_usuario,
			"NombrePostulado1":        nombre_empleador,
			"TituloPuesto":        titulo_puesto,
			"ArchivoCV":          arreglo_postulaciones[i]["SoporteCv"],

		}
		postulaciones = append(postulaciones, postulacion_parcial)
	}
	c.Data["json"] = map[string]interface{}{"Succes": true, "Status": 200, "Message": "Consulta existosa", "Data": postulaciones}
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Postulaciones
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Postulaciones	true		"body for Postulaciones content"
// @Success 200 {object} models.Postulaciones
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PostulacionesController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Postulaciones
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PostulacionesController) Delete() {

}
