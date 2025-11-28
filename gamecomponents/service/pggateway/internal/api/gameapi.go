package api

import (
	"encoding/json"
	"fmt"
	"game/comm/define"
	"game/comm/mq"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PGParams = define.PGParams

//	{
//	    "dt": null,
//	    "err": {
//	        "cd": "1200",
//	        "msg": "OERR: Operator return an error. Failed to verify operator player session",
//	        "tid": "HXTZJF11"
//	    }
//	}

type PGError struct {
	Cd  string `json:"cd"`
	Msg string `json:"msg"`
	Tid string `json:"tid"`
}

func (e *PGError) Error() string {
	return e.Msg
}

type PGRetWrapper struct {
	Dt  *json.RawMessage `json:"dt"`
	Err *PGError         `json:"err"`
}

func checkIp(ip string) (allow bool, err error) {
	var netAllow struct {
		Allow bool
	}
	err = mq.Invoke("/AdminInfo/Interior/whiteList_exists", map[string]any{
		"IP": ip,
	}, &netAllow)
	if err != nil {
		return
	}
	allow = netAllow.Allow
	return
}

func gameapi(w http.ResponseWriter, r *http.Request) {
	var (
		ret PGRetWrapper
		err error
	)

	traceId := r.URL.Query().Get("traceId")
	defer func() {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			ret.Err = &PGError{
				Cd:  "1200",
				Msg: err.Error(),
				Tid: traceId,
			}
			if ec, ok := err.(define.IErrcode); ok {
				code := ec.Code()
				if code != 0 {
					ret.Err.Cd = strconv.Itoa(code)
				}
			}
		}
		jsondata, _ := json.Marshal(ret)
		w.Write(jsondata)
	}()

	/*
		err = r.ParseForm()
		if err != nil {
			return
		}

		tk := r.FormValue("atk")
		pid, err := jwtutil.ParseToken(tk)
		if err != nil {
			err = define.NewErrCode("Invalid player session", 1302)
			return
		}

		_, ip, loc := gamedata.IsBlockLoc(r)
		fmt.Printf("userid=%d,ip=%s,loc=%s\n", pid, ip, loc)
		netAllow, err := checkIp(ip)
		if err != nil {
			err = define.NewErrCode("check ip error", 1302)
			return
		}
		if !netAllow {
			err = define.NewErrCode("Your ip exinclude in net whitelist", 1302)
			return
		}

		ps := &PGParams{
			Path:    r.URL.Path,
			TraceId: traceId,
			Form:    r.Form,
			Pid:     pid,
		}

		var resp json.RawMessage
		err = mq.Invoke(r.URL.Path, ps, &resp)
		// if err != nil {
		// 	return
		// }

		if len(resp) != 0 {
			ret.Dt = &resp
		}

		if err == nil {
			onSpin(ps, resp)
		}
	*/

	var data M
	if strings.Contains(r.URL.Path, "/Spin") {
		spinId := fmt.Sprintf("%d%09d", time.Now().UnixNano()/1000000, rand.Intn(1000000000))

		data = M{
			"si": M{
				"wp":   nil,
				"lw":   nil,
				"irfs": false,
				"itr":  false,
				"itff": false,
				"it":   false,
				"ifsw": false,
				"gm":   0,
				"crtw": 0,
				"imw":  false,
				"orl": []any{
					6,
					7,
					6,
					99,
					6,
					4,
					6,
					7,
					7,
					4,
					7,
					99,
				},
				"rcs": 0,
				"gwt": -1,
				"pmt": nil,
				"ab":  nil,
				"ml":  2,
				"cs":  0.3,
				"rl": []any{
					6,
					7,
					6,
					99,
					6,
					4,
					6,
					7,
					7,
					4,
					7,
					99,
				},
				"ctw":   0,
				"cwc":   0,
				"fstc":  nil,
				"pcwc":  0,
				"rwsp":  nil,
				"hashr": "0:6;6;7#7;4;4#6;6;7#99;7;99#MV#6.0#MT#0#MG#0#",
				"fb":    nil,
				"sid":   spinId,
				"psid":  spinId,
				"st":    1,
				"nst":   1,
				"pf":    1,
				"aw":    0,
				"wid":   0,
				"wt":    "C",
				"wk":    "0_C",
				"wbn":   nil,
				"wfg":   nil,
				"blb":   100000,
				"blab":  99994,
				"bl":    99994,
				"tb":    6,
				"tbb":   6,
				"tw":    0,
				"np":    -6,
				"ocr":   nil,
				"mr":    nil,
				"ge": []any{
					1,
					11,
				},
			},
		}
	} else {
		data = M{
			"cs":    []any{0.03, 0.1, 0.3, 0.9},
			"ml":    []any{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			"mxl":   10,
			"maxwm": 5000,
			"wt": M{
				"mw":   3,
				"bw":   5,
				"mgw":  15,
				"smgw": 35,
			},
			"fb": M{
				"is":  true,
				"bm":  100,
				"bms": nil,
				"t":   0.3,
			},
			"gcs": M{
				"ab": M{
					"is":  false,
					"bm":  0,
					"bms": []any{1.5},
					"t":   0,
				},
			},
			"bl":   100000,
			"inwe": false,
			"iuwe": false,
			"ls": M{
				"si": M{
					"wp":    nil,
					"lw":    nil,
					"irfs":  false,
					"itr":   false,
					"itff":  false,
					"it":    false,
					"ifsw":  true,
					"gm":    0,
					"crtw":  0,
					"imw":   false,
					"orl":   []any{2, 2, 2, 99, 0, 0, 0, 0, 3, 3, 3, 99},
					"rcs":   0,
					"gwt":   0,
					"pmt":   nil,
					"ab":    nil,
					"ml":    2,
					"cs":    0.3,
					"rl":    []any{2, 2, 2, 99, 0, 0, 0, 0, 3, 3, 3, 99},
					"ctw":   0,
					"cwc":   0,
					"fstc":  nil,
					"pcwc":  0,
					"rwsp":  nil,
					"hashr": nil,
					"fb":    nil,
					"sid":   "0",
					"psid":  "0",
					"st":    1,
					"nst":   1,
					"pf":    0,
					"aw":    0,
					"wid":   0,
					"wt":    "C",
					"wk":    "0_C",
					"wbn":   nil,
					"wfg":   nil,
					"blb":   0,
					"blab":  0,
					"bl":    100000,
					"tb":    0,
					"tbb":   0,
					"tw":    0,
					"np":    0,
					"ocr":   nil,
					"mr":    nil,
					"ge":    nil,
				},
			},
			"cc": "PGC",
		}
	}

	jsonData, _ := json.Marshal(data)
	rawMsg := json.RawMessage(jsonData)
	ret.Dt = &rawMsg
}
