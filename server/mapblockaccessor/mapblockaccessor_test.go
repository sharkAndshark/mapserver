package mapblockaccessor

import (
	"io/ioutil"
	"mapserver/coords"
	"mapserver/db/sqlite"
	"mapserver/testutils"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestSimpleAccess(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	tmpfile, err := ioutil.TempFile("", "TestMigrate.*.sqlite")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile.Name())
	testutils.CreateTestDatabase(tmpfile.Name())

	a, err := sqlite.New(tmpfile.Name())
	if err != nil {
		panic(err)
	}

	err = a.Migrate()
	if err != nil {
		panic(err)
	}

	cache := NewMapBlockAccessor(a, 500*time.Millisecond, 1000*time.Millisecond)
	mb, err := cache.GetMapBlock(coords.NewMapBlockCoords(0, 0, 0))

	if err != nil {
		panic(err)
	}

	if mb == nil {
		t.Fatal("Mapblock is nil")
	}
}
