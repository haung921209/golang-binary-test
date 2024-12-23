package service

type IPostgresService interface {
	FunctionCommonS(args []string) string
	FunctionCommon(args []string)
	FunctionSpecificS() string
	FunctionSpecific(args []string)
}

type PostgresService struct {
	IPostgresService
}

func (fs *PostgresService) FunctionCommonS(args []string) string {
	return "hello"
}

func (fs *PostgresService) FunctionCommon(args []string) {

}
