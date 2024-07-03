import java.rmi.Naming;
import java.util.Scanner;

public class MatrixClient {
    public static void main(String[] args) {
        try {
            MatrixMultiplier multiplier = (MatrixMultiplier) Naming.lookup("rmi://localhost/MatrixMultiplier");

            Scanner scanner = new Scanner(System.in);
            System.out.print("Digite o numero de linhas da matriz A: ");
            int linhaA = scanner.nextInt();
            System.out.print("Digite o numero de colunas da matriz A: ");
            int colunaA = scanner.nextInt();
            int[][] matrixA = new int[linhaA][colunaA];
            System.out.println("Preencha a matriz A:");
            for (int i = 0; i < linhaA; i++) {
                for (int j = 0; j < colunaA; j++) {
                    System.out.print("A[" + i + "][" + j + "]: ");
                    matrixA[i][j] = scanner.nextInt();
                }
            }
            System.out.print("Digite o numero de colunas da matriz B: ");
            int colunaB = scanner.nextInt();
            int[][] matrixB = new int[colunaA][colunaB];
            System.out.println("Preencha a matriz B:");
            for (int i = 0; i < colunaA; i++) {
                for (int j = 0; j < colunaB; j++) {
                    System.out.print("B[" + i + "][" + j + "]: ");
                    matrixB[i][j] = scanner.nextInt();
                }
            }
            int[][] result = multiplier.multiply(matrixA, matrixB);
            System.out.println("\nResultado da multiplicação das matrizes:");
            for (int i = 0; i < linhaA; i++) {
                for (int j = 0; j < colunaB; j++) {
                    System.out.print(result[i][j] + " ");
                }
                System.out.println();
            }
            scanner.close();
        } catch (Exception e) {
            System.err.println("Erro no cliente: " + e.toString());
            e.printStackTrace();
        }
    }
}