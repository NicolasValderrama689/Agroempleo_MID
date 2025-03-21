package controllers

import (
	"github.com/astaxie/beego"
)

// UsuarioController operations for Usuario
type UsuarioController struct {
	beego.Controller
}

// URLMapping ...
func (c *UsuarioController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Usuario
// @Param	body		body 	models.Usuario	true		"body for Usuario content"
// @Success 201 {object} models.Usuario
// @Failure 403 body is empty
// @router / [post]
func (c *UsuarioController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Usuario by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usuario
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsuarioController) GetOne() {
	fmt.println("funcion getone")
	fmt.println("funcion getone")

	type User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	
	// Simulando una base de datos de usuarios (en memoria)
	var users = []User{
		{ID: 1, Username: "johndoe", Email: "johndoe@example.com"},
		{ID: 2, Username: "janedoe", Email: "janedoe@example.com"},
	}
	
	func main() {
		// Crear un router de Gin
		r := gin.Default()
	
		// Ruta para obtener un usuario por ID
		r.GET("/users/:id", getUser)
	
		// Correr la API en el puerto 8080
		r.Run(":8080")
	}
	
	// Controlador para obtener un usuario por ID
	func getUser(c *gin.Context) {
		// Obtener el ID del parámetro de la URL
		idParam := c.Param("id")
	
		// Convertir el ID de string a int
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID debe ser un número válido"})
			return
		}
	
		// Buscar el usuario con el ID especificado
		for _, user := range users {
			if user.ID == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}
	
		// Si no se encuentra el usuario
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
	}

}

// GetAll ...
// @Title GetAll
// @Description get Usuario
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Usuario
// @Failure 403
// @router / [get]
func (c *UsuarioController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Usuario
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Usuario	true		"body for Usuario content"
// @Success 200 {object} models.Usuario
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsuarioController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Usuario
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsuarioController) Delete() {

}
