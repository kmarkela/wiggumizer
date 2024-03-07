package scanner

import (
	"github.com/kmarkela/Wiggumizeng/internal/checkers/lfi_checker"
	"github.com/kmarkela/Wiggumizeng/internal/checkers/redirect_checker"
	"github.com/kmarkela/Wiggumizeng/internal/checkers/ssrf_checker"
	"github.com/kmarkela/Wiggumizeng/internal/checkers/xml_checker"
)

func (s *Scanner) registerCheckers() {
	s.checkers["xml"] = xml_checker.New()
	s.checkers["redirect"] = redirect_checker.New()
	s.checkers["ssrf"] = ssrf_checker.New()
	s.checkers["lfi"] = lfi_checker.New()
}
