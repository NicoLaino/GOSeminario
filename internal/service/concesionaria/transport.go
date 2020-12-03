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
		path:     "/messages",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/messages/:id",
		function: getMessageByID(s),
	})

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/messages/:id",
		function: deleteMessageByID(s),
	})
	
	list = append(list, &endpoint{
		method:   "POST",
		path:     "/messages",
		function: postMessage(s),
	})
/*
	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/messages/:id",
		function: updateMessage(s),
	})*/

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, errFindAll := s.FindAll()
		if errFindAll != nil {
			fmt.Println(errFindAll.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"messages": result,
		})
	}
}

func getMessageByID(s Service) gin.HandlerFunc {
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
			"messages": *result,
		})
	}
}

func deleteMessageByID(s Service) gin.HandlerFunc {
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
			"messages": result,
		})
	}
}

func postMessage(s Service) gin.HandlerFunc {

	return func(c *gin.Context) {
		var message Message

		c.BindJSON(&message)
		queryResult, err := s.AddMessage(message)

		if err != nil {
			fmt.Println(err.Error())
		}
			c.JSON(http.StatusOK, gin.H{
				"messages": queryResult,
			})
	}

}




/*func updateBook(s BookService) gin.HandlerFunc {

	return func(c *gin.Context) {
		var book Book
		var httpErrorMsg *ErrorResponse

		ID, errAtoi := strconv.Atoi(c.Param("id"))

		c.BindJSON(&book)
		queryResult, err := s.UpdateBook(ID, book)

		if errAtoi != nil {
			httpErrorMsg = &ErrorResponse{Message: errAtoi.Error()}
		}
		if err != nil {
			httpErrorMsg = &ErrorResponse{Message: err.Error()}
		}
		if queryResult.rowsAffected == 0 {
			httpErrorMsg = &ErrorResponse{Message: fmt.Sprintf("Requested ID: %v not found", ID)}
		}

		if errAtoi != nil || err != nil || queryResult.rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": httpErrorMsg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"Updated books": queryResult.rowsAffected,
			})
		}

	}
}*/

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}