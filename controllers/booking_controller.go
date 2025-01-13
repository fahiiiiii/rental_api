package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"property-listing/conf"
	"property-listing/models"

	"github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

type BookingController struct {
	web.Controller
	uniqueCountries map[string]bool
	uniqueCities    map[string]bool
	countryCities   map[string][]string
	cityProperties  map[string][]string
	mutex           sync.Mutex
	rateLimiter     *rate.Limiter
	rapidAPIKey     string
}

// Constructor method
func NewBookingController() *BookingController {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	rapidAPIKey := os.Getenv("RAPIDAPI_KEY")
	if rapidAPIKey == "" {
		log.Fatalf("RAPIDAPI_KEY is not set in the environment")
	}

	return &BookingController{
		uniqueCountries: make(map[string]bool),
		uniqueCities:    make(map[string]bool),
		countryCities:   make(map[string][]string),
		cityProperties:  make(map[string][]string),
		rateLimiter:     rate.NewLimiter(rate.Every(12*time.Second), 1),
		rapidAPIKey:     rapidAPIKey,
	}
}

// CRUD Operations for RentalProperty

// Create a new property
func (c *BookingController) Create() {
	var property models.RentalProperty
	
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &property)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	if property.PropertyName == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Property name is required"}
		c.ServeJSON()
		return
	}

	result := conf.DB.Create(&property)
	if result.Error != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to create property"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = property
	c.ServeJSON()
}

// Get a specific property
func (c *BookingController) Get() {
	propertyIdStr := c.Ctx.Input.Param(":id")
	propertyId, err := strconv.Atoi(propertyIdStr)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid property ID"}
		c.ServeJSON()
		return
	}

	var property models.RentalProperty
	result := conf.DB.First(&property, propertyId)
	if result.Error != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": "Property not found"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = property
	c.ServeJSON()
}

// Update a property
func (c *BookingController) Put() {
	propertyIdStr := c.Ctx.Input.Param(":id")
	propertyId, err := strconv.Atoi(propertyIdStr)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid property ID"}
		c.ServeJSON()
		return
	}

	var updatedProperty models.RentalProperty
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &updatedProperty)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	var existingProperty models.RentalProperty
	result := conf.DB.First(&existingProperty, propertyId)
	if result.Error != nil {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": "Property not found"}
		c.ServeJSON()
		return
	}

	result = conf.DB.Model(&existingProperty).Updates(updatedProperty)
	if result.Error != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to update property"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = existingProperty
	c.ServeJSON()
}

// Delete a property
func (c *BookingController) Delete() {
	propertyIdStr := c.Ctx.Input.Param(":id")
	propertyId, err := strconv.Atoi(propertyIdStr)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid property ID"}
		c.ServeJSON()
		return
	}

	result := conf.DB.Delete(&models.RentalProperty{}, propertyId)
	if result.Error != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to delete property"}
		c.ServeJSON()
		return
	}

	if result.RowsAffected == 0 {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = map[string]string{"error": "Property not found"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(204)
	c.ServeJSON()
}

// List properties with optional filtering
func (c *BookingController) List() {
	city := c.GetString("city")
	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("pageSize", 10)

	var properties []models.RentalProperty
	query := conf.DB

	if city != "" {
		query = query.Where("city = ?", city)
	}

	offset := (page - 1) * pageSize
	result := query.Offset(offset).Limit(pageSize).Find(&properties)
	
	if result.Error != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to fetch properties"}
		c.ServeJSON()
		return
	}

	var total int64
	conf.DB.Model(&models.RentalProperty{}).Count(&total)

	response := map[string]interface{}{
		"properties": properties,
		"page":       page,
		"pageSize":   pageSize,
		"total":      total,
	}

	c.Data["json"] = response
	c.ServeJSON()
}



