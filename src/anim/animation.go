package anim

type Animation interface {
	Draw()
	Update(float32)
	CheckAnimationFinished() bool
}

func CheckAnimations(arr []Animation) []Animation {
	for i := 0; i < len(arr); i++ {
		if arr[i].CheckAnimationFinished() {
			copy(arr[i:], arr[i+1:])
			arr[len(arr)-1] = nil
			arr = arr[:len(arr)-1]
		}
	}

	return arr
}
