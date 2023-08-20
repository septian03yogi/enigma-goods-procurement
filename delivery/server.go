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
	employeeUC   usecase.EmployeeUseCase
	itemUC       usecase.ItemUseCase
	periodUC     usecase.PeriodUseCase
	roleUC       usecase.RoleUserUseCase
	submissionUC usecase.SubmissionUseCase
	subStatusUC  usecase.SubmisisonStatusUseCase
	uomUC        usecase.UomUseCase
	userUC       usecase.UserUseCase
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
	controller.NewEmployeeController(s.employeeUC, s.engine)
	controller.NewItemController(s.itemUC, s.engine)
	controller.NewPeriodController(s.periodUC, s.engine)
	controller.NewRoleController(s.roleUC, s.engine)
	controller.NewSubmissionController(s.submissionUC, s.engine)
	controller.NewSubStatusController(s.subStatusUC, s.engine)
	controller.NewUomUseController(s.uomUC, s.engine)
	controller.NewUserController(s.userUC, s.engine)
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	exception.CheckErr(err)
	dbConn, _ := config.NewDbConnection(cfg)
	db := dbConn.Conn()
	departmentRepo := repository.NewDepartmentRepository(db)
	employeeRepo := repository.NewEmployeeRepository(db)
	itemRepo := repository.NewItemRepository(db)
	periodRepo := repository.NewPeriodRepository(db)
	roleRepo := repository.NewRoleUserRepository(db)
	submissionRepo := repository.NewSubmissionRepository(db)
	subStatusRepo := repository.NewSubmissionStatusRepository(db)
	uomRepo := repository.NewUomRepository(db)
	userRepo := repository.NewUserRepository(db)
	departmentUseCase := usecase.NewDepartmentUseCase(departmentRepo)
	employeeUseCase := usecase.NewEmployeeUseCase(employeeRepo)
	itemUseCase := usecase.NewItemUseCase(itemRepo)
	periodUseCase := usecase.NewPeriodUseCase(periodRepo)
	roleUseCase := usecase.NewRoleUserUseCase(roleRepo)
	submissionUseCase := usecase.NewSubmissionUseCase(submissionRepo)
	subStatusUseCase := usecase.NewSubmissionStatusUseCase(subStatusRepo)
	uomUseCase := usecase.NewUomUseCase(uomRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)
	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		departmentUC: departmentUseCase,
		employeeUC:   employeeUseCase,
		itemUC:       itemUseCase,
		periodUC:     periodUseCase,
		roleUC:       roleUseCase,
		submissionUC: submissionUseCase,
		subStatusUC:  subStatusUseCase,
		uomUC:        uomUseCase,
		userUC:       userUseCase,
		engine:       engine,
		host:         host,
	}
}
