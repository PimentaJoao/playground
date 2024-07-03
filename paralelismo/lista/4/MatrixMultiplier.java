import java.rmi.Remote;
import java.rmi.RemoteException;

public interface MatrixMultiplier extends Remote {
    int[][] multiply(int[][] matrixA, int[][] matrixB) throws RemoteException;
}