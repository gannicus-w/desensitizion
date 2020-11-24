package ddz

import (
	"fmt"
	//"strconv"
	//"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToDz(t *testing.T) {
	tests := []struct {
		input  interface{}
		preNum int
		sufNum int
		expect string
	}{
		{12345678, 4, 2, "1234**78"},
		{12345678, 4, 4, "12345678"},
		{12345678, 4, 5, "12345678"},
		{12345678, 9, 5, "12345678"},
		{12345678, 0, 0, "********"},
		{12345678, 0, 3, "*****678"},
		{12345678, 3, 0, "123*****"},
		{12345678, -2, 0, "********"},
		{12345678, 0, -3, "********"},
		{12345678, -2, -3, "********"},
		{123456.78, 4, 2, "1234***78"},
		{-123456.78, 4, 3, "-123***.78"},
		{"12345678", 4, 2, "1234**78"},
		{[]byte("123 45678"), 4, 2, "123 ***78"},
		{true, 0, 0, "****"},
		{false, 0, 0, "*****"},
		{nil, 0, 5, ""},
	}

	for _, test := range tests {
		errmsg := fmt.Sprintf("test = %v", test)
		v := ToDz(test.input, test.preNum, test.sufNum)
		assert.Equal(t, test.expect, v, errmsg)
	}
}

func TestDesensitizeString(t *testing.T) {
	type args struct {
		cipherText string
		mask       string
		expect string
	}

	str := ";"
	fmt.Println(str)
	str2 := []rune(str)
	fmt.Println(str2)

/*	str := "abc,bcd,bbb,sefsd"
	fmt.Println(str)

	str2 := strings.Split(str, ",")

	fmt.Println(str2)

	str1 := "+1"
	//fmt.Println(strconv.ParseUint(str1,10,0))
	fmt.Println(strconv.Atoi(str1))
	// direction int, cipherText, mask string*/
	tests := []struct {
		cipherText string
		mask string
		expect string
	}{
		//{"411324199010162819", "3,2;6,4;13,3","411**4****101***19"},
		//{"13838985804", "3,2;6,1","138**9*5804"},
		//{"13838985804", "3,20","138********"},
		{"13838985804", "3","13**8985804"},
		//{"+13838985804", "+3,2","+13**8985804"},
		//{"13838985804", "3,2;","138**985804"},
		//{1, "13838985804", "0x78","****8985804"},
		//{0, "13838985804", "0x078","1383****804"},
		//{1, "13838985804", "0x078","****8985804"},
		//{0, "-12345678", "0xF","-1234****"},
		//{1, "-12345678", "0xF","****45678"},
		//{0, "411324199010162819", "0x0ff0","411324********2819"},
		//{1, "411324199010162819", "0x0ff0","********9010162819"},
		//{1, "411324199010162819", "0x0aaa","*1*3*4*9*0*0162819"},
		//{0, "411324199010162819", "0x0aaa","411324*9*0*0*6*8*9"},
		//{0, "巴拉巴拉小魔仙", "0x7","巴拉巴拉***"},
		//{1, "巴拉巴拉小魔仙", "0x7","***拉小魔仙"},
		//{0, "浙江省杭州市西湖区龙湖大道七号街4-20-88", "0xFF","浙江省杭州市西湖区龙湖大道七号********"},
	}
	for _, test := range tests {
		errmsg := fmt.Sprintf("test = %v", test)
		v := DesensitizeString(test.cipherText, test.mask)
		assert.Equal(t, test.expect, v, errmsg)
	}
}
