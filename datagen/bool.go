package datagen

func Bool() bool {
	return Rand().Int(2) == 1
}
