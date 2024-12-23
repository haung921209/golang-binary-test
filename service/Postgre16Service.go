//go:build pg16
// +build pg16

package service

func (s *PostgresService) FunctionSpecificS() string {
	return "hellllllllo from postgresql16"
}

func (s *PostgresService) FunctionSpecific(args []string) {

}
