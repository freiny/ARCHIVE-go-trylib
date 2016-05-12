# go-trylib

<pre>
- Compare(a, b string) int
- Contains(s, substr string) bool
- ContainsAny(s, chars string) bool
- ContainsRune(s string, r rune) bool
- Count(s, sep string) int
EqualFold(s, t string) bool
Fields(s string) []string
FieldsFunc(s string, f func(rune) bool) []string
- HasPrefix(s, prefix string) bool
- HasSuffix(s, suffix string) bool
- Index(s, sep string) int
- IndexAny(s, chars string) int
- IndexByte(s string, c byte) int
- IndexFunc(s string, f func(rune) bool) int
- IndexRune(s string, r rune) int
- Join(a []string, sep string) string
- LastIndex(s, sep string) int
- LastIndexAny(s, chars string) int
- LastIndexByte(s string, c byte) int
- LastIndexFunc(s string, f func(rune) bool) int
- Map(mapping func(rune) rune, s string) string
- Repeat(s string, count int) string
Replace(s, old, new string, n int) string
- Split(s, sep string) []string
- SplitAfter(s, sep string) []string
- SplitAfterN(s, sep string, n int) []string
- SplitN(s, sep string, n int) []string
Title(s string) string
- ToLower(s string) string
- ToLowerSpecial(_case unicode.SpecialCase, s string) string
ToTitle(s string) string
ToTitleSpecial(_case unicode.SpecialCase, s string) string
- ToUpper(s string) string
- ToUpperSpecial(_case unicode.SpecialCase, s string) string
- Trim(s string, cutset string) string
- TrimFunc(s string, f func(rune) bool) string
- TrimLeft(s string, cutset string) string
- TrimLeftFunc(s string, f func(rune) bool) string
- TrimPrefix(s, prefix string) string
- TrimRight(s string, cutset string) string
- TrimRightFunc(s string, f func(rune) bool) string
- TrimSpace(s string) string
- TrimSuffix(s, suffix string) string
</pre>
