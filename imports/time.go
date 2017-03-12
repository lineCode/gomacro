// this file was generated by gomacro command: import "time"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"time"
)

func init() {
	Binds["time"] = map[string]Value{
		"ANSIC":	ValueOf(time.ANSIC),
		"After":	ValueOf(time.After),
		"AfterFunc":	ValueOf(time.AfterFunc),
		"April":	ValueOf(time.April),
		"August":	ValueOf(time.August),
		"Date":	ValueOf(time.Date),
		"December":	ValueOf(time.December),
		"February":	ValueOf(time.February),
		"FixedZone":	ValueOf(time.FixedZone),
		"Friday":	ValueOf(time.Friday),
		"Hour":	ValueOf(time.Hour),
		"January":	ValueOf(time.January),
		"July":	ValueOf(time.July),
		"June":	ValueOf(time.June),
		"Kitchen":	ValueOf(time.Kitchen),
		"LoadLocation":	ValueOf(time.LoadLocation),
		"Local":	ValueOf(&time.Local).Elem(),
		"March":	ValueOf(time.March),
		"May":	ValueOf(time.May),
		"Microsecond":	ValueOf(time.Microsecond),
		"Millisecond":	ValueOf(time.Millisecond),
		"Minute":	ValueOf(time.Minute),
		"Monday":	ValueOf(time.Monday),
		"Nanosecond":	ValueOf(time.Nanosecond),
		"NewTicker":	ValueOf(time.NewTicker),
		"NewTimer":	ValueOf(time.NewTimer),
		"November":	ValueOf(time.November),
		"Now":	ValueOf(time.Now),
		"October":	ValueOf(time.October),
		"Parse":	ValueOf(time.Parse),
		"ParseDuration":	ValueOf(time.ParseDuration),
		"ParseInLocation":	ValueOf(time.ParseInLocation),
		"RFC1123":	ValueOf(time.RFC1123),
		"RFC1123Z":	ValueOf(time.RFC1123Z),
		"RFC3339":	ValueOf(time.RFC3339),
		"RFC3339Nano":	ValueOf(time.RFC3339Nano),
		"RFC822":	ValueOf(time.RFC822),
		"RFC822Z":	ValueOf(time.RFC822Z),
		"RFC850":	ValueOf(time.RFC850),
		"RubyDate":	ValueOf(time.RubyDate),
		"Saturday":	ValueOf(time.Saturday),
		"Second":	ValueOf(time.Second),
		"September":	ValueOf(time.September),
		"Since":	ValueOf(time.Since),
		"Sleep":	ValueOf(time.Sleep),
		"Stamp":	ValueOf(time.Stamp),
		"StampMicro":	ValueOf(time.StampMicro),
		"StampMilli":	ValueOf(time.StampMilli),
		"StampNano":	ValueOf(time.StampNano),
		"Sunday":	ValueOf(time.Sunday),
		"Thursday":	ValueOf(time.Thursday),
		"Tick":	ValueOf(time.Tick),
		"Tuesday":	ValueOf(time.Tuesday),
		"UTC":	ValueOf(&time.UTC).Elem(),
		"Unix":	ValueOf(time.Unix),
		"UnixDate":	ValueOf(time.UnixDate),
		"Until":	ValueOf(time.Until),
		"Wednesday":	ValueOf(time.Wednesday),
	}
	Types["time"] = map[string]Type{
		"Duration":	TypeOf((*time.Duration)(nil)).Elem(),
		"Location":	TypeOf((*time.Location)(nil)).Elem(),
		"Month":	TypeOf((*time.Month)(nil)).Elem(),
		"ParseError":	TypeOf((*time.ParseError)(nil)).Elem(),
		"Ticker":	TypeOf((*time.Ticker)(nil)).Elem(),
		"Time":	TypeOf((*time.Time)(nil)).Elem(),
		"Timer":	TypeOf((*time.Timer)(nil)).Elem(),
		"Weekday":	TypeOf((*time.Weekday)(nil)).Elem(),
	}
	Proxies["time"] = map[string]Type{
	}
}