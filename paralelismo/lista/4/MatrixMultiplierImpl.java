import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;

public class MatrixMultiplierImpl extends UnicastRemoteObject implements MatrixMultiplier {
    protected MatrixMultiplierImpl() throws RemoteException {
        super();
    }

    @Override
    public int[][] multiply(int[][] matrixA, int[][] matrixB) throws RemoteException {
        int linhaA = matrixA.length;
        int colunaA = matrixA[0].length;
        int colunaB = matrixB[0].length;
        int[][] result = new int[linhaA][colunaB];
        for (int i = 0; i < linhaA; i++) {
            for (int j = 0; j < colunaB; j++) {
                for (int k = 0; k < colunaA; k++) {
                    result[i][j] += matrixA[i][k] * matrixB[k][j];
                }
            }
        }
        return result;
    }

    public static void main(String[] args) {
        try {
            MatrixMultiplierImpl multiplier = new MatrixMultiplierImpl();
            java.rmi.registry.LocateRegistry.createRegistry(1099);
            java.rmi.Naming.rebind("MatrixMultiplier", multiplier);
            System.out.println("Servidor pronto.");
        } catch (Exception e) {
            System.err.println("Erro no servidor: " + e.toString());
            e.printStackTrace();
        }
    }
}