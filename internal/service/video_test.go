package service

import "testing"

func TestCheckPlayExpired(t *testing.T) {
	i := len("6bB+Kl7vDqV3c+dnn6DEpw==")                                                                 // 失败
	l := len("yOnpT29ibinZi6RRCZG9ohrrItBpGi7wwbdrFoW7mqXjfjSIroEZBr00XymUvxu+uMW58RAshFy99lOrE/m3Qg==") // 成功
	t.Log(i)
	t.Log(l)
	ok, err := checkPlayExpired("MCZY-597ebzD43S_OGrQhru8gOFbzJ3MrxYnRvVJ02hn1vEuIbchIx8HA78gemIQhzu1iULAg404XAhxF1VfUvpngPDZ3hn7LHYSp2ZkYHWiFLK7IgBVDBQ")
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(ok)
}
