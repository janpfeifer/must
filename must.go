package must

import "log"

// M logs and panics if `err` is not nil.
//
// This function is used by all other variants (M1, ..., M9), and if you want
// a different error behavior (like `log.Fatalf` or similar), just reassign M
// to your particular use, and all other functions will pick it up.
var M = func(err error) {
	if err != nil {
		log.Printf("Must not error: %+v\nPanicking ...\n\n", err)
		panic(err)
	}
}

/* Code used to generate these in GoNB:

import "fmt"

%%
for ii := 1; ii < 10; ii++ {
    var Tparts, paramParts, retTypeParts, retParts []string
    for jj := 1; jj <= ii; jj++ {
        Tparts = append(Tparts, fmt.Sprintf("T%d any", jj))
        paramParts = append(paramParts, fmt.Sprintf("value%d T%d", jj, jj))
        retTypeParts = append(retTypeParts, fmt.Sprintf("T%d", jj))
        retParts = append(retParts, fmt.Sprintf("value%d", jj))
    }

    fmt.Printf("// M%d checks that there is no error with `M(err)` and then simply returns the values given.\n", ii)
    fmt.Printf("func M%d[%s](%s, err error) (%s) {\n", ii,
               strings.Join(Tparts, ", "), strings.Join(paramParts, ", "), strings.Join(retTypeParts, ", "))
    fmt.Printf("\tM(err)\n\treturn %s\n}\n\n", strings.Join(retParts, ", "))
}

*/

// M1 checks that there is no error with `M(err)` and then simply returns the values given.
func M1[T1 any](value1 T1, err error) T1 {
	M(err)
	return value1
}

// M2 checks that there is no error with `M(err)` and then simply returns the values given.
func M2[T1 any, T2 any](value1 T1, value2 T2, err error) (T1, T2) {
	M(err)
	return value1, value2
}

// M3 checks that there is no error with `M(err)` and then simply returns the values given.
func M3[T1 any, T2 any, T3 any](value1 T1, value2 T2, value3 T3, err error) (T1, T2, T3) {
	M(err)
	return value1, value2, value3
}

// M4 checks that there is no error with `M(err)` and then simply returns the values given.
func M4[T1 any, T2 any, T3 any, T4 any](value1 T1, value2 T2, value3 T3, value4 T4, err error) (T1, T2, T3, T4) {
	M(err)
	return value1, value2, value3, value4
}

// M5 checks that there is no error with `M(err)` and then simply returns the values given.
func M5[T1 any, T2 any, T3 any, T4 any, T5 any](value1 T1, value2 T2, value3 T3, value4 T4, value5 T5, err error) (T1, T2, T3, T4, T5) {
	M(err)
	return value1, value2, value3, value4, value5
}

// M6 checks that there is no error with `M(err)` and then simply returns the values given.
func M6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](value1 T1, value2 T2, value3 T3, value4 T4, value5 T5, value6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	M(err)
	return value1, value2, value3, value4, value5, value6
}

// M7 checks that there is no error with `M(err)` and then simply returns the values given.
func M7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](value1 T1, value2 T2, value3 T3, value4 T4, value5 T5, value6 T6, value7 T7, err error) (T1, T2, T3, T4, T5, T6, T7) {
	M(err)
	return value1, value2, value3, value4, value5, value6, value7
}

// M8 checks that there is no error with `M(err)` and then simply returns the values given.
func M8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](value1 T1, value2 T2, value3 T3, value4 T4, value5 T5, value6 T6, value7 T7, value8 T8, err error) (T1, T2, T3, T4, T5, T6, T7, T8) {
	M(err)
	return value1, value2, value3, value4, value5, value6, value7, value8
}

// M9 checks that there is no error with `M(err)` and then simply returns the values given.
func M9[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any, T9 any](value1 T1, value2 T2, value3 T3, value4 T4, value5 T5, value6 T6, value7 T7, value8 T8, value9 T9, err error) (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	M(err)
	return value1, value2, value3, value4, value5, value6, value7, value8, value9
}
