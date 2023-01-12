package utils

import (
	"fmt"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PageInfoTest struct {
	PageInfo request.PageInfo
	Name     string
}

func TestVerify(t *testing.T) {
	PageInfoVerify := Rules{"Page": {NotEmpty()}, "PageSize": {Gt("0"), Lt("100")}, "Name": {RegexpMatch(`[12]{1}`)}}
	var testInfo PageInfoTest
	testInfo.Name = "12"
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 99
	err := Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能捕捉0值")
	}
	fmt.Println(err)
	testInfo.Name = ""
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 101
	err = Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("校验失败，未能正常检测name为空")
	}
	fmt.Println(err)
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 10
	err = Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("校验失败，未能正常通过检测")
	}
	fmt.Println(err)
}
