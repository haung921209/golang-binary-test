//go:build pg14
// +build pg14

package service


func (s *PostgresService) FunctionSpecificS() string {
	return "hello from postgresql14"
}

func (s *PostgresService) FunctionSpecific(args []string) {

}

type PostGre16Service struct {
	PostgresService
}

func (s *PostGre16Service) FunctionSpecificS() string {
	return "hellllllllo from postgresql16"
}

func (s *PostGre16Service) FunctionSpecific(args []string) {

}
