import java.io.*;
import java.net.*;
import java.util.*;

public class CentralServer {
    public static void main(String[] args) {
        try {
            ServerSocket serverSocket = new ServerSocket(12345); // Porta do servidor
            System.out.println("Servidor esperando conexões na porta 12345...");

            while (true) {
                try {
                    Socket clientSocket = serverSocket.accept(); // Aceita conexões dos clientes
                    
                    System.out.println("Cliente conectado: " + clientSocket.getInetAddress().getHostName());

                    // Iniciar uma thread para lidar com o cliente
                    ClientHandler handler = new ClientHandler(clientSocket);
                    handler.start();
                } catch (Exception e) {
                    serverSocket.close();
                }
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}

class ClientHandler extends Thread {
    private Socket clientSocket;

    public ClientHandler(Socket socket) {
        this.clientSocket = socket;
    }

    @Override
    public void run() {
        try {
            BufferedReader in = new BufferedReader(new InputStreamReader(clientSocket.getInputStream()));
            PrintWriter out = new PrintWriter(clientSocket.getOutputStream(), true);

            // Recebe mensagem do cliente (faixa de números para calcular)
            String inputLine = in.readLine();
            int limit = Integer.parseInt(inputLine);

            // Calcula números primos até o limite usando o Crivo de Eratóstenes
            List<Integer> primes = sieveOfEratosthenes(limit);

            // Converte lista de primos em uma string para enviar de volta ao cliente
            StringBuilder primeNumbers = new StringBuilder();

            for (int prime : primes) {
                primeNumbers.append(prime).append(" ");
            }
            // Envia resposta de volta para o cliente
            out.println(primeNumbers.toString());

            // Fecha os streams e o socket
            in.close();
            out.close();
            clientSocket.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    // Implementação do Crivo de Eratóstenes
    private List<Integer> sieveOfEratosthenes(int limit) {
        boolean[] isPrime = new boolean[limit + 1];
        Arrays.fill(isPrime, true);
        List<Integer> primes = new ArrayList<>();
        for (int p = 2; p * p <= limit; p++) {
            if (isPrime[p]) {
                for (int i = p * p; i <= limit; i += p) {
                    isPrime[i] = false;
                }
            }
        }
        for (int p = 2; p <= limit; p++) {
            if (isPrime[p]) {
                primes.add(p);
            }
        }
        return primes;
    }
}