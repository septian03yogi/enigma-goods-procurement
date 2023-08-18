package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/septian03yogi/config"
	"github.com/septian03yogi/delivery/controller"
	"github.com/septian03yogi/repository"
	"github.com/septian03yogi/usecase"
	"github.com/septian03yogi/utils/exception"
)

type Server struct {
	departmentUC usecase.DepartmentUseCase
	engine       *gin.Engine
	host         string
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) initController() {
	controller.NewDepartmentController(s.departmentUC, s.engine)
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exception.CheckErr(err)
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	departmentRepo := repository.NewDepartmentRepository(db)
	departmentUseCase := usecase.NewDepartmentUseCase(departmentRepo)
	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		departmentUC: departmentUseCase,
		engine:       engine,
		host:         host,
	}
}
