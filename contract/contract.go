package contract

type (
	Initialize = func(url string) Convert
	Convert    = func(code string) (int, error)
)
