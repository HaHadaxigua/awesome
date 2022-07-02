package register

// a code register that help embed new module into code.

type Registry struct {
}

// problem:
// 处理来自Protobuf中的oneof问题。
// oneof将会创建一个接口，但是我们无法获取所有实现了这个接口的类型，因此只能进行枚举。
// 即：实际上存在两个接口 interface in proto and interface which we describe.
// we called the interface in proto as pbInterface, and the other called as weInterface.
// in fact, the req we get is all pbInterface, so we need to assume the real type in our system.
// pbInterface->pbType->weType

// so is there a simple way to describe the process?
// 1. register, register know all the info about pbInterface, pbType and weType
// whenever we meet a pbInterface, it will extrapolate the pbType and convert it to we type.

// 2. auto register, use code generate
