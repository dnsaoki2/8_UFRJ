package aula04;

import java.util.Scanner;

public class janela {
	public static int calculaJanela2(int p1, int p2, int p3) {    
	    int[] vet = new int[601];
	    if (p1 < 0 || p2 < 0 || p3 < 0 || p1 > 400 || p2 > 400 || p3 > 400)
	    	return -1;
	    for (int i = 0; i <= 600; i++) vet[i] = 0;
	    for (int i = p1; i < p1+200 && i <= 600; i++){
	    	vet[i] = 1;
	    }
	    for (int i = p2; i < p2+200 && i <= 600; i++){
	    	vet[i] = 1;
	    }
	    for (int i = p3; i < p3+200 && i <= 600; i++){
	    	vet[i] = 1;
	    }
	    int ans = 0;
	    for (int i = 0; i <= 600-1; i++){
	    	if (vet[i] == 1) continue;
	    	int cnt = 1;
	    	for (int j = i + 1; j<= 600-1; j++){
	    		i = j;
	    		if (vet[j] == 1) {
	    			break;
	    		}
	    		cnt++;
	    	}
	    	ans += cnt*100;
	    }
	    return ans;
	}
	
	//quebra - janela2.java
//	public static int calculaJanela(int p1, int p2, int p3){
//		int i, soma =0;
//		int [] centimetros = new int[600];
//        
//		for(i=0;i<600;i++){
//            centimetros[i] = 0;
//        }
//        
//        for(i=p1;i<p1+200;i++){
//            centimetros[i] = 1;
//        }
//        
//        for(i=p2;i<=p2+200;i++){
//            centimetros[i] = 1;
//        }
//        
//        for(i=p3;i<p3+200;i++){
//            centimetros[i] = 1;
//        }
//        
//        for(i=1;i<600;i++){
//            if(centimetros[i]==1)
//            	soma++;
//        }
//        
//        return soma*100;
//	}
	
	//quebra todas - janela22.java
//	public static int calculaJanela(int p1, int p2, int p3){
//		int i, soma =0;
//		int [] centimetros = new int[600];
//        
//		for(i=0;i<600;i++){
//            centimetros[i] = 0;
//        }
//        
//        for(i=p1;i<p1+200;i++){
//            centimetros[i] = 0;
//        }
//        
//        for(i=0;i<p2+200;i++){
//            centimetros[i] = 1;
//        }
//        
//        for(i=p3;i<200;i++){
//            centimetros[i] = 1;
//        }
//        
//        for(i=0;i<600;i++){
//            if(centimetros[i]==1)
//            	soma++;
//        }
//        
//        return soma*100;
//	}
	
	//quebra algumas janelac2.java
	public static int calculaJanela(int p1, int p2, int p3) {
	    if (p3 - p1 < 200) {
	        return 100 * (400 - p3 + p1);
	    } else if (p2 - p1 < 200 && p3 - p2 < 200) {
	        return 100 * (400 - p3 + p1);
	    } else if (p2 - p1 < 200 && p3 - p2 > 200) {
	        return 100 * (200 - p2 + p1);
	    } else if (p3 - p2 < 200 && p2 - p1 > 200) {
	        return 100 * (200 - p3 + p2);
	    } else {
	        return 0;
	    }
	}

	//mais correto - janelaC
//	public static int calculaJanela(int p1, int p2, int p3) {
//	    
//	    if (p3 - p1 <= 200) {
//	        return 100 * (400 - p3 + p1);
//	    } else if (p2 - p1 <= 200 && p3 - p2 <= 200) {
//	        return 100 * (400 - p3 + p1);
//	    } else if (p2 - p1 <= 200 && p3 - p2 > 200) {
//	        return 100 * (200 - p2 + p1);
//	    } else if (p3 - p2 <= 200 && p2 - p1 > 200) {
//	        return 100 * (200 - p3 + p2);
//	    } else {
//	        return 0;
//	    }
//	    
//	}

	public static void main(String[] args) {
    	
    	Scanner scanner = new Scanner(System.in);
        int posFolha1 = scanner.nextInt();
        int posFolha2 = scanner.nextInt();
        int posFolha3 = scanner.nextInt();
        
        scanner.close();
    	
        System.out.println(janela.calculaJanela(posFolha1, posFolha2, posFolha3));
       // System.out.println(janela.calculaJanelaCorreta(posFolha1, posFolha2, posFolha3));
    	
    	

    }
}