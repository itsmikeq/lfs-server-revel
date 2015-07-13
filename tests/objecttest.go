package tests

import (
	"github.com/revel/revel/testing"
	"github.com/memikequinn/lfs-server-go/app/models"
	"crypto/sha256"
	"encoding/hex"
)

type ObjectTest struct {
	testing.TestSuite
	Object models.Object
}

func (t *ObjectTest) Before() {
	println("Set up")
}

func (t *ObjectTest) TestObjectTestLoads() {
	t.AssertEqual(true, true)
}

func (t *ObjectTest) TestObjectTestSaves() {
	_sha := sha256.New()
	shaStr := hex.EncodeToString(_sha.Sum(nil))
	t.Object.Oid = shaStr
	myobj := models.FindObject(t.Object.Oid)
	t.AssertEqual(shaStr, myobj.Oid)
}

func (t *ObjectTest) After() {
	println("Tear Down")
}

