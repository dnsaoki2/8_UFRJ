
public class triangulo{
	public String avalia_triangulo(int a, int b, int c) {
		if(a <= 0 || b <= 0 || c <= 0){
			return "invalido";
		}
		if(a == b && a == c && b == c) {
			return "equilatero";
		}
		if(!((a > Math.abs(b-c) && a < (b+c)) &&
				(b > Math.abs(a-c) && b < (a+c)) &&
					(c > Math.abs(a-b) && c < (b+a)))){
			return "invalido";
		}
		if((a == b && c != a) ||
				(a == c && b != a) || 
					(b == c && a != c)) {
			return "isosceles";
		}
		if((a != b) && (a != c) && (b != c)){
			return "escaleno";
		}
		return null;
	}
}
