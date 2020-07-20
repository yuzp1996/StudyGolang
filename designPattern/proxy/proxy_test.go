package proxy_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/yuzp1996/StudyGolang/designPattern/proxy"
)

var _ = Describe("Proxy", func() {
	It("simulate proxy", func() {
		server := proxy.NewGitHub("github")
		proxycmd := proxy.NewGitCmd(server,"ready for the clone","gitcmd","after the clone")
		coder := proxy.NewCoder("zpyu",*proxycmd)

		coder.Clone()

	})

})
