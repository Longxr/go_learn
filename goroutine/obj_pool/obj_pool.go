package obj_pool

import (
	"errors"
	"time"
)

// ReusableObj 对象池
type ReusableObj struct {
}

// ObjPool 对象池
type ObjPool struct {
	bufChan chan *ReusableObj
}

// NewObjPool 创建对象池
func NewObjPool(numOfObj int) *ObjPool {
	objPool := &ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return objPool
}

// GetObj 获取对象
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout): //超时控制
		return nil, errors.New("time out")
	}
}

// ReleaseObj 释放对象
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
