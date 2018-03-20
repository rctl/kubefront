package core

//Config contains the configuration parameters needed for kubefront to function
type Config struct {
	JWTSecret string
	Namespace string
}
