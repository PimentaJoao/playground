import java.io.*;
import java.net.*;

public class Client {
    public static void main(String[] args) {
        try {
            Socket socket = new Socket("localhost", 12345); // Conecta ao servidor local na porta 12345

            PrintWriter out = new PrintWriter(socket.getOutputStream(), true);
            BufferedReader in = new BufferedReader(new InputStreamReader(socket.getInputStream()));

            // Envia limite para calcular números primos
            int limit = 100;
            out.println(limit);

            // Recebe resposta do servidor
            String response = in.readLine();
            System.out.println("Números primos até " + limit + ": " + response);

            // Fecha as streams e o socket
            out.close();
            in.close();
            socket.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}