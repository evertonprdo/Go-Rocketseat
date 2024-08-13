package bar

import "advancedConcepts/foo"

func TakeFoo(i foo.Interface) {
	i.Interface()
}
