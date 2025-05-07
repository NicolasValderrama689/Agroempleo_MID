package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:CiudadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:CiudadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:CiudadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:CiudadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:CiudadController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:CiudadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:RegistroController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:RegistroController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:RegistroController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:RegistroController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:RegistroController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:RegistroController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:RegistroController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:RegistroController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:RegistroController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:RegistroController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:VacantesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:VacantesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:VacantesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:VacantesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:VacantesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:VacantesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:VacantesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:VacantesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:VacantesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/Agroempleo_MID/Agroempleo_vacantes_MID/controllers:VacantesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
