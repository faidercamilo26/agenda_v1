package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:CorreoElectronicoController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:CorreoElectronicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:CorreoElectronicoController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:CorreoElectronicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:CorreoElectronicoController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:CorreoElectronicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:CorreoElectronicoController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:CorreoElectronicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:CorreoElectronicoController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:CorreoElectronicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:PersonaController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:PersonaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:TelefonoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:TelefonoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:TelefonoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:TelefonoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:TelefonoController"] = append(beego.GlobalControllerRouter["github.com/faidercamilo26/agenda_v1/controllers:TelefonoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
