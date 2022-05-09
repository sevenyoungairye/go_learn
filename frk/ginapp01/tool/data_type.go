package tool

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	YyyyMmDdHhMmSs = "2006-01-02 15:04:05"
	YyyyMmDd       = "2006-01-02"
	YyyyMmDdHh     = "2006-01-02 15"
	YyyyMmDdHhMm   = "2006-01-02 15:04"
)

//Date 自定义时间
// https://segmentfault.com/a/1190000022264001
// 进行MySQL日期的绑定与读取
type Date time.Time

//UnmarshalJSON 接收数据并转换成日期
func (t *Date) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(YyyyMmDdHhMmSs, timeStr)
	if err != nil {
		t1, err = time.Parse(YyyyMmDd, timeStr)
	}
	if err != nil {
		t1, err = time.Parse(YyyyMmDdHhMm, timeStr)
	}
	if err != nil {
		t1, err = time.Parse(YyyyMmDdHh, timeStr)
	}
	*t = Date(t1)
	return err
}

// MarshalJSON 响应json格式
func (t Date) MarshalJSON() ([]byte, error) {
	timeStr := t.String()
	fmt.Println("日期...", timeStr)
	if strings.Contains(timeStr, "0001-01-01") {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format(t.GetFmt()))
	return []byte(formatted), nil
}

func (t Date) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format(t.GetFmt()), nil
}

func (t *Date) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = Date(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *Date) String() string {
	return fmt.Sprintf("get time:%s", time.Time(*t).String())
}

func (t *Date) GetFmt() string {
	timeStr := time.Time(*t).String()
	splitArr := strings.Split(timeStr, " ")
	if splitArr != nil && len(splitArr) == 4 {
		hms := splitArr[1]
		hmsArr := strings.Split(hms, ":")
		if hms == "00:00:00" {
			return YyyyMmDd
		}
		if hmsArr[1] == "00" && hmsArr[2] == "00" {
			return YyyyMmDdHh
		}
		// yyyyMMdd HH:mm:ss
		// 2022-01-01 12:19 会被默认格式化 2022-01-01 12:19:00
		/*if hmsArr[2] == "00" {
			return YyyyMmDdHhMm
		}*/
	}

	return YyyyMmDdHhMmSs
}

/*func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}*/

// NullString struct
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// NullTime sql.NullTime
// https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267
type NullTime struct {
	sql.NullTime
}

func (nt *NullTime) MarshalJSON() ([]byte, error) {

	if nt.Valid {
		t := nt.Time
		var strFmt string

		var err error
		defer func() {
			r := recover()
			if r != nil {
				errMsg := fmt.Sprintln(r)
				fmt.Println("=========error========", errMsg)
				err = errors.New(errMsg)
			}
		}()

		/*return func() ([]byte, error) {

			return []byte("null"), nil
		}()*/

		timeStr := t.String()
		split := strings.Split(timeStr, " ")
		fmt.Println("string time... ", timeStr)
		hmsArr := strings.Split(split[1], ":")
		if split[1] == "00:00:00" {
			strFmt = YyyyMmDd
		} else if hmsArr[1] == "00" && hmsArr[2] == "00" {
			strFmt = YyyyMmDdHh
		} else {
			strFmt = YyyyMmDdHhMmSs
			// panic("throw a error")
		}

		// fixme: 异常信息未捕捉到...
		// 这一步已经走到了
		if err != nil {
			fmt.Println("捕捉到了异常信息...", err)
			return []byte("null"), nil
		}

		return json.Marshal(t.Format(strFmt))
	}
	return []byte("null"), nil
}

/*func (nt *NullTime) Scan(value interface{}) error {
	var i sql.NullTime
	if err := i.Scan(value); err != nil {
		return err
	}
	// if nil the make Valid false
	if reflect.TypeOf(value) == nil {
		//*nt = NullTime{Time: i.Time}

	} else {
		//*nt = NullTime{Time: i.Time, Valid: true}
	}
	return nil
}
*/

/**
var _ json.Unmarshaler = &MyTime{}

func (mt *MyTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	//t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	t, err := time.Parse("2006-01-02", "s")
	if err != nil {
		return err
	}
	*mt = MyTime(t)
	return nil
}
*/
