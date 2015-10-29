import static org.junit.Assert.assertEquals;
import org.junit.Test;

public class trianguloTest {
	triangulo tr = new triangulo();
	
	@Test
	public void testEscaleno() {
		String result = tr.avalia_triangulo(5,4,3);
		assertEquals("escaleno", result);
	}
	
	@Test
	public void testEquilatero() {
		String result = tr.avalia_triangulo(1, 1, 1);
		assertEquals("equilatero", result);
	}
	
	@Test
	public void testIsosceles() {
		String result = tr.avalia_triangulo(5, 5, 8);
		assertEquals("isosceles", result);
		result = tr.avalia_triangulo(5, 8, 5);
		assertEquals("isosceles", result);
		result = tr.avalia_triangulo(8, 5, 5);
		assertEquals("isosceles", result);
	}
	
	@Test
	public void testZero() {
		String result = tr.avalia_triangulo(0, 1, 2);
		assertEquals("invalido", result);
		result = tr.avalia_triangulo(1, 0, 2);
		assertEquals("invalido", result);
		result = tr.avalia_triangulo(1, 2, 0);
		assertEquals("invalido", result);
		result = tr.avalia_triangulo(0, 0, 0);
		assertEquals("invalido", result);
	}
	
	@Test
	public void testNegativo() {
		String result = tr.avalia_triangulo(5, -3, 4);
		assertEquals("invalido", result);
	}
	
	@Test
	public void testSomaIncorreta() {
		String result = tr.avalia_triangulo(3, 2, 1);
		assertEquals("invalido", result);
		result = tr.avalia_triangulo(2, 3, 1);
		assertEquals("invalido", result);
		result = tr.avalia_triangulo(2, 1, 3);
		assertEquals("invalido", result);
		result = tr.avalia_triangulo(2, 1, 4);
		assertEquals("invalido", result);
		result = tr.avalia_triangulo(1, 4, 2);
		assertEquals("invalido", result);
		result = tr.avalia_triangulo(4, 1, 2);
		assertEquals("invalido", result);
	}
		
	@Test(expected=java.lang.Error.class)
	public void testMenosParametros() {
		tr.avalia_triangulo(2, 3);
	}
	
	@Test(expected=java.lang.Error.class)
	public void testFloat() {
		tr.avalia_triangulo(4.2, 5, 5);
	}
}
