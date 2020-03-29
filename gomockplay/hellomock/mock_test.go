package gomocktest

import (
	"StudyGolang/gomockplay/hellomock/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestCompany_Meeting(t *testing.T) {
	ctl := gomock.NewController(t)

	//返回了一个可以mock的接口
	mock_talker := mock_gomocktest.NewMockTakler(ctl)

	//在这里规定了他的期望返回值
	mock_talker.EXPECT().SayHello(gomock.Eq("王尼玛")).Return("你成功了")

	//这个是个接口 其实相当于是个参数 要带入其他的函数
	company := NewCompany(mock_talker)
	t.Log(company.Meeting("王尼玛"))
}

//操作步骤  参考：https://segmentfault.com/a/1190000009894570