func (c *BookingController) makeRateLimitedRequest(req *http.Request) (*http.Response, error) {
	err := c.rateLimiter.Wait(context.Background())
	if err != nil {
		return nil, fmt.Errorf("rate limiter error: %v", err)
	}


    func (c *BookingController) fetchCities(query string) ([]models.City, error) {
	
        //!using rate limit
        apiURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/auto-complete?query=%s", query)
        
        req, err := http.NewRequest("GET", apiURL, nil)
        if err != nil {
            return nil, fmt.Errorf("error creating request: %v", err)
        }
    
        req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
        req.Header.Add("x-rapidapi-key", c.rapidAPIKey) // Use the stored RapidAPI key
        // req.Header.Add("x-rapidapi-key", "79d933f58amsh0baa13f673b03f0p16d4a2jsnb299a967d295")
    
        resp, err := c.makeRateLimitedRequest(req)
        if err != nil {
            return nil, fmt.Errorf("error sending request: %v", err)
        }
        defer resp.Body.Close()
    
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            return nil, fmt.Errorf("error reading response body: %v", err)
        }
    
        if resp.StatusCode != http.StatusOK {
            return nil, fmt.Errorf("API request failed with status code: %d, body: %s", 
                resp.StatusCode, string(body))
        }
    
        var citiesResp struct {
            Data []models.City `json:"data"`
        }
        err = json.Unmarshal(body, &citiesResp)
        if err != nil {
            return nil, fmt.Errorf("error parsing JSON: %v", err)
        }
    
        return citiesResp.Data, nil
    }
    func (c *BookingController) fetchPropertiesForCity(cityName, country string) ([]models.Property, error) {
        uniqueProperties := make(map[string]models.Property)
        searchQueries := []string{
            cityName,
            fmt.Sprintf("%s hotels", cityName),
            fmt.Sprintf("%s accommodation", cityName),
        }
    
        for _, query := range searchQueries {
            encodedQuery := url.QueryEscape(query)
            apiURL := fmt.Sprintf("https://booking-com18.p.rapidapi.com/stays/auto-complete?query=%s", encodedQuery)
    
            properties, err := c.fetchPropertyData(apiURL)
            if err != nil {
                continue // Skip errors and proceed with other queries
            }
    
            for _, prop := range properties {
                if prop.DestID != "" { // Ensure destId exists
                    uniqueProperties[prop.DestID] = prop
                }
            }
        }
    
        result := make([]models.Property, 0, len(uniqueProperties))
        for _, prop := range uniqueProperties {
            result = append(result, prop)
        }
    
        return result, nil
    }
    
    func (c *BookingController) fetchPropertyData(apiURL string) ([]models.Property, error) {
        req, err := http.NewRequest("GET", apiURL, nil)
        if err != nil {
            return nil, err
        }
    
        req.Header.Add("x-rapidapi-host", "booking-com18.p.rapidapi.com")
        req.Header.Add("x-rapidapi-key", c.rapidAPIKey) // Use the stored RapidAPI key
    
        resp, err := c.makeRateLimitedRequest(req)
        if err != nil {
            return nil, err
        }
        defer resp.Body.Close()
    
        body, err := io.ReadAll(resp.Body)
        if err != nil {
            return nil, err
        }
    
        if resp.StatusCode == 429 || strings.Contains(string(body), "Too many requests") {
            return nil, fmt.Errorf("rate limit exceeded")
        }
    
        // Unmarshal the JSON response
        var response struct {
            Data []struct {
                DestID   string `json:"dest_id"`
                Name     string `json:"name"`
                // Address  string `json:"address"`
                CityName string `json:"city_name"`
                // Add other fields if needed
            } `json:"data"`
        }
    
        err = json.Unmarshal(body, &response)
        if err != nil {
            return nil, err
        }
    
        // Map the API response to your models.Property struct
        properties := make([]models.Property, 0, len(response.Data))
        for _, item := range response.Data {
            properties = append(properties, models.Property{
                DestID:   item.DestID,
                Name:     item.Name,
                // Address:  item.Address,
                CityName: item.CityName,
                // Map other fields as necessary
            })
        }
    
        return properties, nil
    }