package services

import (
	"crypto"
	"encoding/hex"
	"errors"
	_ "golang.org/x/crypto/blake2b"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type JobService struct {
	currentJob job
}

func CreateJobService() *JobService {
	js := new(JobService)
	js.currentJob = job{
		last:       "0x0000000000",
		number:     1,
		created:    time.Now().Unix(),
		difficulty: 100,
	}
	message := js.currentJob.last + strconv.FormatInt(js.currentJob.number, 10) + strconv.FormatInt(js.currentJob.created, 10)
	js.currentJob.hash = getHexedHash(message)
	return js
}

func (js *JobService) GetCurrentJob() Job {
	return Job{
		Created:    js.currentJob.created,
		Last:       js.currentJob.last,
		Hash:       js.currentJob.hash,
		Difficulty: js.currentJob.difficulty,
	}
}

func (js *JobService) AcceptJob(number int64, hash string) (bool, error) {
	if js.currentJob.hash != hash {
		return false, errors.New("wrong hash")
	}
	if js.currentJob.number == number {
		js.currentJob.mx.Lock()
		defer js.currentJob.mx.Unlock()
		drn := time.Now().Unix() - js.currentJob.created
		switch {
		case drn < time.Second.Microseconds(): // 1 second
			js.currentJob.difficulty = 100000000000
		case drn < time.Second.Microseconds()*10: // 10 second
			js.currentJob.difficulty = 100000000
		case drn < time.Second.Microseconds()*30: // 30 second
			js.currentJob.difficulty = 1000000
		case drn < time.Minute.Microseconds(): // 1 minute
			js.currentJob.difficulty = 1000
		}
		js.currentJob.created = time.Now().Unix()
		js.currentJob.last = hash
		js.currentJob.number = rand.Int63n(js.currentJob.difficulty)
		message := js.currentJob.last + strconv.FormatInt(js.currentJob.number, 10) + strconv.FormatInt(js.currentJob.created, 10)
		js.currentJob.hash = getHexedHash(message)
		return true, nil
	}
	return false, nil
}

type Job struct {
	Created    int64  `json:"created"`
	Last       string `json:"last"`
	Difficulty int64  `json:"difficulty"`
	Hash       string `json:"hash"`
}

type job struct {
	created    int64
	last       string
	number     int64
	hash       string
	difficulty int64
	mx         sync.Mutex
}

func getHexedHash(message string) string {
	h := crypto.BLAKE2b_256.New()
	h.Write([]byte(message))
	bs := h.Sum(nil)
	hexBytes := make([]byte, hex.EncodedLen(len(bs)))
	hex.Encode(hexBytes, bs)
	return "0x" + string(hexBytes)
}
