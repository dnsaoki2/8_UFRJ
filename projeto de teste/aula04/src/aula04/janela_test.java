package aula04;

import static org.junit.Assert.*;

import org.junit.Test;

public class janela_test {

	@Test
	public void Folha1B1() {
		int result = janela.calculaJanela(-1, 200, 400);
		//assertTrue(true);
		assertEquals(-1,result);
	}
	
	@Test
	public void Folha1B2() {
		int result = janela.calculaJanela(0, 200, 400);
		assertEquals(0,result);
	}
	
	@Test
	public void Folha1B3() {
		int result = janela.calculaJanela(100, 200, 400);
		assertEquals(10000,result);
	}

	@Test
	public void Folha1B4() {
		int result = janela.calculaJanela(400, 200, 400);
		assertEquals(20000,result);
	}
	
	@Test
	public void Folha1B5() {
		int result = janela.calculaJanela(500, 200, 400);
		assertEquals(-1,result);
	}
	
	@Test
	public void Folha2B1() {
		int result = janela.calculaJanela(0, -1, 400);
		assertEquals(-1,result);
	}
	
	@Test
	public void Folha2B2() {
		int result = janela.calculaJanela(0, 0, 400);
		assertEquals(20000,result);
	}
	
	@Test
	public void Folha2B3() {
		int result = janela.calculaJanela(0, 200, 400);
		assertEquals(0,result);
	}

	@Test
	public void Folha2B4() {
		int result = janela.calculaJanela(0, 400, 400);
		assertEquals(20000,result);
	}
	
	@Test
	public void Folha2B5() {
		int result = janela.calculaJanela(0, 500, 400);
		assertEquals(-1,result);
	}
	
	@Test
	public void Folha3B1() {
		int result = janela.calculaJanela(0, 200, -1);
		assertEquals(-1,result);
	}
	
	@Test
	public void Folha3B2() {
		int result = janela.calculaJanela(0, 200, 0);
		assertEquals(20000,result);
	}
	
	@Test
	public void Folha3B3() {
		int result = janela.calculaJanela(0, 200, 300);
		assertEquals(10000,result);
	}

	@Test
	public void Folha3B4() {
		int result = janela.calculaJanela(0, 200, 400);
		assertEquals(result, 0);
	}
	
	@Test
	public void Folha3B5() {
		int result = janela.calculaJanela(0, 200, 500);
		assertEquals(-1,result);
	}
}
