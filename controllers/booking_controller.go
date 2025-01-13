package controllers

import (
    "encoding/json"
    "net/http"
    "rental_api/models"
    // "rental_api/services"
    "github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
    web.Controller
    bookingService *services.BookingService
}

// @Title List Properties
// @Description Get all properties
// @Success 200 {object} models.RentalProperty
// @router / [get]
func (c *PropertyController) ListProperties() {
    properties, err := models.GetAllProperties()
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.Output.SetStatus(http.StatusInternalServerError)
        c.ServeJSON()
        return
    }
    c.Data["json"] = properties
    c.ServeJSON()
}

// @Title Get Property
// @Description Get property by ID
// @Param id path int true "Property ID"
// @Success 200 {object} models.RentalProperty
// @router /:id [get]
func (c *PropertyController) GetProperty() {
    id, _ := c.GetInt(":id")
    property, err := models.GetPropertyByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.Output.SetStatus(http.StatusNotFound)
        c.ServeJSON()
        return
    }
    c.Data["json"] = property
    c.ServeJSON()
}

// @Title Create Property
// @Description Create new property
// @Param body body models.RentalProperty true "Property information"
// @Success 201 {object} models.RentalProperty
// @router / [post]
func (c *PropertyController) CreateProperty() {
    var property models.RentalProperty
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &property); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.Output.SetStatus(http.StatusBadRequest)
        c.ServeJSON()
        return
    }

    if err := property.Create(); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.Output.SetStatus(http.StatusInternalServerError)
        c.ServeJSON()
        return
    }

    c.Data["json"] = property
    c.Ctx.Output.SetStatus(http.StatusCreated)
    c.ServeJSON()
}

// @Title Update Property
// @Description Update property by ID
// @Param id path int true "Property ID"
// @Param body body models.RentalProperty true "Property information"
// @Success 200 {object} models.RentalProperty
// @router /:id [put]
func (c *PropertyController) UpdateProperty() {
    id, _ := c.GetInt(":id")
    var property models.RentalProperty
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &property); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.Output.SetStatus(http.StatusBadRequest)
        c.ServeJSON()
        return
    }

    property.ID = id
    if err := property.Update(); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.Output.SetStatus(http.StatusInternalServerError)
        c.ServeJSON()
        return
    }

    c.Data["json"] = property
    c.ServeJSON()
}

// @Title Delete Property
// @Description Delete property by ID
// @Param id path int true "Property ID"
// @Success 204 {string} string
// @router /:id [delete]
func (c *PropertyController) DeleteProperty() {
    id, _ := c.GetInt(":id")
    if err := models.DeleteProperty(id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.Output.SetStatus(http.StatusInternalServerError)
        c.ServeJSON()
        return
    }
    c.Ctx.Output.SetStatus(http.StatusNoContent)
    c.ServeJSON()
}

// @Title Get Property Details
// @Description Get property details by ID
// @Param id path int true "Property ID"
// @Success 200 {object} models.PropertyDetails
// @router /:id/details [get]
func (c *PropertyController) GetPropertyDetails() {
    id, _ := c.GetInt(":id")
    details, err := models.GetPropertyDetailsByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.Ctx.Output.SetStatus(http.StatusNotFound)
        c.ServeJSON()
        return
    }
    c.Data["json"] = details
    c.ServeJSON()
}