//autogenerated:yes
//nolint:revive,lll
package std_msgs

import (
	"github.com/aler9/goroslib/pkg/msg"
)

type Int64MultiArray struct {
	msg.Package `ros:"std_msgs"`
	Layout      MultiArrayLayout
	Data        []int64
}