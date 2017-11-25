package btc_utils

import (
    "io/ioutil"
    "net/http"
    "encoding/json"
    "errors"
)

/*
curl https://chain.api.btc.com/v3/tx/unconfirmed/summary
{"data":{"size":97722236,"count":73213},"err_no":0,"err_msg":null}
*/
type CheckMessage struct {
    Data UnconfirmedMessage `json:"data"`
    ErrNo int `json:"err_no"`
    ErrMsg * string `json:"err_msg"`
}

type UnconfirmedMessage struct {
    Size int64 `json:"size"`
    Count int `json:"count"`
}

var (
    CHECK_URL = "https://chain.api.btc.com/v3/tx/unconfirmed/summary"
)

func GetUnconfirmedCount() (int, error) {

    url := CHECK_URL
    res, err := http.Get(url)

    if err != nil {
        panic(err.Error())
    }

    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        return 0, err
    }

    var checkMessage CheckMessage
    err = json.Unmarshal(body,&checkMessage)

    if err != nil {
        return 0, err
    }

    if checkMessage.ErrNo != 0 {
        return 0, errors.New(*checkMessage.ErrMsg)
    }


    return checkMessage.Data.Count,nil
}