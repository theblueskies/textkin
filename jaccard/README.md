# Jaccard similarity

This package implements Jaccard similarity confidence

Usage:
```go
j := NewJaccardSim("some text", "random text")

jaccardConfidence := j.GetConfidence()
fmt.Println(jaccardConfidence) // 0.3333333333333333
```
