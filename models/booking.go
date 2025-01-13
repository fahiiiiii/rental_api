// models/booking.go
package models

import (
    "github.com/beego/beego/v2/client/orm"
)

// City model represents a city from the Booking.com API
type City struct {
    Id        int64   `orm:"auto;pk" json:"id"`
    CC1       string  `orm:"size(2)" json:"cc1"`
    ImageURL  string  `orm:"size(255)" json:"image_url"`
    Longitude float64 `orm:"digits(10);decimals(7)" json:"longitude"`
    CityName  string  `orm:"size(100)" json:"city_name"`
    DestID    string  `orm:"size(50)" json:"dest_id"`
    Timezone  string  `orm:"size(50)" json:"timezone"`
    Hotels    int     `orm:"default(0)" json:"hotels"`
    Label     string  `orm:"size(100)" json:"label"`
    Country   string  `orm:"size(100)" json:"country"`
    Region    string  `orm:"size(100)" json:"region"`
    DestType  string  `orm:"size(50)" json:"dest_type"`
    Name      string  `orm:"size(100)" json:"name"`
    Latitude  float64 `orm:"digits(10);decimals(7)" json:"latitude"`
    Type      string  `orm:"size(50)" json:"type"`
    Created   time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
    Updated   time.Time `orm:"auto_now;type(datetime)" json:"updated"`
}

// Property model represents a property from the Booking.com API
type Property struct {
    Id              int64   `orm:"auto;pk" json:"id"`
    UFI             int64   `orm:"unique" json:"ufi"`
    CheckoutDate    string  `orm:"size(20)" json:"checkout_date"`
    ReviewScoreWord string  `orm:"size(50)" json:"review_score_word"`
    Longitude       float64 `orm:"digits(10);decimals(7)" json:"longitude"`
    IsPreferred     bool    `orm:"default(false)" json:"is_preferred"`
    CountryCode     string  `orm:"size(2)" json:"country_code"`
    Latitude        float64 `orm:"digits(10);decimals(7)" json:"latitude"`
    WishlistName    string  `orm:"size(100)" json:"wishlist_name"`
    Name            string  `orm:"size(255)" json:"name"`
    PropertyClass   float64 `orm:"digits(3);decimals(1)" json:"property_class"`
    DestID          string  `orm:"size(50)" json:"dest_id"`
    CityName        string  `orm:"size(100)" json:"city_name"`
    Country         string  `orm:"size(100)" json:"country"`
    Created         time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
    Updated         time.Time `orm:"auto_now;type(datetime)" json:"updated"`
}

// CityKey model represents a unique identifier for a city
type CityKey struct {
    Id      int64  `orm:"auto;pk" json:"id"`
    Name    string `orm:"size(100)" json:"name"`
    Country string `orm:"size(100)" json:"country"`
    Created time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
    Updated time.Time `orm:"auto_now;type(datetime)" json:"updated"`
}

// HotelDetails model represents the full hotel information
type HotelDetails struct {
    Id           int64      `orm:"auto;pk" json:"id"`
    HotelID      string     `orm:"size(50);unique" json:"hotel_id"`
    PropertyName string     `orm:"size(255)" json:"property_name"`
    Type         string     `orm:"size(50)" json:"type"`
    Bedrooms     int        `orm:"default(0)" json:"bedrooms"`
    Bathroom     int        `orm:"default(0)" json:"bathroom"`
    Amenities    []*Facility `orm:"rel(m2m)" json:"amenities"`
    Description  string     `orm:"type(text)" json:"description"`
    Created      time.Time  `orm:"auto_now_add;type(datetime)" json:"created"`
    Updated      time.Time  `orm:"auto_now;type(datetime)" json:"updated"`
}

// Facility model represents a hotel facility/amenity
type Facility struct {
    Id      int64  `orm:"auto;pk" json:"id"`
    Name    string `orm:"size(100)" json:"name"`
    Created time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
    Updated time.Time `orm:"auto_now;type(datetime)" json:"updated"`
}

// CategorizedImages model represents different categories of property images
type CategorizedImages struct {
    Id               int64    `orm:"auto;pk" json:"id"`
    PropertyBuilding []string `orm:"type(json)" json:"property_building"`
    Property         []string `orm:"type(json)" json:"property"`
    Room             []string `orm:"type(json)" json:"room"`
    Created          time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
    Updated          time.Time `orm:"auto_now;type(datetime)" json:"updated"`
}

// RatingAndReviews model represents property ratings and reviews
type RatingAndReviews struct {
    Id                      int64   `orm:"auto;pk" json:"id"`
    AverageScoreOutOf10    float64 `orm:"digits(4);decimals(2)" json:"average_score_out_of_10"`
    VpmFavorableReviewCount int     `orm:"default(0)" json:"vpm_favorable_review_count"`
    Created                 time.Time `orm:"auto_now_add;type(datetime)" json:"created"`
    Updated                 time.Time `orm:"auto_now;type(datetime)" json:"updated"`
}

func init() {
    // Register models to orm
    orm.RegisterModel(new(City))
    orm.RegisterModel(new(Property))
    orm.RegisterModel(new(CityKey))
    orm.RegisterModel(new(HotelDetails))
    orm.RegisterModel(new(Facility))
    orm.RegisterModel(new(CategorizedImages))
    orm.RegisterModel(new(RatingAndReviews))
}