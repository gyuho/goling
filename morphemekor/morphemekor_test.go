package morphemekor

import "testing"

func TestSegment(t *testing.T) {
	// map testing data to suggested segmentation
	samples := map[string]string{
		"나는산에가고싶다":    "나는 산에 가고싶다",
		"졸리셨겠어요":      "졸리셨겠어요",
		"나는친구한테듣는다안녕": "나는 친구한테 듣는다 안녕",
		"쓰레기통을비우다":    "쓰레기통을 비우다",
		"모기에물리다":      "모기에 물리다",
		"노인은허리를굽히지만":  "노인은 허리를 굽히지만",
		"비가섞이다":       "비가 섞이다",
		"친구에게선물을주다":   "친구에게 선물을 주다",
		"너무졸리다":       "너무 졸리다",
		"학교에가야하지만":    "학교에 가야하지만",
		"나는바다에가고싶다":   "나는 바다에 가고싶다",
		"과자를먹었다":      "과자를 먹었다",
		"사과와배":        "사과와 배",
		"시장에서오다":      "시장에서 오다",
		"영화를보다":       "영화를 보다",
		"선생님에게편지를쓰다":  "선생님에게 편지를 쓰다",
		"한국에서바다를보다":   "한국에서 바다를 보다",
		"미국으로소포를보내다":  "미국으로 소포를 보내다",
		"미국은전쟁에서이겼지만": "미국은 전쟁에서 이겼지만",
		"아버지께죄송하다":    "아버지께 죄송하다",
		"시간과돈":        "시간과 돈",
		"해변으로여행을가다":   "해변으로 여행을 가다",
		"김치를먹었다":      "김치를 먹었다",
		"싸움을보다":       "싸움을 보다",
		"아기에게점심을먹이다":  "아기에게 점심을 먹이다",
	}
	t.Logf("Ok to not be equal. Experimental project.\n\n")
	correct := 0
	for prev, seg1 := range samples {
		seg2 := Segment(prev)
		if seg1 == seg2 {
			t.Logf("CORRECT | %s --> %s\n", prev, seg2)
			correct++
			continue
		}
		t.Logf("  WRONG | %s --> %s\n", prev, seg2)
	}
	t.Logf("\n\nAccuracy: %f %%", float64(correct)/float64(len(samples))*100)
}

/*
=== RUN TestSegment
--- PASS: TestSegment (0.00s)
	morphemekor_test.go:34: Ok to not be equal. Experimental project.

	morphemekor_test.go:39: CORRECT | 아기에게점심을먹이다 --> 아기에게 점심을 먹이다
	morphemekor_test.go:43:   WRONG | 나는산에가고싶다 --> 나는 산에가 고싶다
	morphemekor_test.go:39: CORRECT | 시장에서오다 --> 시장에서 오다
	morphemekor_test.go:39: CORRECT | 김치를먹었다 --> 김치를 먹었다
	morphemekor_test.go:43:   WRONG | 나는바다에가고싶다 --> 나는 바다에가 고싶다
	morphemekor_test.go:39: CORRECT | 선생님에게편지를쓰다 --> 선생님에게 편지를 쓰다
	morphemekor_test.go:39: CORRECT | 미국은전쟁에서이겼지만 --> 미국은 전쟁에서 이겼지만
	morphemekor_test.go:43:   WRONG | 나는친구한테듣는다안녕 --> 나는 친구한테 듣는 다안녕
	morphemekor_test.go:39: CORRECT | 쓰레기통을비우다 --> 쓰레기통을 비우다
	morphemekor_test.go:43:   WRONG | 너무졸리다 --> 너무졸리다
	morphemekor_test.go:39: CORRECT | 비가섞이다 --> 비가 섞이다
	morphemekor_test.go:39: CORRECT | 친구에게선물을주다 --> 친구에게 선물을 주다
	morphemekor_test.go:43:   WRONG | 학교에가야하지만 --> 학교에가 야하지만
	morphemekor_test.go:39: CORRECT | 영화를보다 --> 영화를 보다
	morphemekor_test.go:39: CORRECT | 한국에서바다를보다 --> 한국에서 바다를 보다
	morphemekor_test.go:39: CORRECT | 졸리셨겠어요 --> 졸리셨겠어요
	morphemekor_test.go:43:   WRONG | 모기에물리다 --> 모기에물리다
	morphemekor_test.go:39: CORRECT | 노인은허리를굽히지만 --> 노인은 허리를 굽히지만
	morphemekor_test.go:39: CORRECT | 미국으로소포를보내다 --> 미국으로 소포를 보내다
	morphemekor_test.go:39: CORRECT | 아버지께죄송하다 --> 아버지께 죄송하다
	morphemekor_test.go:39: CORRECT | 해변으로여행을가다 --> 해변으로 여행을 가다
	morphemekor_test.go:39: CORRECT | 싸움을보다 --> 싸움을 보다
	morphemekor_test.go:39: CORRECT | 과자를먹었다 --> 과자를 먹었다
	morphemekor_test.go:43:   WRONG | 사과와배 --> 사과 와배
	morphemekor_test.go:39: CORRECT | 시간과돈 --> 시간과 돈
	morphemekor_test.go:45:

		Accuracy: 72.000000 %
PASS
ok  	github.com/gyuho/goling/morphemekor	0.002s
*/
