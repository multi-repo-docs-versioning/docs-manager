package repository

import (
	"io/ioutil"
	"os"
	"testing"

	"gopkg.in/src-d/go-git.v4"

	"gotest.tools/assert"
)

func TestClone(t *testing.T) {
	dir, err := ioutil.TempDir("/tmp", "test")
	if err != nil {
		panic(err)
	}
	rep := Repository{

		path: dir,
		cloneOptions: git.CloneOptions{
			URL:      "https://github.com/onosproject/onos-config",
			Tags:     git.AllTags,
			Progress: os.Stdout,
		},
	}
	err = rep.Clone()
	assert.Equal(t, err, nil)
	err = os.RemoveAll(dir)
	assert.Equal(t, err, nil)

}

func TestGetTags(t *testing.T) {
	dir, err := ioutil.TempDir("/tmp", "test")
	if err != nil {
		panic(err)
	}
	rep := Repository{
		path: dir,
		cloneOptions: git.CloneOptions{
			URL:      "https://github.com/onosproject/onos-config",
			Tags:     git.AllTags,
			Progress: os.Stdout,
		},
		tagName: "v0.1-onfconnect",
	}
	err = rep.Clone()
	assert.Equal(t, err, nil)
	_, err = rep.GetTag()
	assert.Equal(t, err, nil)
	err = os.RemoveAll(dir)
	assert.Equal(t, err, nil)
}

func TestCheckOut(t *testing.T) {
	dir, err := ioutil.TempDir("/tmp", "test")
	if err != nil {
		panic(err)
	}
	rep := Repository{
		path: dir,
		cloneOptions: git.CloneOptions{
			URL:      "https://github.com/onosproject/onos-config",
			Tags:     git.AllTags,
			Progress: os.Stdout,
		},
		tagName: "v0.1-onfconnect",
	}
	err = rep.Clone()
	assert.Equal(t, err, nil)
	err = rep.CheckOutTag()
	assert.Equal(t, err, nil)
	err = os.RemoveAll(dir)
	assert.Equal(t, err, nil)
}
