package models

// TextSpam defines the interface for predicting spam texts
type TextSpam interface {
	Train()
	Predict()
}
