package api

import (
	"net/http"

	"bookings.server/pkg/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Dao db.DAO
}

func (_ *Server) ServeRoot(c *gin.Context) {
	c.JSON(http.StatusOK, MessageResponse{
		Message: "Hello World!",
	})
}

func (s *Server) ServeStudents(c *gin.Context) {
	students, err := s.Dao.GetStudents()
	if err != nil {
		errStr := err.Error()
		c.JSON(http.StatusInternalServerError, MessageResponse{
			Error:   &errStr,
			Message: "failed to fetch student list",
		})
		return
	}

	c.JSON(http.StatusOK, students)
}

func (s *Server) ServeStudentById(c *gin.Context, id int64) {

}
