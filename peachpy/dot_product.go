package peachpy

//go:generate python3 -m peachpy.x86_64 dot_product.py -S -o dot_product_amd64.s -mabi=goasm
func DotProduct(x *float32, y *float32, length uint) float32
