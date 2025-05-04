package config

import (
	"fmt"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	defaultPage  = 1
	defaultLimit = 10
)

type Pagination struct {
	Page  int
	Limit int
}

type PaginationResponse struct {
	CurrentPage     int `json:"current_page"`
	PageCount       int `json:"page_count"`
	TotalPagesCount int `json:"total_pages_count"`
}

type DateFilterQuery struct {
	FromDate string `form:"from"`
	ToDate   string `form:"to"`
}

func GetPagination(c *gin.Context) Pagination {
	var (
		page  *int
		limit *int
	)
	if c.Query("page") != "" {
		pageInt, err := strconv.Atoi(c.Query("page"))
		if err == nil {
			page = &pageInt
		}
	}
	if c.Query("limit") != "" {
		limitInt, err := strconv.Atoi(c.Query("limit"))
		if err == nil {
			limit = &limitInt
		}
	}

	if page != nil && limit != nil {
		return Pagination{Page: *page, Limit: *limit}
	} else if page == nil && limit != nil {
		return Pagination{Page: defaultPage, Limit: *limit}
	} else if page != nil && limit == nil {
		return Pagination{Page: *page, Limit: defaultLimit}
	} else {
		return Pagination{Page: defaultPage, Limit: defaultLimit}
	}
}

func GetDateFilterQuery(c *gin.Context) (time.Time, time.Time) {
	const layout = "02-01-2006"
	fromStr := c.Query("from")
	toStr := c.Query("to")

	now := time.Now()
	defaultFrom := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	defaultTo := now

	fromDate := defaultFrom
	toDate := defaultTo

	if fromStr != "" {
		if parsed, err := time.Parse(layout, fromStr); err == nil {
			fromDate = parsed
		}
	}

	if toStr != "" {
		if parsed, err := time.Parse(layout, toStr); err == nil {
			toDate = parsed
		}
	}

	return fromDate, toDate
}


func FindOneByField(db *gorm.DB, model interface{}, field string, value interface{}) error {
	result := db.Where(fmt.Sprintf("%s = ?", field), value).First(model)
	return result.Error
}

func FindByTwoFields(db *gorm.DB, model interface{}, field1 string, value1 interface{}, field2 string, value2 interface{}) error {
	result := db.Where(fmt.Sprintf("%s AND %s", field1, field2), value1, value2).First(model)
	return result.Error
}

func FindByThreeFields(db *gorm.DB, model interface{}, field1 string, value1 interface{}, field2 string, value2 interface{}, field3 string, value3 interface{}) error {
	result := db.Where(fmt.Sprintf("%s AND %s AND %s", field1, field2, field3), value1, value2, value3).First(model)
	return result.Error
}

func FindByID(db *gorm.DB, model interface{}, id interface{}) error {
	result := db.Where("id = ?", id).First(model)
	return result.Error
}

func FindByFieldPaginated(db *gorm.DB, model interface{}, field string, value interface{}, pagination Pagination) error {
	offset := (pagination.Page - 1) * pagination.Limit

	result := db.Where(fmt.Sprintf("%s = ?", field), value).
		Limit(pagination.Limit).
		Offset(offset).
		Find(model)

	return result.Error
}

func FindByTwoFieldsPaginated(db *gorm.DB, model interface{}, field1 string, value1 interface{}, field2 string, value2 interface{}, pagination Pagination) error {
	offset := (pagination.Page - 1) * pagination.Limit

	result := db.Where(fmt.Sprintf("%s = ? AND %s = ?", field1, field2), value1, value2).
		Limit(pagination.Limit).
		Offset(offset).
		Find(model)

	return result.Error
}

func FindByThreeFieldsPaginated(db *gorm.DB, model interface{}, field1 string, value1 interface{}, field2 string, value2 interface{}, field3 string, value3 interface{}, pagination Pagination) error {
	offset := (pagination.Page - 1) * pagination.Limit

	result := db.Where(fmt.Sprintf("%s AND %s AND %s", field1, field2, field3), value1, value2, value3).
		Limit(pagination.Limit).
		Offset(offset).
		Find(model)

	return result.Error
}

func FindByUserAndDateRangePaginated(db *gorm.DB, model interface{}, userID string, fromDate, toDate time.Time, pagination Pagination) error {
	offset := (pagination.Page - 1) * pagination.Limit

	result := db.Where("user_id = ? AND created_at BETWEEN ? AND ?", userID, fromDate, toDate).
		Order("created_at DESC").
		Limit(pagination.Limit).
		Offset(offset).
		Find(model)

	return result.Error
}
