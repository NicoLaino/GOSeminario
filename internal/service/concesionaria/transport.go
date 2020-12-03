package concesionaria

import (
	"net/http"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

// NewHTTPTransport ...
func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/cars",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/cars/:id",
		function: getCarByID(s),
	})

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/cars/:id",
		function: deleteCarByID(s),
	})
	
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/cars",
		function: postCar(s),
	})
	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/cars/:id",
		function: updateCar(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, errFindAll := s.FindAll()
		if errFindAll != nil {
			fmt.Println(errFindAll.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"Cars": result,
		})
	}
}

func getCarByID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(err.Error())
		}
		result, errFindByID := s.FindByID(ID)
		if errFindByID != nil {
			fmt.Println(errFindByID.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"Car": *result,
		})
	}
}

func deleteCarByID(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(err.Error())
		}
		result, errDeleteByID := s.DeleteByID(ID)
		if errDeleteByID != nil {
			fmt.Println(errDeleteByID.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"Result": result,
		})
	}
}

func postCar(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var car Car

		c.BindJSON(&car)
		queryResult, err := s.AddCar(car)

		if err != nil {
			fmt.Println(err.Error())
		}
			c.JSON(http.StatusOK, gin.H{
				"Result": queryResult,
			})
	}
}

func updateCar(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var car Car
		ID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(err.Error())
		}
		c.BindJSON(&car)
		result, errUpdateByID := s.UpdateByID(ID, car)
		if errUpdateByID != nil {
			fmt.Println(errUpdateByID.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"Result": result,
		})
	}
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}