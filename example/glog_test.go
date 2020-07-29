package example

import (
	"testing"

	"github.com/huweihuang/logger/glog"
)

func TestGlog(t *testing.T){
	glog.Debugf("test debug")
	glog.Errorf("test error")
	glog.Infof("test info")
	glog.Warningf("test warning")
	glog.V(2).Infof("test v2 info")
	glog.V(4).Infof("test v4 info")
}
