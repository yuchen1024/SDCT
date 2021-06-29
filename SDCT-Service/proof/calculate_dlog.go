package proof

import (
	"crypto/elliptic"
	"encoding/hex"
	"io/ioutil"
	"math/big"
	"os"
	"sync"

	log "github.com/inconshreveable/log15"
	"github.com/sdct/utils"
)

// Global params
var (
	Compressed  = 33
	hashMapFile = "hashMap"
)

var point2Index = make(map[string]int, 0)

// BuildAndLoadMapIfNotExist tries to build map if it's not exist.
// Load map if exist
func BuildAndLoadMapIfNotExist(g *utils.ECPoint, rangeLen, tunning, roNum int) {
	if _, err := os.Stat(hashMapFile); os.IsNotExist(err) {
		log.Info("build hash map, cost about 1 minute")
		buildHashMap(g, rangeLen, tunning, roNum)
		return
	}
	loadHashMap(rangeLen, tunning)
}

func buildHashMap(g *utils.ECPoint, rangeLen, tunning, roNum int) {
	giantStepSize := 2 << (rangeLen/2 + tunning - 1)

	if giantStepSize%roNum != 0 {
		panic("Thread assignment fails")
	}

	l := giantStepSize / roNum
	buffer := make([]byte, giantStepSize*Compressed)

	wg := sync.WaitGroup{}
	wg.Add(roNum)

	for i := 0; i < roNum; i++ {
		go func(i int) {
			startPoint := new(utils.ECPoint).ScalarMult(g, big.NewInt(int64(i*l)))
			calPoints(startPoint, g, buffer, l, i*l)
			wg.Done()
		}(i)
	}

	wg.Wait()

	// write to file
	if err := ioutil.WriteFile(hashMapFile, buffer, 0644); err != nil {
		panic(err)
	}

	// load data
	loadHashMap(rangeLen, tunning)
}

func loadHashMap(rangeLen, tunning int) {
	giantStepSize := 2 << (rangeLen/2 + tunning - 1)
	bytesLen := giantStepSize * Compressed

	// read file
	bytes, err := ioutil.ReadFile(hashMapFile)
	if err != nil {
		panic(err)
	}

	if bytesLen != len(bytes) {
		panic("Invalid hash map file bytes len")
	}

	for i := 0; i < giantStepSize; i++ {
		key := hex.EncodeToString(bytes[i*Compressed : (i+1)*Compressed])
		point2Index[key] = i
	}
}

// ShanksDlog decrypts point using shanks algorithm.
func ShanksDlog(g, msg *utils.ECPoint, rangeLen, tunning int) *big.Int {
	giantStepSize := 2 << (rangeLen/2 + tunning - 1)
	loop := 2 << (rangeLen/2 - tunning - 1)

	giantStep := new(utils.ECPoint).ScalarMult(g, big.NewInt(int64(giantStepSize)))
	giantStep.Negation(giantStep)

	dstPoint := msg.Copy()
	if giantStepSize != len(point2Index) {
		panic("Hash map isn't loaded")
	}

	i, j := 0, 0
	find := false
	for ; j < loop; j++ {
		msgKey := hex.EncodeToString(elliptic.MarshalCompressed(dstPoint.Curve, dstPoint.X, dstPoint.Y))
		r, ok := point2Index[msgKey]
		if ok {
			find = true
			i = r
			break
		}

		dstPoint.Add(dstPoint, giantStep)
	}

	if !find {
		panic("The DLOG is not found in the specified range")
	}

	return big.NewInt(int64(i + j*giantStepSize))
}

func calPoints(start, g *utils.ECPoint, buffer []byte, l, startIndex int) {
	for i := 0; i < l; i++ {
		index := (startIndex + i) * Compressed
		copy(buffer[index:index+Compressed], elliptic.MarshalCompressed(start.Curve, start.X, start.Y))
		start.Add(start, g)
	}
}
