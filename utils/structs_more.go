package utils

type VectRound interface {
	Round() int
}

type VectMod interface {
	Mod(coeff int) int
}

type Vect2D struct {
	x, y int
}

type Vect3D struct {
	Vect2D
	z int
}

func (vm Vect3D) Round() {
	xplus := vm.x + vm.y
	println("Sum: ", xplus)
}

// Impement interface `VectRound`
func (v2 Vect2D) Round() int {
	return v2.x + v2.y
}

// Impement interface  `VectMod`
func (v2 Vect2D) Mod(i int) int {
	return v2.x + v2.y + i
}

// Impement interface `VectRound`
func (v3 Vect3D) Mod(i int) int {
	return v3.x + v3.y + v3.z + i
}

func TestVectRound(vr VectRound) {
	res := vr.Round()
	println(res)
}

func TestVectMod(vr VectMod, i int) {
	res := vr.Mod(i)
	println(res)
}

func TestStructs() {
	println("====>>>")
	v2d := Vect2D{x: 10, y: 40}
	TestVectRound(v2d)
	TestVectMod(v2d, 30)

	v3d := Vect3D{}
	v3d.x = 10
	v3d.y = 20
	v3d.z = 30
	TestVectMod(v3d, 40)
}
