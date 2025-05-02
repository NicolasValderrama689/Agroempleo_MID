package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/services"
)

// VacantesController operations for Vacantes
type VacantesController struct {
	beego.Controller
}

// URLMapping ...
func (c *VacantesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Vacantes
// @Param	body		body 	models.Vacantes	true		"body for Vacantes content"
// @Success 201 {object} models.Vacantes
// @Failure 403 body is empty
// @router / [post]
func (c *VacantesController) Post() {
	fmt.Println("Metodo post")

}

// GetOne ...
// @Title GetOne
// @Description get Vacantes by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Vacantes
// @Failure 403 :id is empty
// @router /:id [get]
func (c *VacantesController) GetOne() {
	fmt.Println("metodo get by id")

}

// GetAll ...
// @Title GetAll
// @Description get Vacantes
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Vacantes
// @Failure 403
// @router / [get]
func (c *VacantesController) GetAll() {
	fmt.Println("metodo Get all")
	Json_vacantes, _ := services.Metodo_get_all("host_api", "Vacantes")
	Json_vacantes2, _ := services.Metodo_get_all("host_api2", "Usuarios")

	fmt.Println("Este es el valor de Json_vacante2 en byte", Json_vacantes2)
	Json_procesado_vacantes, _ := services.ProcessarJson(Json_vacantes)
	Json_procesado_vacantes2, _ := services.ProcessarJson(Json_vacantes2)

	//fmt.Println("Este es el valor de Json_vacante2 en json", Json_procesado_vacantes2)

	Vacantes_json := Json_procesado_vacantes["Consulta de id"]
	Vacantes_json1 := Json_procesado_vacantes2["usuarios consultados"]

	Arreglo_Map_Vacantes_final, _ := services.ConvertInterfaceToSliceMap(Vacantes_json)
	Arreglo_Map_Usuario1, _ := services.ConvertInterfaceToSliceMap(Vacantes_json1)
	fmt.Println("usuarios basedata", Arreglo_Map_Usuario1[0]["Nombre"])

	// Arreglo_Map_Vacantes_final := append(Arreglo_Map_Vacantes, Arreglo_Map_Usuario1...)

	var Resultado_total []map[string]interface{}
	for i := range Arreglo_Map_Vacantes_final {

		fmt.Println("valor de json solo", Arreglo_Map_Vacantes_final[i])
		usuarioID := Arreglo_Map_Vacantes_final[i]["Id_usuarios"]
		usuarioID_string := fmt.Sprintf("%v", usuarioID)

		url := "http://localhost:8080/v1/Usuarios/"


		final,_ := services.Metodo_getid(url, usuarioID_string)
		Json_usuario, _ := services.ProcessarJson(final)
		Arreglo_Map_Usuario, _ := services.ConvertInterfaceToSliceMap(Json_usuario)
		fmt.Println("usuario",Arreglo_Map_Usuario)
		

		Resultado_parcial := map[string]interface{}{
			"TituloPuesto":        Arreglo_Map_Vacantes_final[i]["TituloPuesto"],
			"DescripcionTrabajo":  Arreglo_Map_Vacantes_final[i]["DescripcionTrabajo"],
			"Cargo":               Arreglo_Map_Vacantes_final[i]["Cargo"],
			"Salario":             Arreglo_Map_Vacantes_final[i]["Salario"],
			"Horario":             Arreglo_Map_Vacantes_final[i]["Horario"],
			"Modalidad":           Arreglo_Map_Vacantes_final[i]["Modalidad"],
			"NivelRequerido":      Arreglo_Map_Vacantes_final[i]["NivelRequerido"],
			"ExperienciaRequrida": Arreglo_Map_Vacantes_final[i]["ExperienciaRequrida"],
			"NumeroVacantes":      Arreglo_Map_Vacantes_final[i]["NumeroVacantes"],
			"Activo":              Arreglo_Map_Vacantes_final[i]["Activo"],
			"TipoEmpleo":          Arreglo_Map_Vacantes_final[i]["IdtipoempleoTipodeempleo"].(map[string]interface{})["Nombre"],
			"Ciudad":              Arreglo_Map_Vacantes_final[i]["Idciudadtrabajociudad"].(map[string]interface{})["Nombre"],
			"publicado_por":       Arreglo_Map_Vacantes_final[i]["Id_usuarios"],
		}
		Resultado_total = append(Resultado_total, Resultado_parcial)

	}

	c.Data["json"] = map[string]interface{}{"Succes": true, "Status": 200, "Message": "Consulta existosa", "Data": Resultado_total}
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Vacantes
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Vacantes	true		"body for Vacantes content"
// @Success 200 {object} models.Vacantes
// @Failure 403 :id is not int
// @router /:id [put]
func (c *VacantesController) Put() {
	fmt.Println("metodo put")

}

// Delete ...
// @Title Delete
// @Description delete the Vacantes
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *VacantesController) Delete() {
	fmt.Println("metodo delete")
}
