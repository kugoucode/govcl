// 由res2go自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TFrame3 struct {
    *vcl.TFrame
    Panel1 *vcl.TPanel
    Edit1  *vcl.TEdit

    //::private::
    TFrame3Fields
}


func NewFrame3(owner vcl.IComponent) (root *TFrame3)  {
    vcl.CreateResFrame(owner, &root)
    return
}

var frame3Bytes = []byte("\x54\x50\x46\x30\x07\x54\x46\x72\x61\x6D\x65\x33\x06\x46\x72\x61\x6D\x65\x33\x04\x4C\x65\x66\x74\x02\x00\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\x40\x01\x06\x48\x65\x69\x67\x68\x74\x03\xF0\x00\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x48\x03\x54\x6F\x70\x02\x60\x05\x57\x69\x64\x74\x68\x03\xB9\x00\x06\x48\x65\x69\x67\x68\x74\x02\x29\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x50\x61\x6E\x65\x6C\x31\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x05\x54\x45\x64\x69\x74\x05\x45\x64\x69\x74\x31\x04\x4C\x65\x66\x74\x02\x28\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x79\x06\x48\x65\x69\x67\x68\x74\x02\x15\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x04\x54\x65\x78\x74\x06\x05\x45\x64\x69\x74\x31\x00\x00\x00\x00")

// 注册Form资源
var _ = vcl.RegisterFormResource(TFrame3{}, &frame3Bytes)