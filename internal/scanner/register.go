package scanner

import (
	"github.com/kmarkela/wiggumizer/internal/checkers/lfi_checker"
	"github.com/kmarkela/wiggumizer/internal/checkers/redirect_checker"
	secretchecker "github.com/kmarkela/wiggumizer/internal/checkers/secret_checker"
	"github.com/kmarkela/wiggumizer/internal/checkers/ssrf_checker"
	subdchecker "github.com/kmarkela/wiggumizer/internal/checkers/subd_checker"
	"github.com/kmarkela/wiggumizer/internal/checkers/xml_checker"
)

func (s *Scanner) registerCheckers() {
	s.checkers["xml"] = xml_checker.New()
	s.checkers["redirect"] = redirect_checker.New()
	s.checkers["ssrf"] = ssrf_checker.New()
	s.checkers["subd"] = subdchecker.New()
	s.checkers["secret"] = secretchecker.New()
	s.checkers["lfi"] = lfi_checker.New()
}
